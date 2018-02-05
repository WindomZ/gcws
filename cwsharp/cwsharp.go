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
	Mode gcws.ModeType
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
		Mode:      gcws.ModeDefault,
		Tokenizer: token,
	}
}

// SetMode set cws mode type.
func (c *CWSharp) SetMode(typ gcws.ModeType) {
	c.Mode = typ
}

// Tokenize returns the string array of split.
func (c CWSharp) Tokenize(str string) []string {
	return c.ModeTokenize(c.Mode, str)
}

// ModeTokenize returns the string array of split with cws mode type.
func (c CWSharp) ModeTokenize(typ gcws.ModeType, str string) (ret []string) {
	var iter cwsharp.Iterator
	switch typ {
	case gcws.ModeFast:
		iter = cwsharp.BigramTokenize(strings.NewReader(str))
	case gcws.ModeEnglish:
		iter = cwsharp.WhitespaceTokenize(strings.NewReader(str))
	default:
		iter = c.Tokenizer.Tokenize(strings.NewReader(str))
	}
	for tok := iter.Next(); tok != nil; tok = iter.Next() {
		if tok.Type != cwsharp.PUNCT {
			ret = append(ret, tok.Text)
		}
	}
	return
}

func init() {
	gcws.Register("cwsharp", NewCWSharp)
}
