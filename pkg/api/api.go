package api

import (
	"log"
	"net/http"

	"github.com/gaborszekely/golang-autocomplete-api/pkg/autocomplete"
	"github.com/gorilla/mux"
)

// import (
// 	"encoding/json"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

/**
 * Structs
 */

// ResponseOccurences struct
// type ResponseOccurences struct {
// 	response   string
// 	occurences int
// }

// Response struct
type Response struct {
	SearchPhrase string   `json:"searchphrase"`
	Suggestions  []string `json:"suggestions"`
}

// SuggestionRequest struct
type SuggestionRequest struct {
	Phrase string `json:"phrase"`
}

// Run function
func Run(wordTrie *autocomplete.TrieNode) {
	router := mux.NewRouter()

	// Handle `/api/type` endpoint
	router.HandleFunc("/api/type", func(w http.ResponseWriter, r *http.Request) {
		handleType(w, r, wordTrie)
	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
