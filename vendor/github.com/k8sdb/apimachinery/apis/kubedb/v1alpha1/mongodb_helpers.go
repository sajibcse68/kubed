package v1alpha1

import (
	"fmt"
	"strings"

	"github.com/appscode/kutil/tools/monitoring/api"
	core "k8s.io/api/core/v1"
)

func (p MongoDB) OffshootName() string {
	return p.Name
}

func (p MongoDB) OffshootLabels() map[string]string {
	return map[string]string{
		LabelDatabaseName: p.Name,
		LabelDatabaseKind: ResourceKindMongoDB,
	}
}

func (p MongoDB) StatefulSetLabels() map[string]string {
	labels := p.OffshootLabels()
	for key, val := range p.Labels {
		if !strings.HasPrefix(key, GenericKey+"/") && !strings.HasPrefix(key, MongoDBKey+"/") {
			labels[key] = val
		}
	}
	return labels
}

func (p MongoDB) StatefulSetAnnotations() map[string]string {
	annotations := make(map[string]string)
	for key, val := range p.Annotations {
		if !strings.HasPrefix(key, GenericKey+"/") && !strings.HasPrefix(key, MongoDBKey+"/") {
			annotations[key] = val
		}
	}
	annotations[MongoDBDatabaseVersion] = string(p.Spec.Version)
	return annotations
}

func (p MongoDB) ResourceCode() string {
	return ResourceCodeMongoDB
}

func (p MongoDB) ResourceKind() string {
	return ResourceKindMongoDB
}

func (p MongoDB) ResourceName() string {
	return ResourceNameMongoDB
}

func (p MongoDB) ResourceType() string {
	return ResourceTypeMongoDB
}

func (p MongoDB) ObjectReference() *core.ObjectReference {
	return &core.ObjectReference{
		APIVersion:      SchemeGroupVersion.String(),
		Kind:            p.ResourceKind(),
		Namespace:       p.Namespace,
		Name:            p.Name,
		UID:             p.UID,
		ResourceVersion: p.ResourceVersion,
	}
}

func (p MongoDB) ServiceName() string {
	return p.OffshootName()
}

func (p MongoDB) ServiceMonitorName() string {
	return fmt.Sprintf("kubedb-%s-%s", p.Namespace, p.Name)
}

func (p MongoDB) Path() string {
	return fmt.Sprintf("/kubedb.com/v1alpha1/namespaces/%s/%s/%s/metrics", p.Namespace, p.ResourceType(), p.Name)
}

func (p MongoDB) Scheme() string {
	return ""
}

func (p *MongoDB) StatsAccessor() api.StatsAccessor {
	return p
}
