package wrapper

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubernetesObject presents a consistent way of interacting with various Kubernetes objects
type KubernetesObject interface {
	GetMetadata() v1.ObjectMeta
	GetObjectSpec() interface{}
}

// Wrap accepts an empty interface and returns a KubernetesObject
type Wrap func(interface{}) (KubernetesObject, error)
