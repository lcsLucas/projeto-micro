package endpoints

import "github.com/lcslucas/projeto-micro/services/prova/model"

type CreateAlterRequest struct {
	Prova model.Prova `json:"prova"`
}

type CreateAlterResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

type GetRequest struct {
	ID uint64 `json:"id"`
}

type GetResponse struct {
	Prova  model.Prova `json:"prova"`
	Status bool        `json:"status"`
	Error  string      `json:"error"`
}

type GetProvaAlunoRequest struct {
	IDProva uint64 `json:"id_prova"`
	RaAluno string `json:"ra_aluno"`
}

type GetAllRequest struct {
	Page uint32 `json:"page"`
}

type GetAllResponse struct {
	Provas []model.Prova `json:"prova"`
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
