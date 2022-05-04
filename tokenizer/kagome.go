package tokenizer

import (
	"fmt"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"golang.org/x/text/unicode/norm"
)

// KagomeTokenizer tokenize by kagome.
type KagomeTokenizer struct {
	Form   norm.Form
	Kagome *tokenizer.Tokenizer
}

// Tokenize implements shamoji.Tokenizer interface.
func (kt *KagomeTokenizer) Tokenize(sentence string) [][]byte {
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

			switch ts[i].Features()[0] {
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

// NewKagomeTokenizer generates new KagomeTokenizer.
func NewKagomeTokenizer(f norm.Form) (*KagomeTokenizer, error) {
	k, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, fmt.Errorf("tokenizer: failed to generate tokenizer: %w", err)
	}

	return &KagomeTokenizer{
		Form:   f,
		Kagome: k,
	}, nil
}
