package wrapper

import (
	"fmt"

	v1 "k8s.io/api/apps/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type deployment struct {
	d *v1.Deployment
}

// WrapDeployment wraps a v1.Deployment behind a KubernetesObject interface
func WrapDeployment(i interface{}) (KubernetesObject, error) {
	d, ok := i.(*v1.Deployment)

	if !ok {
		return nil, fmt.Errorf("Expected v1.Deployment received %T", i)
	}

	return &deployment{
		d: d,
	}, nil
}

func (d *deployment) GetMetadata() v1meta.ObjectMeta {
	return d.d.ObjectMeta
}

func (d *deployment) GetObjectSpec() interface{} {
	return d.d.Spec
}

func (d *deployment) GetType() string {
	return "deployment"
}
