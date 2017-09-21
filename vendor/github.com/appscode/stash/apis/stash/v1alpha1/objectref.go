package v1alpha1

import (
	apiv1 "k8s.io/client-go/pkg/api/v1"
)

func (r Restic) ObjectReference() *apiv1.ObjectReference {
	return &apiv1.ObjectReference{
		APIVersion:      SchemeGroupVersion.String(),
		Kind:            ResourceKindRestic,
		Namespace:       r.Namespace,
		Name:            r.Name,
		UID:             r.UID,
		ResourceVersion: r.ResourceVersion,
	}
}
