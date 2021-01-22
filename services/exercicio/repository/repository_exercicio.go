package repository

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/lcslucas/projeto-micro/config"
	"github.com/lcslucas/projeto-micro/services/exercicio/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const tableName = "exercicios"

type repository struct {
	clientMongo *mongo.Client
	logger      log.Logger
	configDB    config.ConfigDB
}

//NewRepository cria um novo repositório para o serviço
func NewRepository(client *mongo.Client, logger log.Logger, configDB config.ConfigDB) model.Repository {
	return &repository{
		clientMongo: client,
		logger:      log.With(logger, "repository", "exercicio", "sql"),
		configDB:    configDB,
	}
}

func (r *repository) Create(ctx context.Context, exe model.Exercicio) (bool, error) {
	return false, nil
}

func (r *repository) Alter(ctx context.Context, exe model.Exercicio) (bool, error) {
	return false, nil
}

func (r *repository) Get(ctx context.Context, id uint64) (model.Exercicio, error) {
	return model.Exercicio{}, nil
}

func (r *repository) GetSomes(ctx context.Context, ids []uint64) ([]model.Exercicio, error) {
	db := r.clientMongo.Database(r.configDB.DBName)
	collection := db.Collection(tableName)

	var pIds []interface{}

	for _, v := range ids {
		pIds = append(pIds, bson.D{{Key: "id", Value: v}})
	}

	filter := bson.D{
		{Key: "$or", Value: pIds},
	}

	exes := []model.Exercicio{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return []model.Exercicio{}, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var e model.Exercicio
		if err = cursor.Decode(&e); err != nil {
			return []model.Exercicio{}, err
		}

		exes = append(exes, e)
	}

	return exes, nil
}

func (r *repository) Delete(ctx context.Context, id uint64) (bool, error) {
	return false, nil
}

func (r *repository) StatusService(ctx context.Context) (bool, error) {
	return false, nil
}
