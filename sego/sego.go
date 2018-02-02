package sego

import (
	"os"
	"path"

	"github.com/WindomZ/gcws"
	"github.com/huichen/sego"
)

// Segmenter is sego.Segmenter adapter.
type Segmenter struct {
	Mode gcws.ModeType
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
	segmenter := new(Segmenter)
	segmenter.LoadDictionary(filePath)
	return segmenter
}

// SetMode set cws mode type.
func (s *Segmenter) SetMode(typ gcws.ModeType) {
	s.Mode = typ
}

// Tokenize returns the string array of split.
func (s Segmenter) Tokenize(str string) []string {
	return s.ModeTokenize(s.Mode, str)
}

// ModeTokenize returns the string array of split with cws mode type.
func (s Segmenter) ModeTokenize(typ gcws.ModeType, str string) []string {
	if typ == gcws.ModeDefault {
		return sego.SegmentsToSlice(s.Segment([]byte(str)), false)
	}
	return sego.SegmentsToSlice(s.Segment([]byte(str)), true)
}

func init() {
	gcws.Register("sego", NewSegmenter)
}
