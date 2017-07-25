package japanese

import (
	"testing"

	"github.com/ikawaha/kagome.ipadic/tokenizer"
	"github.com/stretchr/testify/assert"
	"github.com/tylertreat/BoomFilters"
	"golang.org/x/net/context"
)

var blacklist = []string{
	"死ね",
	"教えて",
	"LINE",
}

func TestNewServe(t *testing.T) {
	s := NewServe(blacklist)
	assert.NotNil(t, s)
	assert.NotNil(t, s.Tokenizer)
	assert.NotNil(t, s.Filer)

	ret, _ := s.Do(context.Background(), "LINE教えて、今すぐに死ね。")
	assert.True(t, ret)

	ret, _ = s.Do(context.Background(), "すもももももももものうち")
	assert.False(t, ret)
}

func TestTokenizer_Tokenize(t *testing.T) {
	tr := &Tokenizer{
		Kagome: tokenizer.NewWithDic(tokenizer.SysDicIPASimple()),
	}
	ts := tr.Tokenize("すもももももももものうち")
	assert.NotEmpty(t, ts)
	assert.Len(t, ts, 4)
}

func TestFilter_Test(t *testing.T) {
	f := &Filter{
		Bloom: boom.NewBloomFilter(uint(len(blacklist)), 0.01),
	}
	for i := range blacklist {
		f.Bloom.Add([]byte(blacklist[i]))
	}
	for i := range blacklist {
		assert.True(t, f.Test([]byte(blacklist[i])))
	}
}
