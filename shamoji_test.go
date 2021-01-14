package shamoji_test

import (
	"context"
	"strings"
	"testing"

	"github.com/osamingo/shamoji"
)

type example struct {
	DenyList [][]byte
}

func newExample(deny ...[]byte) *shamoji.Serve {
	e := &example{
		DenyList: deny,
	}
	return &shamoji.Serve{
		Tokenizer: e,
		Filer:     e,
	}
}

func (e *example) Tokenize(sentence string) [][]byte {
	fs := strings.Fields(sentence)
	ts := make([][]byte, len(fs))
	for i := range fs {
		ts[i] = []byte(fs[i])
	}

	return ts
}

func (e *example) Test(src []byte) bool {
	for i := range e.DenyList {
		if string(src) == string(e.DenyList[i]) {
			return true
		}
	}

	return false
}

func TestServe_Do(t *testing.T) {
	s := newExample([]byte("fuck"), []byte("fucker"))

	cases := map[string]struct {
		sentence string
		result   bool
		expect   string
	}{
		"Empty sentence":             {"", false, ""},
		"Sentence with deny word":    {"I'm a student.", false, ""},
		"Sentence without deny word": {"fuck you.", true, "fuck"},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			ret, token := s.Do(c.sentence)
			if ret != c.result {
				t.Error("shoud be", c.result)
			}
			if token != c.expect {
				t.Error("shoud be", c.result)
			}
		})
	}
}

func TestServe_DoAsync(t *testing.T) {
	s := newExample([]byte("fuck"), []byte("fucker"))

	cases := map[string]struct {
		sentence string
		result   bool
		expect   string
	}{
		"Empty sentence":             {"", false, ""},
		"Sentence with deny word":    {"I'm a student.", false, ""},
		"Sentence without deny word": {"fuck you.", true, "fuck"},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			ret, token := s.DoAsync(context.Background(), c.sentence)
			if ret != c.result {
				t.Error("shoud be ", c.result)
			}
			if token != c.expect {
				t.Error("shoud be ", c.result)
			}
		})
	}
}
