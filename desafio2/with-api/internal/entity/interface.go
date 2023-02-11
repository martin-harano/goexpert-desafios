package entity

type CepProviderInterface interface {
	provide(cepCode string, result chan<- Cep)
}
