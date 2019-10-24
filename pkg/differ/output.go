package differ

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/apex/log"
)

// Output abstracts a straightforward way to write
type Output interface {
	WriteAdded(name string, namespace string, objectType string)
	WriteDeleted(name string, namespace string, objectType string)
	WriteUpdated(name string, namespace string, objectType string, diffs []string)
}

// OutputFormat encodes
type OutputFormat int

const (
	// Text outputs the diffs in a simple text based format
	Text OutputFormat = iota
	// JSON outputs the diffs in json
	JSON
)

type output struct {
	format     OutputFormat
	logAdded   bool
	logDeleted bool
}

type jsonformat struct {
	Timestamp  string `json:"timestamp"`
	Verb       string `json:"verb"`
	ObjectType string `json:"type"`
	Notes      string `json:"notes"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
}

// NewOutput constructs a new outputter
func NewOutput(fmt OutputFormat, logAdded bool, logDeleted bool) Output {
	return &output{
		format:     fmt,
		logAdded:   logAdded,
		logDeleted: logDeleted,
	}
}

func (f *output) WriteAdded(name string, namespace string, objectType string) {
	if !f.logAdded {
		return
	}

	f.write(name, namespace, "added", objectType, nil)
}

func (f *output) WriteDeleted(name string, namespace string, objectType string) {
	if !f.logDeleted {
		return
	}

	f.write(name, namespace, "deleted", objectType, nil)
}

func (f *output) WriteUpdated(name string, namespace string, objectType string, diffs []string) {
	f.write(name, namespace, "updated", objectType, diffs)
}

func (f *output) write(name string, namespace string, verb string, objectType string, etc []string) {

	switch f.format {
	case Text:
		fmt.Printf("%s %s : %s %s (%s) %v\n", time.Now().UTC().Format(time.RFC3339), verb, namespace, name, objectType, etc)
	case JSON:
		bytes, err := json.Marshal(jsonformat{
			Timestamp:  time.Now().UTC().Format(time.RFC3339),
			Name:       name,
			Namespace:  namespace,
			Verb:       verb,
			ObjectType: objectType,
			Notes:      fmt.Sprintf("%v", etc),
		})

		if err != nil {
			log.Errorf("Failed to convert to json")
			return
		}

		fmt.Println(string(bytes))
	}
}
