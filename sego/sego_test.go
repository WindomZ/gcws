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

func TestSegmenter_SetMode(t *testing.T) {
	demo.SetMode(gcws.ModeDefault)
}

func TestSegmenter_Tokenize(t *testing.T) {
	assert.Equal(t, []string{"喜欢", "就", "坚持", "，", "爱", "就", "别", "放弃"},
		demo.Tokenize("喜欢就坚持，爱就别放弃"))

	demo.SetMode(gcws.ModeSearch)
	assert.Equal(t, []string{"喜欢", "就", "坚持", "，", "爱", "就", "别", "放弃"},
		demo.Tokenize("喜欢就坚持，爱就别放弃"))
}

func TestNewSegmenter(t *testing.T) {
	demo, _ = gcws.NewCWS("sego", "")
	TestSegmenter_Tokenize(t)
}
