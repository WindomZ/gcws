package gse

import (
	"github.com/WindomZ/gcws"
	"github.com/go-ego/gse"
)

// Segmenter is gse.Segmenter adapter.
type Segmenter struct {
	Config gcws.Config
	gse.Segmenter
}

// NewSegmenter create new gse.Segmenter with paths.
func NewSegmenter(paths ...string) gcws.CWS {
	segmenter := &Segmenter{
		Config: gcws.DefaultConfig,
	}
	segmenter.LoadDict(paths...)
	return segmenter
}

// Parent is the interface of adapter parent structure.
func (c Segmenter) Parent() interface{} {
	return &c.Segmenter
}

// SetConfig set cws configuration.
func (c *Segmenter) SetConfig(conf gcws.Config) {
	c.Config = conf
}

// Tokenize returns the string array of split.
func (c Segmenter) Tokenize(str string) []string {
	return c.ConfigTokenize(c.Config, str)
}

// ConfigTokenize returns the string array of split with cws configuration.
func (c Segmenter) ConfigTokenize(conf gcws.Config, str string) (ret []string) {
	switch conf.Mode {
	case gcws.ModeSearch:
		ret = gse.ToSlice(c.ModeSegment([]byte(str), true), true)
	default:
		ret = gse.ToSlice(c.ModeSegment([]byte(str), false), false)
	}

	if conf.FilterStopWords {
		ret = gcws.FilterPunct(ret)
	}
	return
}

func init() {
	gcws.Register("gse", NewSegmenter)
}
