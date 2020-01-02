package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gaborszekely/golang-autocomplete-api/pkg/autocomplete"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func handleType(w http.ResponseWriter, r *http.Request, wordTrie *autocomplete.TrieNode, collection *mongo.Collection) {
	w.Header().Set("Content-Type", "application/json")

	var suggestionRequest SuggestionRequest
	_ = json.NewDecoder(r.Body).Decode(&suggestionRequest)
	phrase := suggestionRequest.Phrase

	// Avoid recursing on empty phrase
	if phrase == "" {
		apiResponse := Response{phrase, []string{}}
		json.NewEncoder(w).Encode(apiResponse)
		return
	}

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
		autocomplete.AddWord(wordTrie, getLastWord(phrase[0:len(phrase)-1]))
		apiResponse := Response{phrase, []string{}} // TODO - Implement sentence suggestions
		json.NewEncoder(w).Encode(apiResponse)
		updateTrie(collection, wordTrie)

	} else if isSentenceEnd {
		// If sentence has been typed, add to sentence Trie
		// TODO - Implement sentence trie
		apiResponse := Response{phrase, []string{}} // TODO - Implement sentence suggestions
		json.NewEncoder(w).Encode(apiResponse)
	}
}

func handleClear(w http.ResponseWriter, r *http.Request, wordTrie *autocomplete.TrieNode, collection *mongo.Collection) {
	w.Header().Set("Content-Type", "application/json")
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": wordTrie.ID})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted Count: %d\n", res.DeletedCount)

	*wordTrie = *getRootNode(collection)

	json.NewEncoder(w).Encode(GenericResponse{
		Status:  200,
		Message: "Successfully cleared suggestions",
	})
}

func updateTrie(collection *mongo.Collection, wordTrie *autocomplete.TrieNode) {
	// Update trie in database
	resultUpdate, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": wordTrie.ID},
		bson.M{
			"$set": wordTrie,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Modified Count: %d\n", resultUpdate.ModifiedCount)
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
