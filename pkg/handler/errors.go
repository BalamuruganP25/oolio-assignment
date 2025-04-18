package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrResponse struct {
	Title   string `json:"title"`
	Details string `json:"details"`
}

func ErrorResponse(w http.ResponseWriter, status int, ErrRes ErrResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(ErrRes); err != nil {
		log.Printf("Failed to encode error response: %v", err)
	}
}

// SendResponse is a utility function to send the response to the client
func SendResponse(w http.ResponseWriter, response any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
