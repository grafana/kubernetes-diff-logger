package differ

import (
	"fmt"
	"time"
)

// Output abstracts a straightforward way to write
type Output interface {
	WriteAdded(name string, objectType string)
	WriteDeleted(name string, objectType string)
	WriteUpdated(name string, objectType string, diffs []string)
}

// OutputFormat encodes
type OutputFormat int

const (
	// Text outputs the diffs in a simple text based format
	Text OutputFormat = iota
)

type output struct {
	format     OutputFormat
	logAdded   bool
	logDeleted bool
}

// NewOutput constructs a new outputter
func NewOutput(fmt OutputFormat, logAdded bool, logDeleted bool) Output {
	return &output{
		format:     fmt,
		logAdded:   logAdded,
		logDeleted: logDeleted,
	}
}

func (f *output) WriteAdded(name string, objectType string) {
	if !f.logAdded {
		return
	}

	f.write(name, "added", objectType, nil)
}

func (f *output) WriteDeleted(name string, objectType string) {
	if !f.logDeleted {
		return
	}

	f.write(name, "deleted", objectType, nil)
}

func (f *output) WriteUpdated(name string, objectType string, diffs []string) {
	f.write(name, "updated", objectType, diffs)
}

func (f *output) write(name string, verb string, objectType string, etc interface{}) {
	fmt.Printf("%s %s (%s) %s : %s %v\n", time.Now().UTC().Format(time.RFC3339), name, objectType, verb, name, etc)
}
