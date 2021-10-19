package exercicio

import (
	"context"

	"github.com/lcslucas/projeto-micro/services/exercicio/model"
)

/*Service é uma interface do serviço de exercicios, onde se encontra todos os métodos necessário para implementação dessa interface. */
type Service interface {
	//Create: Cria novo exercicio
	Create(ctx context.Context, exe model.Exercicio) error
	//Alter: Altera um exercicio existente
	Alter(ctx context.Context, exe model.Exercicio) error
	//Get: Obtém um exercicio específico
	Get(ctx context.Context, id uint64) (model.Exercicio, error)
	//GetSomes obtém uma lista de exercícios específicos
	GetSomes(ctx context.Context, ids []uint64) ([]model.Exercicio, error)
	//Delete: Deleta um exercicio específico
	Delete(ctx context.Context, id uint64) error
	//StatusService: status do serviço, serve para verificar se o serviço está funcionando
	StatusService(ctx context.Context) error
}
