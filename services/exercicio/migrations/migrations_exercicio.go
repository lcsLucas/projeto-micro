package migrations

import (
	"context"

	"github.com/lcslucas/projeto-micro/services/exercicio/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ExecMigrationExercicios(ctx context.Context, database string, clientMongo *mongo.Client) error {
	var err error

	db := clientMongo.Database(database)

	collecNames, err := db.ListCollectionNames(ctx, bson.D{{Key: "name", Value: "exercicios"}})

	if err == nil && len(collecNames) > 0 { // colection exercicios já existe na base de dados
		return nil
	}

	exerciciosAleatorios := []interface{}{
		model.Exercicio{
			ID:        1,
			Nome:      "EXERCÍCIOS SOBRE DISTÂNCIA ENTRE DOIS PONTOS",
			Descricao: "Calcule a distância entre os pontos A e B, sabendo que suas coordenadas são A (2,5) e B (– 5, – 2).",
			Materia:   "Matemática",
			Ativo:     true,
		},
		model.Exercicio{
			ID:        2,
			Nome:      "EXERCÍCIOS SOBRE DISTÂNCIA ENTRE DOIS PONTOS",
			Descricao: "Calcule o valor da coordenada x do ponto A (x,2) sabendo que a distância entre A e B (4,8) é 10.",
			Materia:   "Matemática",
			Ativo:     true,
		},
		model.Exercicio{
			ID:        3,
			Nome:      "EXERCÍCIOS SOBRE AS APLICAÇÕES DE UMA FUNÇÃO DE 1º GRAU",
			Descricao: "(Fuvest – SP) Determine a função que representa o valor a ser pago após um desconto de 3%% sobre o valor x de uma mercadoria.",
			Materia:   "Matemática",
			Ativo:     true,
		},
		model.Exercicio{
			ID:        4,
			Nome:      "EXERCÍCIOS SOBRE OS CLIMAS NO MUNDO",
			Descricao: "Existem vários tipos de clima no mundo, que variam conforme as zonas climáticas da Terra. Portanto, apresente os aspectos que influenciam na caracterização do clima de um determinado local.",
			Materia:   "Geografia",
			Ativo:     true,
		},
		model.Exercicio{
			ID:        5,
			Nome:      "EXERCÍCIOS SOBRE O CARVÃO MINERAL",
			Descricao: "Como ocorre o processo de formação do carvão mineral?",
			Materia:   "Geografia",
			Ativo:     true,
		},
		model.Exercicio{
			ID:        6,
			Nome:      "EXERCÍCIOS SOBRE O CARVÃO MINERAL",
			Descricao: "Apesar da descoberta e desenvolvimento de outras fontes energéticas, o carvão mineral continua sendo bastante utilizado como fonte de energia. Qual a principal utilidade do carvão mineral?",
			Materia:   "Geografia",
			Ativo:     true,
		},
		model.Exercicio{
			ID:        7,
			Nome:      "Exercícios Sobre A Classificação Dos Fungos",
			Descricao: "Os fungos são organismos pertencentes ao Reino Fungi. Indique algumas características desse grupo.",
			Materia:   "Biologia",
			Ativo:     true,
		},
		model.Exercicio{
			ID:        8,
			Nome:      "Exercícios Sobre A Célula Vegetal",
			Descricao: "Qual o nome da organela encontrada na célula vegetal responsável pela germinação de sementes de oleaginosas.",
			Materia:   "Biologia",
			Ativo:     true,
		},
		model.Exercicio{
			ID:        9,
			Nome:      "Exercícios Sobre A Célula Vegetal",
			Descricao: "A célula vegetal é eucarionte, assim como a célula animal, entretanto, apresenta algumas características que permitem diferenciá-la dessa última célula. Quais estruturas são exclusivas da célula vegetal",
			Materia:   "Biologia",
			Ativo:     true,
		},
	}

	collection := db.Collection("exercicios")
	err = collection.Drop(ctx)
	if err != nil {
		return err
	}

	_, err = collection.InsertMany(ctx, exerciciosAleatorios)
	return err
}
