package shamoji

import (
	"context"
	"errors"

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
func (s *Serve) DoAsync(ctx context.Context, sentence string) (bool, string) {
	ts := s.Tokenizer.Tokenize(sentence)
	if len(ts) == 0 {
		return false, ""
	}

	eg, _ := errgroup.WithContext(ctx)

	for i := range ts {
		eg.Go(func() error {
			if s.Filer.Test(ts[i]) {
				// define found error type...
				//nolint: goerr113
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
