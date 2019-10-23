package wrapper

import (
	"fmt"

	v1 "k8s.io/api/apps/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type daemonset struct {
	d *v1.DaemonSet
}

// WrapDaemonSet wraps a v1.DaemonSet behind a KubernetesObject interface
func WrapDaemonSet(i interface{}) (KubernetesObject, error) {
	d, ok := i.(*v1.DaemonSet)

	if !ok {
		return nil, fmt.Errorf("Expected v1.DaemonSet received %T", i)
	}

	return &daemonset{
		d: d,
	}, nil
}

func (d *daemonset) GetMetadata() v1meta.ObjectMeta {
	return d.d.ObjectMeta
}

func (d *daemonset) GetObjectSpec() interface{} {
	return d.d.Spec
}

func (d *daemonset) GetType() string {
	return "daemonset"
}
