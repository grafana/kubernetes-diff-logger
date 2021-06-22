package wrapper

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type pod struct {
	p *v1.Pod
}

// WrapPod wraps a v1.Pod behind a KubernetesObject interface
func WrapPod(i interface{}) (KubernetesObject, error) {
	p, ok := i.(*v1.Pod)

	if !ok {
		return nil, fmt.Errorf("Expected v1.Pod received %T", i)
	}

	return &pod{
		p: p,
	}, nil
}

func (p *pod) GetMetadata() v1meta.ObjectMeta {
	return p.p.ObjectMeta
}

func (p *pod) GetObjectSpec() interface{} {
	return p.p.Spec
}

func (p *pod) GetType() string {
	return "pod"
}
