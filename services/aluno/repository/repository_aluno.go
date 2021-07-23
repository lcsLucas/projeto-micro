package repository

import (
	"context"

	"github.com/lcslucas/projeto-micro/config"
	"github.com/lcslucas/projeto-micro/services/aluno/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const TableName = "alunos"

type repository struct {
	clientMongo *mongo.Client
	configDB    config.ConfigDB
}

//NewRepository cria um novo repositório para o serviço
func NewRepository(client *mongo.Client, configDB config.ConfigDB) model.Repository {
	return &repository{
		clientMongo: client,
		configDB:    configDB,
	}
}

func (r *repository) Create(ctx context.Context, alu model.Aluno) (bool, error) {
	return false, nil
}

func (r *repository) Alter(ctx context.Context, alu model.Aluno) (bool, error) {
	return false, nil
}

func (r *repository) Get(ctx context.Context, ra string) (model.Aluno, error) {
	db := r.clientMongo.Database(r.configDB.DBName)
	collection := db.Collection(TableName)

	filter := bson.M{"ra": ra}
	a := model.Aluno{}

	err := collection.FindOne(ctx, filter).Decode(&a)
	if err != nil {
		return model.Aluno{}, err
	}

	return a, nil
}

func (r *repository) GetAll(ctx context.Context, page uint32) ([]model.Aluno, error) {
	return []model.Aluno{}, nil
}

func (r *repository) Delete(ctx context.Context, ra string) (bool, error) {
	return false, nil
}

func (r *repository) StatusService(ctx context.Context) (bool, error) {
	return false, nil
}
