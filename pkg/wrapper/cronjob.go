package wrapper

import (
	"fmt"

	v1 "k8s.io/api/batch/v1beta1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CronJob struct {
	d *v1.CronJob
}

// WrapCronJob wraps a v1.CronJob behind a KubernetesObject interface
func WrapCronJob(i interface{}) (KubernetesObject, error) {
	d, ok := i.(*v1.CronJob)

	if !ok {
		return nil, fmt.Errorf("Expected v1.CronJob received %T", i)
	}

	return &CronJob{
		d: d,
	}, nil
}

func (d *CronJob) GetMetadata() v1meta.ObjectMeta {
	return d.d.ObjectMeta
}

func (d *CronJob) GetObjectSpec() interface{} {
	return d.d.Spec
}

func (d *CronJob) GetType() string {
	return "CronJob"
}
