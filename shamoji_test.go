package shamoji

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

type Example struct {
	Blacklist [][]byte
}

func (e *Example) Tokenize(sentence string) [][]byte {
	fs := strings.Fields(sentence)
	ts := make([][]byte, len(fs))
	for i := range fs {
		ts[i] = []byte(fs[i])
	}
	return ts
}

func (e *Example) Test(src []byte) bool {
	for i := range e.Blacklist {
		if string(src) == string(e.Blacklist[i]) {
			return true
		}
	}
	return false
}

func TestServe_Do(t *testing.T) {
	e := &Example{
		Blacklist: [][]byte{
			[]byte("fuck"),
			[]byte("fucker"),
		},
	}
	s := &Serve{
		Tokenizer: e,
		Filer:     e,
	}

	ret, token := s.Do("")
	assert.False(t, ret)
	assert.Empty(t, token)

	ret, token = s.Do("fuck you.")
	assert.True(t, ret)
	assert.Equal(t, "fuck", token)

	ret, token = s.Do("I'm a student.")
	assert.False(t, ret)
	assert.Empty(t, token)
}

func TestServe_DoAsync(t *testing.T) {
	e := &Example{
		Blacklist: [][]byte{
			[]byte("fuck"),
			[]byte("fucker"),
		},
	}
	s := &Serve{
		Tokenizer: e,
		Filer:     e,
	}

	ret, token := s.DoAsync(context.Background(), "")
	assert.False(t, ret)
	assert.Empty(t, token)

	ret, token = s.DoAsync(context.Background(), "fuck you.")
	assert.True(t, ret)
	assert.Equal(t, "fuck", token)

	ret, token = s.DoAsync(context.Background(), "I'm a student.")
	assert.False(t, ret)
	assert.Empty(t, token)
}
