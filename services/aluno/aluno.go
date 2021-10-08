package aluno

import (
	"context"
	"errors"

	"github.com/lcslucas/projeto-micro/services/aluno/model"
)

type alunoService struct {
	repository model.Repository
}

/*NewService cria um novo serviço de aluno */
func NewService(rep model.Repository) Service {
	return &alunoService{
		repository: rep,
	}
}

func (a *alunoService) Create(ctx context.Context, alu model.Aluno) (bool, error) {
	return false, errors.New("not implemented")
}

func (a *alunoService) Alter(ctx context.Context, alu model.Aluno) (bool, error) {
	return false, errors.New("not implemented")
}

func (a *alunoService) Get(ctx context.Context, ra string) (model.Aluno, error) {
	aluRes, err := a.repository.Get(ctx, ra)
	if err != nil {
		return model.Aluno{}, err
	}

	return aluRes, nil
}

func (a *alunoService) GetAll(ctx context.Context, page uint32) ([]model.Aluno, error) {
	return []model.Aluno{}, nil
}

func (a *alunoService) Delete(ctx context.Context, ra string) (bool, error) {
	return false, errors.New("not implemented")
}

func (a *alunoService) StatusService(ctx context.Context) (bool, error) {
	return true, nil
}
