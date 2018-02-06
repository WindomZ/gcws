package jieba

import (
	"github.com/WindomZ/gcws"
	"github.com/yanyiwu/gojieba"
)

// Jieba is gojieba.Jieba adapter.
type Jieba struct {
	Config gcws.Config
	gojieba.Jieba
}

// NewJieba create new gojieba.Jieba with paths.
func NewJieba(paths ...string) gcws.CWS {
	return &Jieba{
		Config: gcws.DefaultConfig,
		Jieba:  *gojieba.NewJieba(paths...),
	}
}

// Parent is the interface of adapter parent structure.
func (j Jieba) Parent() interface{} {
	return &j.Jieba
}

// SetConfig set cws configuration.
func (j *Jieba) SetConfig(conf gcws.Config) {
	j.Config = conf
}

// Tokenize returns the string array of split.
func (j Jieba) Tokenize(str string) []string {
	return j.ConfigTokenize(j.Config, str)
}

// ConfigTokenize returns the string array of split with cws configuration.
func (j Jieba) ConfigTokenize(conf gcws.Config, str string) (ret []string) {
	if conf.Mode == gcws.ModeSearch {
		ret = j.CutForSearch(str, true)
	} else {
		ret = j.Cut(str, true)
	}
	if conf.SkipPunct {
		ret = gcws.FilterPunct(ret)
	}
	return
}

func init() {
	gcws.Register("jieba", NewJieba)
}
