package endpoints

import (
	"github.com/lcslucas/projeto-micro/services/aluno/model"
)

type CreateAlterRequest struct {
	Aluno model.Aluno `json:"aluno"`
}

type CreateAlterResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

type GetRequest struct {
	ID uint64 `json:"id"`
}

type GetResponse struct {
	Aluno  model.Aluno `json:"aluno"`
	Status bool        `json:"status"`
	Error  string      `json:"error"`
}

type GetAllRequest struct {
	Page uint32 `json:"page"`
}

type GetAllResponse struct {
	Alunos []model.Aluno `json:"alunos"`
	Status bool          `json:"status"`
	Error  string        `json:"error"`
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
