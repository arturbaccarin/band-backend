package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func ErrorResponse(w http.ResponseWriter, statusCode int, errorRequest error) {
	w.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(w).Encode(ErrorMessage{Error: errorRequest.Error()})
	if err != nil {
		panic(err)
	}
}
