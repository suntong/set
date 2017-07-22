
# set

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/set?status.svg)](http://godoc.org/github.com/suntong/set)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/set)](https://goreportcard.com/report/github.com/suntong/set)
[![travis Status](https://travis-ci.org/suntong/set.svg?branch=master)](https://travis-ci.org/suntong/set)

## TOC
- [API](#api)
  - [> example_test.go](#-example_testgo)

Extremely light-weighted SET handlings.

# API

#### > example_test.go
```go
package set_test

import (
	"fmt"

	"github.com/suntong/set"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleSet() {
	{
		// == case-sensitive SET
		s1 := set.NewSet()
		s1.Add("Apple")
		s1.Add("Orange")
		fmt.Println(s1.Has("Apple"))
		fmt.Println(s1.Has("Grape"))

		s2 := set.NewSetFrom([]string{"Apple", "Orange"})
		fmt.Println(s2.Has("Apple"))
		s2.Add("Grape")
		fmt.Println(s2.Len(), s2.Has("Grape"), s2.Has("grape"))
		s2.Del("Grape")
		fmt.Println(s2.Len(), s2.Has("Grape"))
		fmt.Println(s2.Get())
	}
	fmt.Println()

	{
		// == case-insensitive SET
		s1 := set.NewSet0()
		s1.Add("Apple")
		s1.Add("Orange")
		fmt.Println(s1.Has("Apple"))
		fmt.Println(s1.Has("Grape"))

		s2 := set.NewSet0From([]string{"Apple", "Orange"})
		fmt.Println(s2.Has("Apple"))
		s2.Add("Grape")
		fmt.Println(s2.Len(), s2.Has("Grape"), s2.Has("grape"))
		s2.Del("Grape")
		fmt.Println(s2.Len(), s2.Has("Grape"))
		fmt.Println(s2.Get())
	}

	// Output:
	// true
	// false
	// true
	// 3 true false
	// 2 false
	// [Apple Orange]
	//
	// true
	// false
	// true
	// 3 true true
	// 2 false
	// [apple orange]
}

// To show the full code in GoDoc
var s set.Set = set.NewSet()
```

All patches welcome.
