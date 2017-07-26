package tokenizer

import (
	"testing"

	"github.com/osamingo/shamoji"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/unicode/norm"
)

func TestNewKagomeKagomeTokenizer(t *testing.T) {

	kt := NewKagomeSimpleTokenizer(norm.NFKC)

	assert.NotNil(t, kt)
	assert.Equal(t, kt.Form, norm.NFKC)
	assert.NotNil(t, kt.Kagome)
	assert.Implements(t, (*shamoji.Tokenizer)(nil), kt)
}

func TestKagomeTokenizer_Tokenize(t *testing.T) {

	kt := NewKagomeSimpleTokenizer(norm.NFKC)

	ts := kt.Tokenize("")
	assert.Empty(t, ts)

	ts = kt.Tokenize("すもももももももものうち")
	assert.NotEmpty(t, ts)
	assert.Len(t, ts, 4)
}
