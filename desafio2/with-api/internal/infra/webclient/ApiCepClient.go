package webclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/martin-harano/goexpert-desafios/tree/main/desafio2/with-api/internal/entity"
)

type ApiCepClient struct {
	Discard *atomic.Bool
}

type apiCepDtoResponse struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func (c *ApiCepClient) Provide(cep string, result chan<- entity.Cep) {
	resp, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ApiCepClient: Erro ao fazer requisição: %v\n", err)
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ApiCepClient: Erro ao ler resposta: %v\n", err)
	}
	var apiCep apiCepDtoResponse
	err = json.Unmarshal(res, &apiCep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ApiCepClient: Erro ao fazer parse da resposta: %v\n", err)
	}
	if c.Discard.Load() {
		return
	}
	result <- entity.Cep{
		Code:     apiCep.Code,
		State:    apiCep.State,
		City:     apiCep.City,
		District: apiCep.District,
		Address:  apiCep.Address,
	}
}
