package segment

import (
	"testing"

	"github.com/WindomZ/gcws"
	"github.com/WindomZ/gosegment"
	"github.com/WindomZ/testify/assert"
)

var demo gcws.CWS

func init() {
	demo, _ = gcws.NewCWS("segment")
}

func TestRegister(t *testing.T) {
	defer func() {
		assert.NotEmpty(t, recover())
	}()
	gcws.Register("segment", NewSegment)
}

func TestSegment_Parent(t *testing.T) {
	assert.NotEmpty(t, demo.Parent())

	_, ok := demo.Parent().(*gosegment.Segment)
	assert.True(t, ok)
}

func TestSegment_SetConfig(t *testing.T) {
	demo.SetConfig(gcws.Config{
		Mode: gcws.ModeSearch,
	})
	demo.SetConfig(gcws.DefaultConfig)
}

func TestSegment_Tokenize(t *testing.T) {
	assert.Equal(t, []string{"喜欢", "就", "坚持", "，", "爱就别", "放弃"},
		demo.Tokenize("喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"喜欢", "就", "坚持", "，", "爱就别", "放弃"},
		demo.ConfigTokenize(gcws.Config{
			Mode: gcws.ModeSearch,
		}, "喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"喜欢", "就", "坚持", "爱就别", "放弃"},
		demo.ConfigTokenize(gcws.Config{
			Mode:            gcws.ModeSearch,
			FilterStopWords: true,
		}, "喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"For", "man", "is", "man", "and", "master", "of", "his", "fate"},
		demo.ConfigTokenize(gcws.Config{
			Mode:            gcws.ModeEnglish,
			FilterStopWords: true,
		}, "For man is man and master of his fate."))
}

func TestNewSegment(t *testing.T) {
	demo, _ = gcws.NewCWS("segment", "")
	TestSegment_Tokenize(t)
}
