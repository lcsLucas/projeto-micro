package aluno

import (
	"context"

	"github.com/lcslucas/projeto-micro/services/aluno/model"
)

type alunoService struct{}

/*NewService cria um novo serviço de aluno */
func NewService() Service {
	return &alunoService{}
}

func (a *alunoService) Create(ctx context.Context, alu model.Aluno) (bool, error) {
	return false, nil
}

func (a *alunoService) Alter(ctx context.Context, alu model.Aluno) (bool, error) {
	return false, nil
}

func (a *alunoService) Get(ctx context.Context, id uint64) (model.Aluno, error) {
	return model.Aluno{}, nil
}

func (a *alunoService) GetAll(ctx context.Context, page uint32) ([]model.Aluno, error) {
	return []model.Aluno{}, nil
}

func (a *alunoService) Delete(ctx context.Context, id uint64) (bool, error) {
	return false, nil
}

func (a *alunoService) StatusService(ctx context.Context) (bool, error) {
	return false, nil
}
