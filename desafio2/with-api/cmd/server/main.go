package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/martin-harano/goexpert-desafios/tree/main/desafio2/with-api/docs"
	"github.com/martin-harano/goexpert-desafios/tree/main/desafio2/with-api/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Desafio2 - Multithread
// @version         1.0
// @description     API to demonstrate multithreading
// @termsOfService  http://swagger.io/terms/

// @contact.name   Martin Harano
// @contact.url    https://github.com/martin-harano
// @contact.email  martin.harano@fullcycle.com.br

// @license.name   Martin Harano License
// @license.url    https://github.com/martin-harano

// @host      localhost:8000
// @BasePath  /
// @in header
// @name Authorization
func main() {

	cepHandler := handlers.CepHandler{}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/cep/{code}", cepHandler.GetCep)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}
