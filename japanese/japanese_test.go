package japanese

import (
	"testing"

	"github.com/ikawaha/kagome.ipadic/tokenizer"
	"github.com/irfansharif/cfilter"
	"github.com/stretchr/testify/assert"
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

	ret, _ := s.Do("LINE教えて、今すぐに死ね。")
	assert.True(t, ret)

	ret, _ = s.Do("すもももももももものうち")
	assert.False(t, ret)
}

func TestTokenizer_Tokenize(t *testing.T) {
	tr := &Tokenizer{
		Kagome: tokenizer.NewWithDic(tokenizer.SysDicIPASimple()),
	}

	ts := tr.Tokenize("")
	assert.Empty(t, ts)

	ts = tr.Tokenize("すもももももももものうち")
	assert.NotEmpty(t, ts)
	assert.Len(t, ts, 4)
}

func TestFilter_Test(t *testing.T) {
	f := &Filter{
		Cuckoo: cfilter.New(cfilter.Size(uint(len(blacklist)))),
	}
	for i := range blacklist {
		f.Cuckoo.Insert([]byte(blacklist[i]))
	}
	for i := range blacklist {
		assert.True(t, f.Test([]byte(blacklist[i])))
	}
}
