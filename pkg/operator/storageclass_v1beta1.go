package operator

import (
	"errors"
	"reflect"

	"github.com/appscode/go/log"
	acrt "github.com/appscode/go/runtime"
	"github.com/appscode/kubed/pkg/util"
	kutil "github.com/appscode/kutil/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	storage "k8s.io/client-go/pkg/apis/storage/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// Blocks caller. Intended to be called as a Go routine.
func (op *Operator) WatchStorageClassV1beta1() {
	if !util.IsPreferredAPIResource(op.KubeClient, storage.SchemeGroupVersion.String(), "StorageClass") {
		log.Warningf("Skipping watching non-preferred GroupVersion:%s Kind:%s", storage.SchemeGroupVersion.String(), "StorageClass")
		return
	}

	defer acrt.HandleCrash()

	lw := &cache.ListWatch{
		ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
			return op.KubeClient.StorageV1beta1().StorageClasses().List(metav1.ListOptions{})
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return op.KubeClient.StorageV1beta1().StorageClasses().Watch(metav1.ListOptions{})
		},
	}
	_, ctrl := cache.NewInformer(lw,
		&storage.StorageClass{},
		op.Opt.ResyncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				if res, ok := obj.(*storage.StorageClass); ok {
					log.Infof("StorageClass %s@%s added", res.Name, res.Namespace)
					kutil.AssignTypeKind(res)

					si := op.SearchIndex()
					if si != nil {
						if err := si.HandleAdd(obj); err != nil {
							log.Errorln(err)
						}
					}

					if op.eventer != nil &&
						op.Config.EventForwarder.StorageAdded.Handle &&
						util.IsRecent(res.ObjectMeta.CreationTimestamp) {
						err := op.eventer.Forward(res.TypeMeta, res.ObjectMeta, "added", obj)
						if err != nil {
							log.Errorln(err)
						}
					}
				}
			},
			DeleteFunc: func(obj interface{}) {
				if res, ok := obj.(*storage.StorageClass); ok {
					log.Infof("StorageClass %s@%s deleted", res.Name, res.Namespace)
					kutil.AssignTypeKind(res)

					si := op.SearchIndex()
					if si != nil {
						if err := si.HandleDelete(obj); err != nil {
							log.Errorln(err)
						}
					}
					tc := op.TrashCan()
					if tc != nil {
						tc.Delete(res.TypeMeta, res.ObjectMeta, obj)
					}
				}
			},
			UpdateFunc: func(old, new interface{}) {
				oldRes, ok := old.(*storage.StorageClass)
				if !ok {
					log.Errorln(errors.New("Invalid StorageClass object"))
					return
				}
				newRes, ok := new.(*storage.StorageClass)
				if !ok {
					log.Errorln(errors.New("Invalid StorageClass object"))
					return
				}
				kutil.AssignTypeKind(oldRes)
				kutil.AssignTypeKind(newRes)

				si := op.SearchIndex()
				if si != nil {
					si.HandleUpdate(old, new)
				}
				tc := op.TrashCan()
				if tc != nil && op.Config.RecycleBin.HandleUpdates {
					if !reflect.DeepEqual(oldRes.Labels, newRes.Labels) ||
						!reflect.DeepEqual(oldRes.Annotations, newRes.Annotations) ||
						!reflect.DeepEqual(oldRes.Parameters, newRes.Parameters) {
						tc.Update(newRes.TypeMeta, newRes.ObjectMeta, old, new)
					}
				}
			},
		},
	)
	ctrl.Run(wait.NeverStop)
}
