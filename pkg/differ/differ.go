package differ

import (
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"

	"github.com/go-test/deep"
	"github.com/joe-elliott/kubernetes-diff-logger/pkg/wrapper"
	"github.com/ryanuber/go-glob"
)

// Differ is responsible for subscribing to an informer an filtering out events
type Differ struct {
	matchGlob string
	wrap      wrapper.Wrap
	informer  cache.SharedInformer
	output    Output
}

// NewDiffer constructs a Differ
func NewDiffer(m string, f wrapper.Wrap, i cache.SharedInformer, o Output) *Differ {
	d := &Differ{
		matchGlob: m,
		wrap:      f,
		informer:  i,
		output:    o,
	}

	return d
}

// Run sets up eventhandlers, sync informer caches and blocks until stop is closed
func (d *Differ) Run(stopCh <-chan struct{}) error {
	defer runtime.HandleCrash()

	if ok := cache.WaitForCacheSync(stopCh, d.informer.HasSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	d.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    d.added,
		UpdateFunc: d.updated,
		DeleteFunc: d.deleted,
	})

	<-stopCh

	return nil
}

func (d *Differ) added(added interface{}) {
	object := d.mustWrap(added)

	if d.matches(object) {
		d.output.WriteAdded(object.GetMetadata().Name)
	}
}

func (d *Differ) updated(old interface{}, new interface{}) {
	oldObject := d.mustWrap(old)
	newObject := d.mustWrap(new)

	if d.matches(oldObject) || d.matches(newObject) {
		if diff := deep.Equal(oldObject.GetObject(), newObject.GetObject()); diff != nil {
			d.output.WriteUpdated(newObject.GetMetadata().Name, diff)
		}
	}
}

func (d *Differ) deleted(deleted interface{}) {
	object := d.mustWrap(deleted)

	if d.matches(object) {
		d.output.WriteDeleted(object.GetMetadata().Name)
	}
}

func (d *Differ) matches(o wrapper.KubernetesObject) bool {
	return glob.Glob(d.matchGlob, o.GetMetadata().Name)
}

func (d *Differ) mustWrap(i interface{}) wrapper.KubernetesObject {
	o, err := d.wrap(i)

	if err != nil {
		log.Fatalf("Failed to wrap interface %v", err)
	}

	return o
}
