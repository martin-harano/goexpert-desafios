package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/martin-harano/goexpert-desafios/tree/main/desafio2/with-api/internal/entity"
	"github.com/martin-harano/goexpert-desafios/tree/main/desafio2/with-api/internal/infra/webclient"
)

type CepHandler struct {
}

func NewCepHandler() *CepHandler {
	return &CepHandler{}
}

func (h *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if code == "" {
		response := Error{StatusCode: http.StatusBadRequest, Message: "Please set a CEP code"}
		response.WriteResponse(w)
		return
	}
	err := h.validateCep(code)
	if err != nil {
		response := Error{StatusCode: http.StatusBadRequest, Message: err.Error()}
		response.WriteResponse(w)
		return
	}
	code = h.normalizeCep(code)

	requestDone := atomic.Bool{}
	requestDone.Store(false)

	apiCepClient := webclient.ApiCepClient{Discard: &requestDone}
	viaCepClient := webclient.ViaCepClient{Discard: &requestDone}

	result1 := make(chan entity.Cep)
	result2 := make(chan entity.Cep)

	go apiCepClient.Provide(code, result1)
	go viaCepClient.Provide(code, result2)

	select {
	case cep := <-result1:
		requestDone.Store(true)
		h.WriteResponse(w, "APICEP", &cep)
	case cep := <-result2:
		requestDone.Store(true)
		h.WriteResponse(w, "VIACEP", &cep)
	case <-time.After(time.Second):
		requestDone.Store(true)
		println("timeout")
		response := Error{StatusCode: http.StatusGatewayTimeout, Message: "Server waited too long"}
		response.WriteResponse(w)
	}
}

func (h *CepHandler) WriteResponse(w http.ResponseWriter, apiName string, cep *entity.Cep) {
	result, err := json.Marshal(cep)
	if err != nil {
		response := Error{StatusCode: http.StatusBadGateway, Message: err.Error()}
		response.WriteResponse(w)
		return
	}
	fmt.Fprintf(os.Stderr, "CEP obtido pela api %s: %v\n", apiName, string(result))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}

func (h *CepHandler) validateCep(code string) error {
	_, err := strconv.Atoi(code)
	if err != nil {
		return errors.New("Only numbers are allowed for CEP code")
	}
	if len(code) != 8 {
		return errors.New("CEP code should have exactly 8 numeric digits")
	}
	return nil
}

func (h *CepHandler) normalizeCep(code string) string {
	return code[0:5] + "-" + code[5:]
}
