package handler

import (
	"encoding/json"
	"net/http"
)

type ErrResponse struct {
	Tittle  string `json:"tittle"`
	Details string `json:"details"`
}

func ErrorResponse(w http.ResponseWriter, status int, ErrRes ErrResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrRes)
}

// SendResponse is a utility function to send the response to the client
func SendResponse(w http.ResponseWriter, response any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
