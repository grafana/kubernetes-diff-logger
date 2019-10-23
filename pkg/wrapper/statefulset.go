package wrapper

import (
	"fmt"

	v1 "k8s.io/api/apps/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type statefulset struct {
	s *v1.StatefulSet
}

// WrapStatefulSet wraps a v1.WrapStatefulSet behind a KubernetesObject interface
func WrapStatefulSet(i interface{}) (KubernetesObject, error) {
	s, ok := i.(*v1.StatefulSet)

	if !ok {
		return nil, fmt.Errorf("Expected v1.StatefulSet received %T", i)
	}

	return &statefulset{
		s: s,
	}, nil
}

func (s *statefulset) GetMetadata() v1meta.ObjectMeta {
	return s.s.ObjectMeta
}

func (s *statefulset) GetObjectSpec() interface{} {
	return s.s.Spec
}

func (s *statefulset) GetType() string {
	return "statefulset"
}
