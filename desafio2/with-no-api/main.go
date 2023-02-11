package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	code := "06233030"
	if len(os.Args) > 1 {
		code = os.Args[1]
	}

	err := validateCep(code)
	if err != nil {
		panic(err)
	}

	code = normalizeCep(code)

	result1 := make(chan string)
	result2 := make(chan string)

	go cepApiCep(result1, code)
	go cepViaCep(result2, code)

	select {
	case msg := <-result1:
		fmt.Println("APICEP: ", msg)
	case msg := <-result2:
		fmt.Println("VIACEP: ", msg)
	case <-time.After(time.Second):
		println("timeout")
	}

}

func cepApiCep(result chan<- string, cep string) {
	req, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	result <- string(res)
}

func cepViaCep(result chan<- string, cep string) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	result <- string(res)
}

func validateCep(code string) error {
	_, err := strconv.Atoi(code)
	if err != nil {
		return errors.New("Only numbers are allowed for CEP code")
	}
	if len(code) != 8 {
		return errors.New("CEP code should have exactly 8 numeric digits")
	}
	return nil
}

func normalizeCep(code string) string {
	return code[0:5] + "-" + code[5:]
}
