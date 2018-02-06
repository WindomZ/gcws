package jieba

import (
	"testing"

	"github.com/WindomZ/gcws"
	"github.com/WindomZ/testify/assert"
	"github.com/yanyiwu/gojieba"
)

var demo gcws.CWS

func init() {
	demo, _ = gcws.NewCWS("jieba")
}

func TestRegister(t *testing.T) {
	defer func() {
		assert.NotEmpty(t, recover())
	}()
	gcws.Register("jieba", NewJieba)
}

func TestJieba_Parent(t *testing.T) {
	assert.NotEmpty(t, demo.Parent())

	_, ok := demo.Parent().(*gojieba.Jieba)
	assert.True(t, ok)
}

func TestJieba_SetConfig(t *testing.T) {
	demo.SetConfig(gcws.Config{
		Mode: gcws.ModeSearch,
	})
	demo.SetConfig(gcws.DefaultConfig)
}

func TestJieba_Tokenize(t *testing.T) {
	assert.Equal(t, []string{"喜欢", "就", "坚持", "，", "爱", "就", "别", "放弃"},
		demo.Tokenize("喜欢就坚持，爱就别放弃"))

	assert.Equal(t, []string{"喜欢", "就", "坚持", "爱", "就", "别", "放弃"},
		demo.ConfigTokenize(gcws.Config{
			Mode:      gcws.ModeSearch,
			SkipPunct: true,
		}, "喜欢就坚持，爱就别放弃"))
}
