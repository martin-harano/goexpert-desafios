package dto

type GetCepOutput struct {
	Code     string `json:"code" example:"06233-030"`
	State    string `json:"state" example:"SP"`
	City     string `json:"city" example:"Osasco"`
	District string `json:"district" example:"Piratininga"`
	Address  string `json:"address" example:"Rua Paula Rodrigues"`
} // @name CEP

type ErrorOutput struct {
	Message string `json:"message"`
} // @name ERROR
