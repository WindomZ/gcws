package cwsharp

import (
	"testing"

	"github.com/WindomZ/gcws"
	"github.com/WindomZ/testify/assert"
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

func TestCWSharp_SetMode(t *testing.T) {
	demo.SetMode(gcws.ModeSearch)
	demo.SetMode(gcws.ModeDefault)
}

func TestCWSharp_Tokenize(t *testing.T) {
	assert.Equal(t, []string{"喜欢", "就", "坚持", "爱", "就", "别", "放弃"},
		demo.Tokenize("喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"喜欢", "就", "坚持", "爱", "就", "别", "放弃"},
		demo.ModeTokenize(gcws.ModeSearch, "喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"喜欢", "欢就", "就坚", "坚持", "爱就", "就别", "别放", "放弃"},
		demo.ModeTokenize(gcws.ModeFast, "喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"for", "man", "is", "man", "and", "master", "of", "his", "fate"},
		demo.ModeTokenize(gcws.ModeEnglish, "For man is man and master of his fate."))
}
