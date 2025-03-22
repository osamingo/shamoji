# shamoji

[![CI](https://github.com/osamingo/shamoji/actions/workflows/actions.yml/badge.svg)](https://github.com/osamingo/shamoji/actions/workflows/actions.yml)
[![codecov](https://codecov.io/gh/osamingo/shamoji/branch/master/graph/badge.svg)](https://codecov.io/gh/osamingo/shamoji)
[![Go Report Card](https://goreportcard.com/badge/github.com/osamingo/shamoji)](https://goreportcard.com/report/github.com/osamingo/shamoji)
[![codebeat badge](https://codebeat.co/badges/9d9fdf3d-0c6d-455f-8444-8399a07d49ae)](https://codebeat.co/projects/github-com-osamingo-shamoji-master)
[![GoDoc](https://godoc.org/github.com/osamingo/shamoji?status.svg)](https://godoc.org/github.com/osamingo/shamoji)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/osamingo/shamoji/master/LICENSE)

## About

The shamoji (杓文字) is word filtering package.

## Install

```
$ go get github.com/osamingo/shamoji@latest
```

## Usage

```go
package main

import (
	"fmt"
	"sync"

	"github.com/osamingo/shamoji"
	"github.com/osamingo/shamoji/filter"
	"github.com/osamingo/shamoji/tokenizer"
	"golang.org/x/text/unicode/norm"
)

var (
	o sync.Once
	s *shamoji.Serve
)

func main() {
	yes, word := Contains("我が生涯に一片の悔い無し")
	fmt.Printf("Result: %v, Word: %s", yes, word)
}

func Contains(sentence string) (bool, string) {
	o.Do(func() {
		tok, err := tokenizer.NewKagomeTokenizer(norm.NFKC)
		if err != nil {
			panic(err)
		}
		s = &shamoji.Serve{
			Tokenizer: tok,
			Filer:     filter.NewCuckooFilter("涯に", "悔い"),
		}
	})
	return s.Do(sentence)
}
```

## License

Released under the [MIT License](https://github.com/osamingo/shamoji/blob/master/LICENSE).
