package operator

import (
	"reflect"

	"github.com/appscode/go/log"
	acrt "github.com/appscode/go/runtime"
	"github.com/appscode/kubed/pkg/util"
	kutil "github.com/appscode/kutil/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
)

// Blocks caller. Intended to be called as a Go routine.
func (op *Operator) WatchEndpoints() {
	if !util.IsPreferredAPIResource(op.KubeClient, apiv1.SchemeGroupVersion.String(), "Endpoints") {
		log.Warningf("Skipping watching non-preferred GroupVersion:%s Kind:%s", apiv1.SchemeGroupVersion.String(), "Endpoints")
		return
	}

	defer acrt.HandleCrash()

	lw := &cache.ListWatch{
		ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
			return op.KubeClient.CoreV1().Endpoints(apiv1.NamespaceAll).List(metav1.ListOptions{})
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return op.KubeClient.CoreV1().Endpoints(apiv1.NamespaceAll).Watch(metav1.ListOptions{})
		},
	}
	_, ctrl := cache.NewInformer(lw,
		&apiv1.Endpoints{},
		op.Opt.ResyncPeriod,
		cache.ResourceEventHandlerFuncs{
			UpdateFunc: func(oldObj, newObj interface{}) {
				oldRes, ok := oldObj.(*apiv1.Endpoints)
				if !ok {
					log.Errorln("Invalid Endpoint object")
					return
				}
				newRes, ok := newObj.(*apiv1.Endpoints)
				if !ok {
					log.Errorln("Invalid Endpoint object")
					return
				}
				kutil.AssignTypeKind(oldRes)
				kutil.AssignTypeKind(newRes)

				ri := op.ReverseIndex()
				if reflect.DeepEqual(oldRes.Subsets, newRes.Subsets) || ri == nil {
					return
				}

				svc, err := op.KubeClient.CoreV1().Services(newRes.Namespace).Get(newRes.Name, metav1.GetOptions{})
				if err != nil {
					log.Errorln(err)
					return
				}

				oldPods := make(map[string]*apiv1.Pod)
				for _, oldEPSubsets := range oldRes.Subsets {
					for _, oldEPPods := range oldEPSubsets.Addresses {
						if podRef := oldEPPods.TargetRef; podRef != nil {
							pod := &apiv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: podRef.Name, Namespace: podRef.Namespace}}
							oldPods[podRef.String()] = pod
						}
					}
				}

				newPods := make(map[string]*apiv1.Pod)
				for _, newEPSubsets := range newRes.Subsets {
					for _, newEPPods := range newEPSubsets.Addresses {
						if podRef := newEPPods.TargetRef; podRef != nil {
							pod := &apiv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: podRef.Name, Namespace: podRef.Namespace}}
							newPods[podRef.String()] = pod
							if _, ok := oldPods[podRef.String()]; !ok {
								// This Pod reference is in update Endpoint, New Pod Added
								ri.Service.AddPodForService(svc, pod)
							}
						}
					}
				}

				for ref, pod := range oldPods {
					if _, ok := newPods[ref]; !ok {
						// Pod ref not found in New Endpoint, Removed
						ri.Service.DeletePodForService(svc, pod)
					}
				}
			},
		},
	)
	ctrl.Run(wait.NeverStop)
}
