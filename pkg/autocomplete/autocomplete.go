package autocomplete

import (
	"sort"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Trie struct
type Trie struct {
	root *TrieNode
}

// TrieNode struct
type TrieNode struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Val        string             `bson:"value" json:"value,omitempty"`
	Children   [](*TrieNode)      `bson:"children" json:"children,omitempty"`
	IsEnd      bool               `isEnd:"value" json:"isEnd,omitempty"`
	Occurences int                `bson:"occurences" json:"occurences,omitempty"`
}

// ResponseOccurences struct
type ResponseOccurences struct {
	Response   string
	Occurences int
}

// AddWord function
func AddWord(node *TrieNode, value string) {
	if len(value) == 0 {
		return
	}

	currChar := value[0:1]
	isLastChar := len(value) == 1

	for _, child := range node.Children {
		if child.Val == currChar {
			if isLastChar {
				child.IsEnd = true
				child.Occurences++
			}
			AddWord(child, value[1:])
			return
		}
	}

	// Children does not have current character
	occurences := 0
	if isLastChar {
		occurences = 1
	}
	newNode := TrieNode{
		Val:        currChar,
		Children:   make([]*TrieNode, 0),
		IsEnd:      isLastChar,
		Occurences: occurences,
	}
	node.Children = append(node.Children, &newNode)
	AddWord(&newNode, value[1:])
}

// GetWords function
func GetWords(node *TrieNode, prefix string) []string {
	res := make([](*ResponseOccurences), 0)

	if len(prefix) > 0 {
		recurseWords(node, "", prefix, &res)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Occurences > res[j].Occurences
	})

	response := make([]string, 0)

	for _, el := range res {
		response = append(response, el.Response)
	}

	return response
}

func recurseWords(currNode *TrieNode, currWord string, prefix string, res *[](*ResponseOccurences)) {
	if currNode.IsEnd && len(prefix) == 0 {
		*res = append(*res, &ResponseOccurences{
			Response:   currWord,
			Occurences: currNode.Occurences,
		})
	}

	prefixChar, prefixSlice := "", ""
	if len(prefix) > 0 {
		prefixChar = prefix[0:1]
		prefixSlice = prefix[1:]
	}

	for _, child := range currNode.Children {
		if child.Val == prefixChar || prefixChar == "" {
			recurseWords(child, currWord+child.Val, prefixSlice, res)
		}
	}
}

// GetOccurences function
func GetOccurences(node *TrieNode, word string) int {
	occurences := 0

	if len(word) == 0 {
		occurences = node.Occurences
	} else {
		for _, child := range node.Children {
			if child.Val == word[0:1] {
				occurences = GetOccurences(child, word[1:])
				break
			}
		}
	}

	return occurences
}
