package differ

import (
	"fmt"
	"time"

	"k8s.io/client-go/tools/cache"

	"github.com/joe-elliott/kubernetes-diff-logger/pkg/wrapper"
)

// Differ is responsible for subscribing to an informer an filtering out events
type Differ struct {
	name      string
	matchGlob string
	wrap      wrapper.Wrap
}

// NewDiffer constructs a Differ
func NewDiffer(n string, m string, r time.Duration, f wrapper.Wrap, i cache.SharedInformer) *Differ {
	d := &Differ{
		matchGlob: m,
		name:      n,
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
	o, _ := d.wrap(added)

	fmt.Println(o)
}

func (d *Differ) updated(old interface{}, new interface{}) {
	o, _ := d.wrap(old)
	n, _ := d.wrap(new)

	fmt.Println(o)
	fmt.Println(n)
}

func (d *Differ) deleted(deleted interface{}) {
	o, _ := d.wrap(deleted)

	fmt.Println(o)
}
