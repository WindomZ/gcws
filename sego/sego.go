package sego

import (
	"os"
	"path"

	"github.com/WindomZ/gcws"
	"github.com/huichen/sego"
)

// Segmenter is sego.Segmenter adapter.
type Segmenter struct {
	Config gcws.Config
	sego.Segmenter
}

// NewSegmenter create new sego.Segmenter with paths.
func NewSegmenter(paths ...string) gcws.CWS {
	var filePath string
	if len(paths) != 0 {
		filePath = paths[0]
	}
	if filePath == "" {
		filePath = path.Join(os.Getenv("GOPATH"),
			"src/github.com/huichen/sego/data/dictionary.txt")
	}
	segmenter := &Segmenter{
		Config: gcws.DefaultConfig,
	}
	segmenter.LoadDictionary(filePath)
	return segmenter
}

// SetConfig set cws configuration.
func (s *Segmenter) SetConfig(conf gcws.Config) {
	s.Config = conf
}

// Tokenize returns the string array of split.
func (s Segmenter) Tokenize(str string) []string {
	return s.ConfigTokenize(s.Config, str)
}

// ConfigTokenize returns the string array of split with cws configuration.
func (s Segmenter) ConfigTokenize(conf gcws.Config, str string) (ret []string) {
	if conf.Mode == gcws.ModeSearch {
		ret = sego.SegmentsToSlice(s.Segment([]byte(str)), true)
	} else {
		ret = sego.SegmentsToSlice(s.Segment([]byte(str)), false)
	}
	if conf.SkipPunct {
		ret = gcws.FilterPunct(ret)
	}
	return
}

func init() {
	gcws.Register("sego", NewSegmenter)
}
