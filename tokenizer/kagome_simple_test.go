package tokenizer_test

import (
	"testing"

	"github.com/osamingo/shamoji"
	"github.com/osamingo/shamoji/tokenizer"
	"golang.org/x/text/unicode/norm"
)

func TestNewKagomeSimpleTokenizer(t *testing.T) {
	kt := tokenizer.NewKagomeSimpleTokenizer(norm.NFKC)

	if kt == nil {
		t.Error("should not be nil")
	}
	if kt.Form != norm.NFKC {
		t.Error("should be NFKC")
	}
	var i interface{} = kt
	if _, ok := i.(shamoji.Tokenizer); !ok {
		t.Error("should be implements shamoji.Tokenizer")
	}
}

func TestKagomeTokenizer_Tokenize(t *testing.T) {
	kt := tokenizer.NewKagomeSimpleTokenizer(norm.NFKC)

	cases := map[string]struct {
		sentence string
		expect   int
	}{
		"Empty sentence":    {"", 0},
		"Japanese sentence": {"すもももももももものうち", 4},
	}
	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			ts := kt.Tokenize(c.sentence)
			if len(ts) != c.expect {
				t.Error("shound be length", c.expect)
			}
		})
	}
}
