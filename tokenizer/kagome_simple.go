package tokenizer

import (
	"github.com/ikawaha/kagome.ipadic/tokenizer"
	"golang.org/x/text/unicode/norm"
)

// KagomeSimpleTokenizer tokenize by kagome.
type KagomeSimpleTokenizer struct {
	Form   norm.Form
	Kagome tokenizer.Tokenizer
}

// Tokenize implements shamoji.Tokenizer interface.
func (kt *KagomeSimpleTokenizer) Tokenize(sentence string) [][]byte {

	ts := kt.Kagome.Analyze(kt.Form.String(sentence), tokenizer.Normal)
	ch := make(chan []byte, len(ts))

	for i := range ts {
		i := i
		go func() {
			var s []byte
			defer func() {
				ch <- s
			}()
			if ts[i].Class == tokenizer.DUMMY {
				return
			}
			switch ts[i].Pos() {
			case "", "連体詞", "接続詞", "助詞", "助動詞", "記号", "フィラー", "その他":
				return
			default:
				s = []byte(ts[i].Surface)
			}
		}()
	}

	ret := make([][]byte, 0, len(ts))
	for range ts {
		if t := <-ch; t != nil {
			ret = append(ret, t)
		}
	}

	return ret
}

// NewKagomeSimpleTokenizer generates new KagomeSimpleTokenizer.
func NewKagomeSimpleTokenizer(f norm.Form) *KagomeSimpleTokenizer {
	return &KagomeSimpleTokenizer{
		Form:   f,
		Kagome: tokenizer.NewWithDic(tokenizer.SysDicIPASimple()),
	}
}
