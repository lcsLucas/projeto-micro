package repository

import (
	"context"
	"errors"

	"github.com/lcslucas/projeto-micro/config"
	"github.com/lcslucas/projeto-micro/services/prova/model"
	"gorm.io/gorm"
)

//const tableName = "provas"

type repository struct {
	db       *gorm.DB
	configDB config.ConfigDB
}

//NewRepository cria um novo repositório para o serviço
func NewRepository(db *gorm.DB, configDB config.ConfigDB) model.Repository {
	return &repository{
		db:       db,
		configDB: configDB,
	}
}

func (r *repository) Create(ctx context.Context, pro model.Prova) (bool, error) {
	return false, nil
}

func (r *repository) Alter(ctx context.Context, pro model.Prova) (bool, error) {
	return false, nil
}

func (r *repository) Get(ctx context.Context, id uint64) (model.Prova, error) {
	return model.Prova{}, nil
}

func (r *repository) GetProvaAluno(ctx context.Context, idProva uint64, raAluno string) (model.Prova, error) {
	var err error
	var p model.Prova
	var pa model.ProvaAluno
	var pe []model.ProvaExercicios

	tx := r.db.WithContext(ctx)

	err = tx.Debug().Model(model.Prova{}).Where("id = ?", idProva).Take(&p).Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Prova{}, errors.New("nenhuma prova encontrada")
		}

		return model.Prova{}, err
	}

	err = tx.Debug().Model(model.ProvaAluno{}).Where("prova_id = ? AND aluno_ra = ?", p.ID, raAluno).Take(&pa).Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Prova{}, errors.New("prova não relacionada para esse aluno")
		}

		return model.Prova{}, err

	}

	err = tx.Debug().Model(model.ProvaExercicios{}).Where("prova_aluno_id = ?", pa.ID).Find(&pe).Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Prova{}, errors.New("nenhum exercício relacionado com essa Prova")
		}

		return model.Prova{}, err
	}

	p.Aluno.RA = pa.AlunoRA
	p.ProvaExercicios = pe

	return p, nil
}

func (r *repository) GetAll(ctx context.Context, page uint32) ([]model.Prova, error) {
	var err error

	provas := []model.Prova{}

	tx := r.db.WithContext(ctx)

	err = tx.Debug().Model(&model.Prova{}).Find(&provas).Error
	if err != nil {
		return []model.Prova{}, err
	}

	return provas, nil
}

func (r *repository) Delete(ctx context.Context, id uint64) (bool, error) {
	return false, nil
}

func (r *repository) StatusService(ctx context.Context) (bool, error) {
	return false, nil
}
