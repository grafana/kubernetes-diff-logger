package differ

import (
	"fmt"

	"github.com/joe-elliott/kubernetes-diff-logger/pkg/wrapper"
)

// Output abstracts a straightforward way to write
type Output interface {
	WriteAdded(added wrapper.KubernetesObject)
	WriteDeleted(deleted wrapper.KubernetesObject)
	WriteUpdated(old wrapper.KubernetesObject, new wrapper.KubernetesObject)
}

// OutputFormat encodes
type OutputFormat int

const (
	// Text outputs the diffs in a simple text based format
	Text OutputFormat = iota
)

type output struct {
	format OutputFormat
}

// NewOutput constructs a new outputter
func NewOutput(fmt OutputFormat) Output {
	return &output{
		format: fmt,
	}
}

func (f *output) WriteAdded(added wrapper.KubernetesObject) {
	fmt.Printf("added : %s\n", added.GetMetadata().Name)
}

func (f *output) WriteDeleted(deleted wrapper.KubernetesObject) {
	fmt.Printf("deleted : %s\n", deleted.GetMetadata().Name)
}

func (f *output) WriteUpdated(old wrapper.KubernetesObject, new wrapper.KubernetesObject) {
	fmt.Printf("updated : %s\n", new.GetMetadata().Name)
}
