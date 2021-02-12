package model

import (
	"context"
	"time"

	"github.com/lcslucas/projeto-micro/services/aluno/model"
	mod_exe "github.com/lcslucas/projeto-micro/services/exercicio/model"
)

type Tabler interface {
	TableName() string
}

type Prova struct {
	ID              uint64              `gorm:"primary_key; auto_increment; not null" json:"id"`
	Nome            string              `gorm:"size:255; not null" json:"nome"`
	DataCadastro    time.Time           `gorm:"not null" json:"dt_cadastro"`
	DataInicio      time.Time           `gorm:"not null" json:"dt_inicio"`
	DataFinal       time.Time           `gorm:"not null" json:"dt_final"`
	Serie           string              `gorm:"not null" json:"serie,omitempty"`
	Materia         string              `gorm:"-" json:"materia,omitempty"`
	Bimestre        uint16              `gorm:"not null" json:"bimestre,omitempty"`
	Finalizada      bool                `json:"finalizada,omitempty"`
	Aluno           model.Aluno         `gorm:"-" json:"aluno"`
	Exercicios      []mod_exe.Exercicio `gorm:"-" json:"exercicios"`
	ProvaAlunos     []ProvaAluno        `gorm:"-" json:"-"`
	ProvaExercicios []ProvaExercicios   `gorm:"-" json:"-"`
}

func (Prova) TableName() string {
	return "provas"
}

type ProvaExercicios struct {
	ProvaAluno   ProvaAluno `gorm:"ForeignKey:ProvaAlunoID;association_foreignkey:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	ProvaAlunoID uint64     `json:"-"`
	ExercicioID  uint64     `gorm:"not null" json:"-"`
}

func (ProvaExercicios) TableName() string {
	return "provas_exercicios"
}

type ProvaAluno struct {
	ID      uint64 `gorm:"primary_key; auto_increment; not null" json:"-"`
	Prova   Prova  `gorm:"ForeignKey:ProvaID;association_foreignkey:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"usuario"`
	ProvaID uint64 `json:"-"`
	AlunoRA string `gorm:"size:15; not null" json:"-"`
}

func (ProvaAluno) TableName() string {
	return "provas_alunos"
}

type Repository interface {
	Create(ctx context.Context, pro Prova) (bool, error)
	Alter(ctx context.Context, pro Prova) (bool, error)
	Get(ctx context.Context, id uint64) (Prova, error)
	GetProvaAluno(ctx context.Context, idProva uint64, raAluno string) (Prova, error)
	GetAll(ctx context.Context, page uint32) ([]Prova, error)
	Delete(ctx context.Context, id uint64) (bool, error)
	StatusService(ctx context.Context) (bool, error)
}
