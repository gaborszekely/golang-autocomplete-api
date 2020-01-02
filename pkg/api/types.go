package api

// Response struct
type Response struct {
	SearchPhrase string   `json:"searchphrase"`
	Suggestions  []string `json:"suggestions"`
}

// GenericResponse struct
type GenericResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// SuggestionRequest struct
type SuggestionRequest struct {
	Phrase string `json:"phrase"`
}
