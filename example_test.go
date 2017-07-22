package set_test

import (
	"fmt"

	"github.com/suntong/set"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleSet() {
	s1 := set.NewSet()
	s1.Add("Apple")
	s1.Add("Orange")
	fmt.Println(s1.Has("Apple"))
	fmt.Println(s1.Has("Grape"))

	s2 := set.NewSetFrom([]string{"Apple", "Orange"})
	fmt.Println(s2.Has("Apple"))
	fmt.Println(s2.Has("Grape"))

	// Output:
	// true
	// false
	// true
	// false
}

// To show the full code in GoDoc
type dummy struct {
}
