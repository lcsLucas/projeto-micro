package model

import "context"

type Exercicio struct {
	ID        uint64 `json:"id"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
	Materia   string `json:"materia"`
	Ativo     bool   `json:"ativo"`
}

type Repository interface {
	Create(ctx context.Context, exe Exercicio) (bool, error)
	Alter(ctx context.Context, exe Exercicio) (bool, error)
	Get(ctx context.Context, id uint64) (Exercicio, error)
	GetSomes(ctx context.Context, id []uint64) ([]Exercicio, error)
	Delete(ctx context.Context, id uint64) (bool, error)
	StatusService(ctx context.Context) (bool, error)
}
