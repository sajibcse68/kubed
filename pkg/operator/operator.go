package operator

import (
	"encoding/json"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/appscode/envconfig"
	"github.com/appscode/go/log"
	v "github.com/appscode/go/version"
	"github.com/appscode/kubed/pkg/api"
	"github.com/appscode/kubed/pkg/backup"
	"github.com/appscode/kubed/pkg/config"
	"github.com/appscode/kubed/pkg/elasticsearch"
	"github.com/appscode/kubed/pkg/eventer"
	"github.com/appscode/kubed/pkg/indexers"
	"github.com/appscode/kubed/pkg/influxdb"
	rbin "github.com/appscode/kubed/pkg/recyclebin"
	"github.com/appscode/kubed/pkg/storage"
	"github.com/appscode/kubed/pkg/syncer"
	"github.com/appscode/kubed/pkg/util"
	"github.com/appscode/pat"
	srch_cs "github.com/appscode/searchlight/client/typed/monitoring/v1alpha1"
	scs "github.com/appscode/stash/client/typed/stash/v1alpha1"
	vcs "github.com/appscode/voyager/client/typed/voyager/v1beta1"
	shell "github.com/codeskyblue/go-sh"
	pcm "github.com/coreos/prometheus-operator/pkg/client/monitoring/v1"
	prom "github.com/coreos/prometheus-operator/pkg/client/monitoring/v1"
	kcs "github.com/k8sdb/apimachinery/client/typed/kubedb/v1alpha1"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Options struct {
	Master     string
	KubeConfig string

	ConfigPath string
	APIAddress string
	WebAddress string

	EnableConfigSync  bool
	ScratchDir        string
	OperatorNamespace string

	ResyncPeriod time.Duration
}

type Operator struct {
	KubeClient        clientset.Interface
	VoyagerClient     vcs.VoyagerV1beta1Interface
	SearchlightClient srch_cs.MonitoringV1alpha1Interface
	StashClient       scs.StashV1alpha1Interface
	PromClient        pcm.MonitoringV1Interface
	KubeDBClient      kcs.KubedbV1alpha1Interface

	Opt    Options
	Config config.ClusterConfig

	searchIndex    *indexers.ResourceIndexer
	reverseIndex   *indexers.ReverseIndexer
	trashCan       *rbin.RecycleBin
	eventer        *eventer.EventForwarder
	cron           *cron.Cron
	notifierLoader envconfig.LoaderFunc
	configSyncer   *syncer.ConfigSyncer

	m sync.Mutex
}

func (op *Operator) SearchIndex() *indexers.ResourceIndexer {
	op.m.Lock()
	defer op.m.Unlock()
	return op.searchIndex
}

func (op *Operator) ReverseIndex() *indexers.ReverseIndexer {
	op.m.Lock()
	defer op.m.Unlock()
	return op.reverseIndex
}

func (op *Operator) TrashCan() *rbin.RecycleBin {
	op.m.Lock()
	defer op.m.Unlock()
	return op.trashCan
}

func (op *Operator) Eventer() *eventer.EventForwarder {
	op.m.Lock()
	defer op.m.Unlock()
	return op.eventer
}

func (op *Operator) NotifierLoader() envconfig.LoaderFunc {
	op.m.Lock()
	defer op.m.Unlock()
	return op.notifierLoader
}

func (op *Operator) Setup() error {
	cfg, err := config.LoadConfig(op.Opt.ConfigPath)
	if err != nil {
		return err
	}
	if op.Opt.APIAddress != "" {
		cfg.APIServer.Address = op.Opt.APIAddress
	}
	err = cfg.Validate()
	if err != nil {
		return err
	}
	op.Config = *cfg

	op.notifierLoader, err = op.getLoader()
	if err != nil {
		return err
	}

	if op.Config.RecycleBin != nil {
		if op.Config.RecycleBin.Path == "" {
			op.Config.RecycleBin.Path = filepath.Join(op.Opt.ScratchDir, "transhcan")
		}
		op.trashCan = &rbin.RecycleBin{
			ClusterName: op.Config.ClusterName,
			Spec:        *op.Config.RecycleBin,
			Loader:      op.notifierLoader,
		}
	}

	if op.Config.EventForwarder != nil {
		op.eventer = &eventer.EventForwarder{
			ClusterName: op.Config.ClusterName,
			Receivers:   op.Config.EventForwarder.Receivers,
			Loader:      op.notifierLoader,
		}
	}

	if op.Config.EnableConfigSyncer {
		op.configSyncer = &syncer.ConfigSyncer{KubeClient: op.KubeClient}
	}

	op.cron = cron.New()
	op.cron.Start()

	for _, j := range cfg.Janitors {
		if j.Kind == config.JanitorInfluxDB {
			janitor := influx.Janitor{Spec: *j.InfluxDB, TTL: j.TTL.Duration}
			err = janitor.Cleanup()
			if err != nil {
				return err
			}
		}
	}

	// Enable full text indexing to have search feature
	indexDir := filepath.Join(op.Opt.ScratchDir, "bleve")
	if op.Config.APIServer.EnableSearchIndex {
		si, err := indexers.NewResourceIndexer(indexDir)
		if err != nil {
			return err
		}
		op.searchIndex = si
	}
	// Enable pod -> service, service -> serviceMonitor indexing
	if op.Config.APIServer.EnableReverseIndex {
		ri, err := indexers.NewReverseIndexer(op.KubeClient, op.PromClient, indexDir)
		if err != nil {
			return err
		}
		op.reverseIndex = ri
	}

	op.Opt.ResyncPeriod = time.Minute * 2
	return nil
}

func (op *Operator) getLoader() (envconfig.LoaderFunc, error) {
	if op.Config.NotifierSecretName == "" {
		return func(key string) (string, bool) {
			return "", false
		}, nil
	}
	cfg, err := op.KubeClient.CoreV1().
		Secrets(op.Opt.OperatorNamespace).
		Get(op.Config.NotifierSecretName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return func(key string) (value string, found bool) {
		var bytes []byte
		bytes, found = cfg.Data[key]
		value = string(bytes)
		return
	}, nil
}

func (op *Operator) RunWatchers() {
	go op.WatchAlertmanagerV1()
	go op.WatchCertificateSigningRequests()
	go op.WatchClusterAlerts()
	go op.WatchClusterRoleBindingV1alpha1()
	go op.WatchClusterRoleBindingV1beta1()
	go op.WatchClusterRoleV1alpha1()
	go op.WatchClusterRoleV1beta1()
	go op.WatchConfigMaps()
	go op.WatchDaemonSets()
	go op.WatchDeploymentApps()
	go op.WatchDeploymentExtensions()
	go op.WatchDormantDatabases()
	go op.WatchElasticsearches()
	go op.WatchEvents()
	go op.WatchIngresses()
	go op.WatchJobs()
	go op.watchNamespaces()
	go op.WatchNodeAlerts()
	go op.WatchNodes()
	go op.WatchPersistentVolumeClaims()
	go op.WatchPersistentVolumes()
	go op.WatchPodAlerts()
	go op.WatchPostgreses()
	go op.WatchPrometheusV1()
	go op.WatchReplicaSets()
	go op.WatchReplicationControllers()
	go op.WatchRestics()
	go op.WatchRoleBindingV1alpha1()
	go op.WatchRoleBindingV1beta1()
	go op.WatchRoleV1alpha1()
	go op.WatchRoleV1beta1()
	go op.WatchSecrets()
	go op.watchService()
	go op.WatchEndpoints()
	go op.WatchServiceMonitorV1()
	go op.WatchStatefulSets()
	go op.WatchStorageClassV1()
	go op.WatchStorageClassV1beta1()
	go op.WatchVoyagerCertificates()
	go op.WatchVoyagerIngresses()
}

func (op *Operator) RunAPIServer() {
	// router is default HTTP request multiplexer for kubed. It matches the URL of each
	// incoming request against a list of registered patterns with their associated
	// methods and calls the handler for the pattern that most closely matches the
	// URL.
	//
	// Pattern matching attempts each pattern in the order in which they were
	// registered.
	router := pat.New()
	if op.Config.APIServer.EnableSearchIndex {
		op.searchIndex.RegisterRouters(router)
	}
	// Enable pod -> service, service -> serviceMonitor indexing
	if op.Config.APIServer.EnableReverseIndex {
		router.Get("/api/v1/namespaces/:namespace/:resource/:name/services", http.HandlerFunc(op.reverseIndex.Service.ServeHTTP))
		if util.IsPreferredAPIResource(op.KubeClient, prom.Group+"/"+prom.Version, prom.ServiceMonitorsKind) {
			// Add Indexer only if Server support this resource
			router.Get("/apis/"+prom.Group+"/"+prom.Version+"/namespaces/:namespace/:resource/:name/"+prom.ServiceMonitorName, http.HandlerFunc(op.reverseIndex.ServiceMonitor.ServeHTTP))
		}
		if util.IsPreferredAPIResource(op.KubeClient, prom.Group+"/"+prom.Version, prom.PrometheusesKind) {
			// Add Indexer only if Server support this resource
			router.Get("/apis/"+prom.Group+"/"+prom.Version+"/namespaces/:namespace/:resource/:name/"+prom.PrometheusName, http.HandlerFunc(op.reverseIndex.Prometheus.ServeHTTP))
		}
	}

	router.Get("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }))
	router.Get("/metadata", http.HandlerFunc(op.metadataHandler))
	log.Fatalln(http.ListenAndServe(op.Config.APIServer.Address, router))
}

func (op *Operator) metadataHandler(w http.ResponseWriter, r *http.Request) {
	resp := &api.KubedMetadata{
		OperatorNamespace:   op.Opt.OperatorNamespace,
		SearchEnabled:       op.Config.APIServer.EnableSearchIndex,
		ReverseIndexEnabled: op.Config.APIServer.EnableReverseIndex,
		Version:             &v.Version,
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-content-type-options", "nosniff")
}

func (op *Operator) RunElasticsearchCleaner() error {
	for _, j := range op.Config.Janitors {
		if j.Kind == config.JanitorElasticsearch {
			janitor := es.Janitor{Spec: *j.Elasticsearch, TTL: j.TTL.Duration}
			err := janitor.Cleanup()
			if err != nil {
				return err
			}
			op.cron.AddFunc("@every 1h", func() {
				err := janitor.Cleanup()
				if err != nil {
					log.Errorln(err)
				}
			})
		}
	}
	return nil
}

func (op *Operator) RunTrashCanCleaner() error {
	if op.trashCan == nil {
		return nil
	}

	return op.cron.AddFunc("@every 1h", func() {
		err := op.trashCan.Cleanup()
		if err != nil {
			log.Errorln(err)
		}
	})
}

func (op *Operator) RunSnapshotter() error {
	if op.Config.Snapshotter == nil {
		return nil
	}

	osmconfigPath := filepath.Join(op.Opt.ScratchDir, "osm", "config.yaml")
	err := storage.WriteOSMConfig(op.KubeClient, op.Config.Snapshotter.Backend, op.Opt.OperatorNamespace, osmconfigPath)
	if err != nil {
		return err
	}

	container, err := op.Config.Snapshotter.Backend.Container()
	if err != nil {
		return err
	}

	// test credentials
	sh := shell.NewSession()
	sh.SetDir(op.Opt.ScratchDir)
	sh.ShowCMD = true
	err = sh.Command("osm", "lc", "--osmconfig", osmconfigPath).Run()
	if err != nil {
		return err
	}
	snapshotter := func() error {
		cfg, err := clientcmd.BuildConfigFromFlags(op.Opt.Master, op.Opt.KubeConfig)
		if err != nil {
			return err
		}

		t := time.Now().UTC()
		ts := t.Format(config.TimestampFormat)
		snapshotRoot := filepath.Join(op.Opt.ScratchDir, "snapshot")
		snapshotDir := filepath.Join(snapshotRoot, ts)
		snapshotFile := filepath.Join(snapshotRoot, ts+".tar.gz")
		err = backup.SnapshotCluster(cfg, snapshotDir, op.Config.Snapshotter.Sanitize)
		if err != nil {
			return err
		}
		defer func() {
			if err := os.RemoveAll(snapshotDir); err != nil {
				log.Errorln(err)
			}
			if err := os.Remove(snapshotFile); err != nil {
				log.Errorln(err)
			}
		}()

		dest, err := op.Config.Snapshotter.Location(t)
		if err != nil {
			return err
		}

		sh := shell.NewSession()
		sh.SetDir(snapshotRoot)
		sh.ShowCMD = true

		err = sh.Command("tar", "-czf", ts+".tar.gz", ts).Run()
		if err != nil {
			return err
		}
		return sh.Command("osm", "push", "--osmconfig", osmconfigPath, "-c", container, snapshotFile, dest).Run()
	}
	// start taking first backup
	go func() {
		err := snapshotter()
		if err != nil {
			log.Errorln(err)
		}
	}()
	return op.cron.AddFunc(op.Config.Snapshotter.Schedule, func() {
		err := snapshotter()
		if err != nil {
			log.Errorln(err)
		}
	})
}

func (op *Operator) RunAndHold() {
	var err error

	err = op.RunElasticsearchCleaner()
	if err != nil {
		log.Fatalln(err)
	}

	err = op.RunTrashCanCleaner()
	if err != nil {
		log.Fatalln(err)
	}

	err = op.RunSnapshotter()
	if err != nil {
		log.Fatalln(err)
	}

	op.RunWatchers()
	go op.RunAPIServer()

	m := pat.New()
	m.Get("/metrics", promhttp.Handler())
	http.Handle("/", m)
	log.Infoln("Listening on", op.Opt.WebAddress)
	log.Fatal(http.ListenAndServe(op.Opt.WebAddress, nil))
}
