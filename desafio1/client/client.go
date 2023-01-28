package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	c := http.Client{Timeout: time.Second}
	resp, err := c.Get("http://localhost:8080/cotacao")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	cotacao, err := parseCotacao(body)
	if err != nil {
		panic(err)
	}
	err = saveCotacao(cotacao)
	if err != nil {
		panic(err)
	}
	println("Dolar price written at cotacao.txt")
}

func parseCotacao(b []byte) (*Cotacao, error) {
	var cotacao Cotacao
	err := json.Unmarshal(b, &cotacao)
	if err != nil {
		return nil, err
	}
	return &cotacao, nil
}

func saveCotacao(c *Cotacao) error {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("Dolar: " + c.Usdbrl.Bid)
	if err != nil {
		return err
	}
	return nil
}
