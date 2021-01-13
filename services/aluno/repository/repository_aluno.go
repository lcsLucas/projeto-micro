package repository

import (
	"aluno/model"
	"context"

	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	clientMongo *mongo.Client
	logger      log.Logger
}

func NewRepository(client *mongo.Client, logger log.Logger) model.Repository {
	return &repository{
		clientMongo: client,
		logger:      log.With(logger, "repository", "aluno", "sql"),
	}
}

func (r *repository) Create(ctx context.Context, alu model.Aluno) (bool, error) {
	return false, nil
}

func (r *repository) Alter(ctx context.Context, alu model.Aluno) (bool, error) {
	return false, nil
}

func (r *repository) Get(ctx context.Context, id uint64) (model.Aluno, error) {
	return model.Aluno{}, nil
}

func (r *repository) GetAll(ctx context.Context, page uint32) ([]model.Aluno, error) {
	return []model.Aluno{}, nil
}

func (r *repository) Delete(ctx context.Context, id uint64) (bool, error) {
	return false, nil
}

func (r *repository) StatusService(ctx context.Context) (bool, error) {
	return false, nil
}
