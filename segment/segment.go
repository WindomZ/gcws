package segment

import (
	"container/list"

	"github.com/WindomZ/gcws"
	"github.com/WindomZ/gosegment"
	"github.com/WindomZ/gosegment/dict"
	"github.com/WindomZ/gosegment/match"
)

// Segment is segment.Segment adapter.
type Segment struct {
	Config gcws.Config
	gosegment.Segment
}

// NewSegment create new segment.Segment with paths.
func NewSegment(paths ...string) gcws.CWS {
	return &Segment{
		Config:  gcws.DefaultConfig,
		Segment: *gosegment.NewSegment(paths...),
	}
}

// Parent is the interface of adapter parent structure.
func (c Segment) Parent() interface{} {
	return &c.Segment
}

// SetConfig set cws configuration.
func (c *Segment) SetConfig(conf gcws.Config) {
	c.Config = conf
}

// Tokenize returns the string array of split.
func (c Segment) Tokenize(str string) []string {
	return c.ConfigTokenize(c.Config, str)
}

var (
	optDefault *match.MatchOptions
	optSearch  *match.MatchOptions
	optEnglish *match.MatchOptions
)

func init() {
	optDefault = match.NewMatchOptions()
	optSearch = match.NewMatchOptions()
	optSearch.FrequencyFirst = true
	optSearch.MultiDimensionality = true
	optEnglish = match.NewMatchOptions()
	optEnglish.EnglishMultiDimensionality = true
}

// ConfigTokenize returns the string array of split with cws configuration.
func (c Segment) ConfigTokenize(conf gcws.Config, str string) (ret []string) {
	var iter *list.List

	switch conf.Mode {
	case gcws.ModeSearch:
		optSearch.FilterStopWords = conf.FilterStopWords
		iter = c.Segment.DoSegmentWithOption(str, optSearch)
	case gcws.ModeEnglish:
		optEnglish.FilterStopWords = conf.FilterStopWords
		iter = c.Segment.DoSegmentWithOption(str, optEnglish)
	default:
		optDefault.FilterStopWords = conf.FilterStopWords
		iter = c.Segment.DoSegmentWithOption(str, optDefault)
	}

	for cur := iter.Front(); cur != nil; cur = cur.Next() {
		if w, ok := cur.Value.(*dict.WordInfo); ok {
			ret = append(ret, w.Word)
		}
	}

	return
}

func init() {
	gcws.Register("segment", NewSegment)
}
