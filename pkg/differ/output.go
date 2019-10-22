package differ

import (
	"fmt"
)

// Output abstracts a straightforward way to write
type Output interface {
	WriteAdded(name string)
	WriteDeleted(name string)
	WriteUpdated(name string, diffs []string)
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

func (f *output) WriteAdded(name string) {
	fmt.Printf("added : %s\n", name)
}

func (f *output) WriteDeleted(name string) {
	fmt.Printf("deleted : %s\n", name)
}

func (f *output) WriteUpdated(name string, diffs []string) {
	fmt.Printf("updated : %s %v\n", name, diffs)
}
