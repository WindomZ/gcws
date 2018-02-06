package gcws

import "fmt"

// CWS short for Chinese Word Segmentation,
// interface contains all behaviors for cws adapter.
type CWS interface {
	// Parent is the interface of adapter parent structure.
	Parent() interface{}
	// SetConfig set cws configuration.
	SetConfig(Config)
	// Tokenize returns the string array of split.
	Tokenize(string) []string
	// ConfigTokenize returns the string array of split with cws configuration.
	ConfigTokenize(Config, string) []string
}

// Instance is a function create a new CWS Instance
type Instance func(paths ...string) CWS

var adapters = make(map[string]Instance)

// Register makes a cws adapter available by the adapter name.
// If Register is called twice with the same name or if driver is nil, it panics.
func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("cws: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("cws: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

// NewCWS Create a new cws driver by adapter name and dictionary paths.
// paths are file paths of the dictionary.
func NewCWS(name string, paths ...string) (adapter CWS, err error) {
	f, ok := adapters[name]
	if !ok {
		err = fmt.Errorf("cws: unknown adapter name %q (forgot to import?)", name)
		return
	}
	adapter = f(paths...)
	return
}
