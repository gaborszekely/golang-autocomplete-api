package autocomplete

// Trie struct
type Trie struct {
	root *TrieNode
}

// TrieNode struct
type TrieNode struct {
	Val        string
	Children   [](*TrieNode)
	IsEnd      bool
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
	newNode := TrieNode{currChar, make([]*TrieNode, 0), isLastChar, occurences}
	node.Children = append(node.Children, &newNode)
	AddWord(&newNode, value[1:])
}

// GetWords function
func GetWords(node *TrieNode, prefix string) []string {
	res := make([]string, 0)
	if len(prefix) > 0 {
		recurse(node, "", prefix, &res)
	}
	return res
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

func recurse(currNode *TrieNode, currWord string, prefix string, res *[]string) {
	// currWord += currNode.val
	if currNode.IsEnd && len(prefix) == 0 {
		*res = append(*res, currWord)
	}

	prefixChar, prefixSlice := "", ""
	if len(prefix) > 0 {
		prefixChar = prefix[0:1]
		prefixSlice = prefix[1:]
	}

	for _, child := range currNode.Children {
		if child.Val == prefixChar || prefixChar == "" {
			recurse(child, currWord+child.Val, prefixSlice, res)
		}
	}
}
