package japanese

import (
	"github.com/ikawaha/kagome.ipadic/tokenizer"
	"github.com/osamingo/shamoji"
	"github.com/willf/bloom"
	"golang.org/x/text/unicode/norm"
)

type (
	// Tokenizer has kagome tokenizer.
	Tokenizer struct {
		Kagome tokenizer.Tokenizer
	}
	// Filter has bloom filter.
	Filter struct {
		Bloom *bloom.BloomFilter
	}
)

// NewServe generates shamoji.Serve for japanese.
func NewServe(blacklist []string) *shamoji.Serve {
	bf := bloom.NewWithEstimates(uint(len(blacklist)), 0.0001)
	for i := range blacklist {
		bf.AddString(blacklist[i])
	}
	return &shamoji.Serve{
		Tokenizer: &Tokenizer{
			Kagome: tokenizer.NewWithDic(tokenizer.SysDicIPASimple()),
		},
		Filer: &Filter{
			Bloom: bf,
		},
	}
}

// Tokenize implements shamoji.Tokenizer interface.
func (t *Tokenizer) Tokenize(sentence string) [][]byte {

	ts := t.Kagome.Analyze(norm.NFKC.String(sentence), tokenizer.Normal)
	if len(ts) == 0 {
		return nil
	}

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
			case "", "記号", "連体詞", "助動詞", "接頭詞", "助詞", "感動詞", "フィラー":
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

// Test implements shamoji.Filter interface.
func (f *Filter) Test(src []byte) bool {
	return f.Bloom.Test(src)
}
