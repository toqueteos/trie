package trie_test

import (
	"fmt"

	"github.com/cespare/trie"
)

func ExampleSimple() {
	t := trie.New()
	t.Insert([]byte("hello"))
	t.Insert([]byte("world"))
	for _, s := range []string{"hello", "world", "he", "h", "worlds", ""} {
		fmt.Println(t.Contains([]byte(s)))
	}
	// Output:
	// true
	// true
	// false
	// false
	// false
	// false
}
