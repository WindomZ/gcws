package sego

import (
	"testing"

	"github.com/WindomZ/gcws"
	"github.com/WindomZ/testify/assert"
)

var demo gcws.CWS

func init() {
	demo, _ = gcws.NewCWS("sego")
}

func TestRegister(t *testing.T) {
	defer func() {
		assert.NotEmpty(t, recover())
	}()
	gcws.Register("sego", NewSegmenter)
}

func TestSegmenter_SetConfig(t *testing.T) {
	demo.SetConfig(gcws.Config{
		Mode: gcws.ModeSearch,
	})
	demo.SetConfig(gcws.DefaultConfig)
}
func TestSegmenter_Tokenize(t *testing.T) {
	assert.Equal(t, []string{"喜欢", "就", "坚持", "，", "爱", "就", "别", "放弃"},
		demo.Tokenize("喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"喜欢", "就", "坚持", "爱", "就", "别", "放弃"},
		demo.ConfigTokenize(gcws.Config{
			Mode:      gcws.ModeSearch,
			SkipPunct: true,
		}, "喜欢就坚持，爱就别放弃"))
}

func TestNewSegmenter(t *testing.T) {
	demo, _ = gcws.NewCWS("sego", "")
	TestSegmenter_Tokenize(t)
}
