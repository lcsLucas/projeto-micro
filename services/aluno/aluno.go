package aluno

import (
	"context"
	"fmt"

	"github.com/lcslucas/projeto-micro/services/aluno/model"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type alunoService struct {
	repository model.Repository
	logger     log.Logger
}

/*NewService cria um novo serviço de aluno */
func NewService(rep model.Repository, logger log.Logger) Service {
	return &alunoService{
		repository: rep,
		logger:     logger,
	}
}

func (a *alunoService) Create(ctx context.Context, alu model.Aluno) (bool, error) {
	logger := log.With(a.logger, "method", "Create")
	level.Info(logger).Log("msg", fmt.Sprintf("Criando o registro: %v", alu))

	return false, nil
}

func (a *alunoService) Alter(ctx context.Context, alu model.Aluno) (bool, error) {
	logger := log.With(a.logger, "method", "Alter")
	level.Info(logger).Log("msg", fmt.Sprintf("Alterando o registro: %v", alu))

	return false, nil
}

func (a *alunoService) Get(ctx context.Context, ra string) (model.Aluno, error) {
	logger := log.With(a.logger, "method", "Get")
	level.Info(logger).Log("msg", fmt.Sprintf("Buscando o registro com RA: %s", ra))

	aluRes, err := a.repository.Get(ctx, ra)
	if err != nil {
		level.Error(logger).Log("error", err)
		return model.Aluno{}, err
	}

	level.Info(logger).Log("msg", fmt.Sprintf("%v", aluRes))
	return aluRes, nil
}

func (a *alunoService) GetAll(ctx context.Context, page uint32) ([]model.Aluno, error) {
	logger := log.With(a.logger, "method", "GetAll")
	level.Info(logger).Log("msg", fmt.Sprintf("Buscando uma lista de registro da página: %d", page))

	return []model.Aluno{}, nil
}

func (a *alunoService) Delete(ctx context.Context, ra string) (bool, error) {
	logger := log.With(a.logger, "method", "Delete")
	level.Info(logger).Log("msg", fmt.Sprintf("Deletando o registro RA: %s", ra))

	return false, nil
}

func (a *alunoService) StatusService(ctx context.Context) (bool, error) {
	logger := log.With(a.logger, "method", "StatusService")
	level.Info(logger).Log("msg", fmt.Sprint("Status do serviço: OK"))

	return false, nil
}
