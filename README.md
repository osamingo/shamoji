# shamoji

[![Travis branch](https://img.shields.io/travis/osamingo/shamoji/master.svg)](https://travis-ci.org/osamingo/shamoji)
[![codecov](https://codecov.io/gh/osamingo/shamoji/branch/master/graph/badge.svg)](https://codecov.io/gh/osamingo/shamoji)
[![Go Report Card](https://goreportcard.com/badge/osamingo/shamoji)](https://goreportcard.com/report/osamingo/shamoji)
[![codebeat badge](https://codebeat.co/badges/9d9fdf3d-0c6d-455f-8444-8399a07d49ae)](https://codebeat.co/projects/github-com-osamingo-shamoji-master)
[![GoDoc](https://godoc.org/github.com/osamingo/shamoji?status.svg)](https://godoc.org/github.com/osamingo/shamoji)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/osamingo/shamoji/master/LICENSE)

## About

The shamoji (杓文字) is word filtering package.

## Install

```
$ go get -u github.com/osamingo/shamoji
```

## Usage

```go
package main

import (
	"fmt"
	"sync"

	"github.com/osamingo/shamoji"
	"github.com/osamingo/shamoji/japanese"
)

var (
	o sync.Once
	s *shamoji.Serve
)

func main() {
	yes, word := Contains("いつか殺す")
	fmt.Printf("Result: %v, Word: %s", yes, word)
}

func Contains(sentence string) (bool, string) {
	o.Do(func() {
		s = japanese.NewServe([]string{
			"死ね",
			"殺す",
		})
	})
	return s.Do(sentence)
}
```

## License

Released under the [MIT License](https://github.com/osamingo/shamoji/blob/master/LICENSE).
