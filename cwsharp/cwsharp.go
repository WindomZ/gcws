package cwsharp

import (
	"os"
	"path"
	"strings"

	"github.com/WindomZ/gcws"
	"github.com/zhengchun/cwsharp-go"
)

// CWSharp is cwsharp.CWSharp adapter.
type CWSharp struct {
	Config gcws.Config
	cwsharp.Tokenizer
}

// NewCWSharp create new cwsharp.CWSharp with paths.
func NewCWSharp(paths ...string) gcws.CWS {
	var filePath string
	if len(paths) != 0 {
		filePath = paths[0]
	}
	if filePath == "" {
		filePath = path.Join(os.Getenv("GOPATH"),
			"src/github.com/zhengchun/cwsharp-go/data/cwsharp.dawg")
	}
	token, err := cwsharp.New(filePath)
	if err != nil {
		panic(err)
	}
	return &CWSharp{
		Config:    gcws.DefaultConfig,
		Tokenizer: token,
	}
}

// Parent is the interface of adapter parent structure.
func (c CWSharp) Parent() interface{} {
	return c.Tokenizer
}

// SetConfig set cws configuration.
func (c *CWSharp) SetConfig(conf gcws.Config) {
	c.Config = conf
}

// Tokenize returns the string array of split.
func (c CWSharp) Tokenize(str string) []string {
	return c.ConfigTokenize(c.Config, str)
}

// ConfigTokenize returns the string array of split with cws configuration.
func (c CWSharp) ConfigTokenize(conf gcws.Config, str string) (ret []string) {
	var iter cwsharp.Iterator

	switch conf.Mode {
	case gcws.ModeFast:
		iter = cwsharp.BigramTokenize(strings.NewReader(str))
	case gcws.ModeEnglish:
		iter = cwsharp.WhitespaceTokenize(strings.NewReader(str))
	default:
		iter = c.Tokenizer.Tokenize(strings.NewReader(str))
	}

	if conf.SkipPunct {
		for tok := iter.Next(); tok != nil; tok = iter.Next() {
			if tok.Type != cwsharp.PUNCT {
				ret = append(ret, tok.Text)
			}
		}
	} else {
		for tok := iter.Next(); tok != nil; tok = iter.Next() {
			ret = append(ret, tok.Text)
		}
	}

	return
}

func init() {
	gcws.Register("cwsharp", NewCWSharp)
}
