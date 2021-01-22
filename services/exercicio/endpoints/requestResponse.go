package endpoints

import (
	"github.com/lcslucas/projeto-micro/services/exercicio/model"
)

type CreateAlterRequest struct {
	Exercicio model.Exercicio `json:"exercicio"`
}

type CreateAlterResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

type GetRequest struct {
	ID uint64 `json:"id"`
}

type GetResponse struct {
	Exercicio model.Exercicio `json:"exercicio"`
	Status    bool            `json:"status"`
	Error     string          `json:"error"`
}

type GetSomesRequest struct {
	Ids []uint64 `json:"ids"`
}

type GetSomesResponse struct {
	Exercicios []model.Exercicio
	Status     bool   `json:"status"`
	Error      string `json:"error"`
}

type DeleteRequest struct {
	ID uint64 `json:"id"`
}

type DeleteResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

type StatusServiceRequest struct{}

type StatusServiceResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}
