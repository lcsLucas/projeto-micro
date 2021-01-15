package model

import "context"

type Aluno struct {
	RA      string `gorm:"primary_key;size:15; not null" json:"ra"`
	Nome    string `gorm:"size:255; not null" json:"nome"`
	Email   string `gorm:"size:255; not null" json:"email"`
	Celular string `gorm:"size:15; not null" json:"celular"`
}

type Repository interface {
	Create(ctx context.Context, alu Aluno) (bool, error)
	Alter(ctx context.Context, alu Aluno) (bool, error)
	Get(ctx context.Context, ra string) (Aluno, error)
	GetAll(ctx context.Context, page uint32) ([]Aluno, error)
	Delete(ctx context.Context, ra string) (bool, error)
	StatusService(ctx context.Context) (bool, error)
}
