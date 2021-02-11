package exercicio

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/lcslucas/projeto-micro/services/exercicio/model"
)

type exercicioService struct {
	repository model.Repository
	logger     log.Logger
}

/*NewService cria um novo serviço de exercicio */
func NewService(rep model.Repository, logger log.Logger) Service {
	return &exercicioService{
		repository: rep,
		logger:     logger,
	}
}

func (e exercicioService) Create(ctx context.Context, exe model.Exercicio) (bool, error) {
	logger := log.With(e.logger, "method", "Create")
	level.Info(logger).Log("msg", fmt.Sprintf("Criando o registro: %v", exe))

	return false, errors.New("Not implemented")
}

func (e exercicioService) Alter(ctx context.Context, exe model.Exercicio) (bool, error) {
	logger := log.With(e.logger, "method", "Alter")
	level.Info(logger).Log("msg", fmt.Sprintf("Alterando o registro: %v", exe))

	return false, errors.New("Not implemented")
}

func (e exercicioService) Get(ctx context.Context, id uint64) (model.Exercicio, error) {
	logger := log.With(e.logger, "method", "Get")
	level.Info(logger).Log("msg", fmt.Sprintf("Buscando o registro com ID: %d", id))

	return model.Exercicio{}, errors.New("Not implemented")
}

func (e exercicioService) GetSomes(ctx context.Context, ids []uint64) ([]model.Exercicio, error) {
	logger := log.With(e.logger, "method", "GetSomes")
	level.Info(logger).Log("msg", fmt.Sprintf("Buscando os exercícios com ids: %v", ids))

	exeRes, err := e.repository.GetSomes(ctx, ids)
	if err != nil {
		level.Error(logger).Log("error", err)
		return []model.Exercicio{}, err
	}

	return exeRes, nil
}

func (e exercicioService) Delete(ctx context.Context, id uint64) (bool, error) {
	logger := log.With(e.logger, "method", "Delete")
	level.Info(logger).Log("msg", fmt.Sprintf("Deletando o registro ID: %d", id))

	return false, errors.New("Not implemented")
}

func (e exercicioService) StatusService(ctx context.Context) (bool, error) {
	logger := log.With(e.logger, "method", "StatusService")
	level.Info(logger).Log("msg", fmt.Sprint("Status do serviço: OK"))

	return true, nil
}
