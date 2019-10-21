package differ

import (
	"fmt"
	"log"
	"time"

	"k8s.io/client-go/tools/cache"

	"github.com/joe-elliott/kubernetes-diff-logger/pkg/wrapper"
	"github.com/ryanuber/go-glob"
)

// Differ is responsible for subscribing to an informer an filtering out events
type Differ struct {
	matchGlob string
	wrap      wrapper.Wrap
}

// NewDiffer constructs a Differ
func NewDiffer(m string, r time.Duration, f wrapper.Wrap, i cache.SharedInformer) *Differ {
	d := &Differ{
		matchGlob: m,
		wrap:      f,
	}

	i.AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
		AddFunc:    d.added,
		UpdateFunc: d.updated,
		DeleteFunc: d.deleted,
	}, r)

	return d
}

func (d *Differ) added(added interface{}) {
	object := d.mustWrap(added)

	if d.matches(object) {
		fmt.Printf("added: %s\n", object.GetMetadata().Name)
	}
}

func (d *Differ) updated(old interface{}, new interface{}) {
	oldObject := d.mustWrap(old)
	newObject := d.mustWrap(new)

	if d.matches(oldObject) ||
	   d.matches(newObject) {
		fmt.Printf("updated: %s\n", object.GetMetadata().Name)
	}
}

func (d *Differ) deleted(deleted interface{}) {
	object := d.mustWrap(deleted)

	if d.matches(object) {
		fmt.Printf("deleted: %s\n", object.GetMetadata().Name)
	}
}

func (d *Differ) matches(o wrapper.KubernetesObject) bool {
	return glob.Glob(d.matchGlob, o.GetMetadata().Name)
}

func (d *Differ) mustWrap(i interface{}) wrapper.KubernetesObject {
	o, err := d.wrap(i)

	if err != nil {
		log.Fatalf("Failed to wrap interface %v", o)
	}

	return o
}
