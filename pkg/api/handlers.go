package api

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gaborszekely/golang-autocomplete-api/pkg/autocomplete"
)

// handleType function
func handleType(w http.ResponseWriter, r *http.Request, wordTrie *autocomplete.TrieNode) {
	w.Header().Set("Content-Type", "application/json")

	var suggestionRequest SuggestionRequest
	_ = json.NewDecoder(r.Body).Decode(&suggestionRequest)
	phrase := suggestionRequest.Phrase

	isWord := getIsWord(phrase)
	isWordEnd := getIsWordEnd(phrase)
	isSentenceEnd := getIsSentenceEnd(phrase)

	if isWord {
		// If currently typing a word, return suggestions
		lastWord := getLastWord(phrase)
		suggestions := autocomplete.GetWords(wordTrie, lastWord)
		apiResponse := Response{phrase, suggestions}
		json.NewEncoder(w).Encode(apiResponse)

	} else if isWordEnd {
		// If word has been typed, add to word Trie, return sentence suggestions
		autocomplete.AddWord(wordTrie, phrase[0:len(phrase)-1])
		apiResponse := Response{phrase, []string{}} // TODO - Implement sentence suggestions
		json.NewEncoder(w).Encode(apiResponse)

	} else if isSentenceEnd {
		// If sentence has been typed, add to sentence Trie
		// TODO - Implement sentence trie
		apiResponse := Response{phrase, []string{}} // TODO - Implement sentence suggestions
		json.NewEncoder(w).Encode(apiResponse)
	}

	// Check if phrase is currently on a word
	// Check this by making sure we did not terminate a senetence

	// Check if phrase is end of sentence
	// Check for '. '
	// If so, add to Sentence PrefixTree
}

func getLastWord(phrase string) string {
	chars := []rune(phrase)
	word := ""
	for i := len(chars) - 1; i >= 0; i-- {
		char := string(chars[i])
		if char == " " {
			break
		}
		word = char + word
	}
	return word
}

func getIsSentenceEnd(phrase string) bool {
	match, _ := regexp.MatchString("\\.\\s$", phrase)
	return match
}

func getIsWordEnd(phrase string) bool {
	match, _ := regexp.MatchString("[^\\.]\\s$", phrase)
	return match
}

func getIsWord(phrase string) bool {
	match, _ := regexp.MatchString("\\w$", phrase)
	return match
}
