package handlers

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func (e *Error) WriteResponse(w http.ResponseWriter) {
	w.WriteHeader(e.StatusCode)
	json.NewEncoder(w).Encode(*e)
}
