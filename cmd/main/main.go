package main

import (
	"github.com/gaborszekely/golang-autocomplete-api/pkg/api"
	"github.com/gaborszekely/golang-autocomplete-api/pkg/autocomplete"
)

func main() {
	wordTrie := autocomplete.TrieNode{
		Val:        "",
		Children:   make([](*autocomplete.TrieNode), 0),
		IsEnd:      false,
		Occurences: 0,
	}
	api.Run(&wordTrie)
}

/*
autocomplete.AddWord(&trie, "test")
autocomplete.AddWord(&trie, "app")
autocomplete.AddWord(&trie, "apple")
autocomplete.AddWord(&trie, "apple")

fmt.Println(autocomplete.GetWords(&trie, ""))
fmt.Println(autocomplete.GetWords(&trie, "ap"))
fmt.Println(autocomplete.GetWords(&trie, "appl"))

fmt.Println(autocomplete.GetOccurences(&trie, "apple"))
*/
