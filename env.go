package gcws

import "os"

var GOPATH string

func init() {
	GOPATH = os.Getenv("GOPATH")
}
