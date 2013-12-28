package trie_test

import (
	"fmt"

	"github.com/cespare/trie"
)

// This example shows simple usage of a Trie as a []byte set.
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
