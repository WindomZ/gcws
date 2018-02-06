package cwsharp

import (
	"testing"

	"github.com/WindomZ/gcws"
	"github.com/WindomZ/testify/assert"
	"github.com/zhengchun/cwsharp-go"
)

var demo gcws.CWS

func init() {
	demo, _ = gcws.NewCWS("cwsharp")
}

func TestRegister(t *testing.T) {
	defer func() {
		assert.NotEmpty(t, recover())
	}()
	gcws.Register("cwsharp", NewCWSharp)
}

func TestCWSharp_Parent(t *testing.T) {
	assert.NotEmpty(t, demo.Parent())

	_, ok := demo.Parent().(cwsharp.Tokenizer)
	assert.True(t, ok)
}

func TestCWSharp_SetConfig(t *testing.T) {
	demo.SetConfig(gcws.Config{
		Mode: gcws.ModeSearch,
	})
	demo.SetConfig(gcws.DefaultConfig)
}

func TestCWSharp_Tokenize(t *testing.T) {
	assert.Equal(t, []string{"喜欢", "就", "坚持", ",", "爱", "就", "别", "放弃"},
		demo.Tokenize("喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"喜欢", "就", "坚持", ",", "爱", "就", "别", "放弃"},
		demo.ConfigTokenize(gcws.Config{
			Mode: gcws.ModeSearch,
		}, "喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"喜欢", "欢就", "就坚", "坚持", "爱就", "就别", "别放", "放弃"},
		demo.ConfigTokenize(gcws.Config{
			Mode:      gcws.ModeFast,
			SkipPunct: true,
		}, "喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"for", "man", "is", "man", "and", "master", "of", "his", "fate"},
		demo.ConfigTokenize(gcws.Config{
			Mode:      gcws.ModeEnglish,
			SkipPunct: true,
		}, "For man is man and master of his fate."))
}

func TestNewCWSharp(t *testing.T) {
	demo, _ = gcws.NewCWS("cwsharp", "")
	TestCWSharp_Tokenize(t)
}
