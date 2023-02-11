package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/martin-harano/goexpert-desafios/tree/main/desafio2/with-api/internal/infra/webserver/handlers"
)

func main() {

	cepHandler := handlers.CepHandler{}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/cep/{code}", cepHandler.GetCep)

	http.ListenAndServe(":8000", r)
}
