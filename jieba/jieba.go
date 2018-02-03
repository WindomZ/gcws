package jieba

import (
	"github.com/WindomZ/gcws"
	"github.com/yanyiwu/gojieba"
)

// Jieba is gojieba.Jieba adapter.
type Jieba struct {
	Mode gcws.ModeType
	gojieba.Jieba
}

// NewJieba create new gojieba.Jieba with paths.
func NewJieba(paths ...string) gcws.CWS {
	cws := &Jieba{
		Mode:  gcws.ModeDefault,
		Jieba: *gojieba.NewJieba(paths...),
	}
	return cws
}

// SetMode set cws mode type.
func (j *Jieba) SetMode(typ gcws.ModeType) {
	j.Mode = typ
}

// Tokenize returns the string array of split.
func (j Jieba) Tokenize(str string) []string {
	return j.ModeTokenize(j.Mode, str)
}

// ModeTokenize returns the string array of split with cws mode type.
func (j Jieba) ModeTokenize(typ gcws.ModeType, str string) []string {
	if typ == gcws.ModeDefault {
		return j.Cut(str, true)
	}
	return j.CutForSearch(str, true)
}

func init() {
	gcws.Register("jieba", NewJieba)
}
