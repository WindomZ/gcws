package gcws

import "os"

// GOPATH the gopath system env directory.
var GOPATH string

func init() {
	GOPATH = os.Getenv("GOPATH")
}
