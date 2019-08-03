package main

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("hi")
	trie.Insert("imgoogege")
	trie.Insert("scokinglove")
	trie.Insert("sockingfuck")
	trie.Insert("sockinggirl")
	trie.Insert("sockingwoman")
	fmt.Println(trie.Search("sockingwoman"))
	fmt.Println(trie.StartsWith("socking"))
}
