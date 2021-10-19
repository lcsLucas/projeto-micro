package exercicio

import (
	"context"
	"errors"

	"github.com/lcslucas/projeto-micro/services/exercicio/model"
)

type exercicioService struct {
	repository model.Repository
}

/*NewService cria um novo serviço de exercicio */
func NewService(rep model.Repository) Service {
	return &exercicioService{
		repository: rep,
	}
}

func (e exercicioService) Create(ctx context.Context, exe model.Exercicio) error {
	return errors.New("not implemented")
}

func (e exercicioService) Alter(ctx context.Context, exe model.Exercicio) error {
	return errors.New("not implemented")
}

func (e exercicioService) Get(ctx context.Context, id uint64) (model.Exercicio, error) {
	return model.Exercicio{}, errors.New("not implemented")
}

func (e exercicioService) GetSomes(ctx context.Context, ids []uint64) ([]model.Exercicio, error) {

	exeRes, err := e.repository.GetSomes(ctx, ids)
	if err != nil {
		return []model.Exercicio{}, err
	}

	return exeRes, nil
}

func (e exercicioService) Delete(ctx context.Context, id uint64) error {
	return errors.New("not implemented")
}

func (e exercicioService) StatusService(ctx context.Context) error {
	return nil
}
