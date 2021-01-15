package migrations

import (
	"context"

	"github.com/lcslucas/projeto-micro/services/aluno/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func ExecMigrationAlunos(ctx context.Context, database string, clientMongo *mongo.Client) error {
	var err error

	db := clientMongo.Database(database)

	pessoas_aleatorias := []interface{}{
		model.Aluno{
			RA:      "16.128.614-8",
			Nome:    "Nicolas Arthur Cauê Silveira",
			Email:   "nicolasarthurcauesilveira@email.com",
			Celular: "(27) 99888-5555",
		},
		model.Aluno{
			RA:      "10.846.074-5",
			Nome:    "Mariane Raquel da Mata",
			Email:   "mariamarianeraqueldamata@email.com.br",
			Celular: "(28) 99777-6666",
		},
		model.Aluno{
			RA:      "40.759.343-3",
			Nome:    "Carla Renata Baptista",
			Email:   "ccarlarenatabaptista@email.com.br",
			Celular: "(87) 99999-3333",
		},
		model.Aluno{
			RA:      "40.738.017-6",
			Nome:    "Hugo Francisco Lima",
			Email:   "hhugofranciscolima@email.biz",
			Celular: "(92) 98222-2222",
		},
		model.Aluno{
			RA:      "35.329.394-5",
			Nome:    "Kaique Raimundo Marcelo Melo",
			Email:   "kaiqueraimundomarcelomelo-89@email.com.br",
			Celular: "(12) 91111-1111",
		},
	}

	collection := db.Collection("alunos")
	_, err = collection.InsertMany(ctx, pessoas_aleatorias)
	if err != nil {
		return err
	}

	return nil
}
