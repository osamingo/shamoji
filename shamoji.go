package shamoji

import (
	"errors"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

type (
	// Tokenizer implements Tokenize method.
	Tokenizer interface {
		Tokenize(sentence string) (tokens [][]byte)
	}
	// Filter implements Test method.
	Filter interface {
		Test(src []byte) (result bool)
	}
	// Serve has Tokenizer and Filter interfaces.
	Serve struct {
		Tokenizer Tokenizer
		Filer     Filter
	}
)

// Do filtering sentence.
func (s *Serve) Do(sentence string) (bool, string) {

	ts := s.Tokenizer.Tokenize(sentence)
	if len(ts) == 0 {
		return false, ""
	}

	for i := range ts {
		if s.Filer.Test(ts[i]) {
			return true, string(ts[i])
		}
	}

	return false, ""
}

// DoAsync filtering sentence.
func (s *Serve) DoAsync(c context.Context, sentence string) (bool, string) {

	ts := s.Tokenizer.Tokenize(sentence)
	if len(ts) == 0 {
		return false, ""
	}

	eg, _ := errgroup.WithContext(c)
	for i := range ts {
		i := i
		eg.Go(func() error {
			if s.Filer.Test(ts[i]) {
				return errors.New(string(ts[i]))
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return true, err.Error()
	}

	return false, ""
}
