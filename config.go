package main

// Config represents the configuration options for kubernetes-diff-logger
type Config struct {
	Differs []DifferConfig `yaml:"differs"`
}

// DifferConfig represents the configuration options for a single Diffing process
type DifferConfig struct {
	// NameFilter is a glob-based filter for the object type
	NameFilter string `yaml:"nameFilter"`
	// Type specifies the Kubernetes object type to watch.  Currently supporting statefulsets, daemonsets, deployments
	Type string `yaml:"type"`
}

// DefaultConfig returns a default deployment watching config
func DefaultConfig() Config {
	return Config{
		Differs: []DifferConfig{
			DifferConfig{
				NameFilter: "*",
				Type:       "deployment",
			},
		},
	}
}
