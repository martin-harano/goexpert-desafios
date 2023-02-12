package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/martin-harano/goexpert-desafios/tree/main/desafio2/with-api/internal/dto"
)

type Error struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func (e *Error) WriteResponse(w http.ResponseWriter) {
	w.WriteHeader(e.StatusCode)
	err := dto.ErrorOutput{
		Message: e.Message,
	}
	json.NewEncoder(w).Encode(&err)
}
