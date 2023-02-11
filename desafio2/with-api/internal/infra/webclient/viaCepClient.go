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

type ViaCepClient struct {
	Discard *atomic.Bool
}

type viaCepDtoResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (c *ViaCepClient) Provide(cep string, result chan<- entity.Cep) {
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ViaCepClient: Erro ao fazer requisição: %v\n", err)
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ViaCepClient: Erro ao ler resposta: %v\n", err)
	}
	var apiCep viaCepDtoResponse
	err = json.Unmarshal(res, &apiCep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ViaCepClient: Erro ao fazer parse da resposta: %v\n", err)
	}
	if c.Discard.Load() {
		return
	}
	result <- entity.Cep{
		Code:     apiCep.Cep,
		State:    apiCep.Uf,
		City:     apiCep.Localidade,
		District: apiCep.Bairro,
		Address:  fmt.Sprintf("%s%s", apiCep.Logradouro, c.addAddressComplement(apiCep.Complemento)),
	}
}

func (c *ViaCepClient) addAddressComplement(AddressComplement string) string {
	if len(AddressComplement) == 0 {
		return ""
	}
	return " - " + AddressComplement
}
