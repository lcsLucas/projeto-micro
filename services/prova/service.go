package prova

import (
	"context"

	"github.com/lcslucas/projeto-micro/services/prova/model"
)

/*Service é uma interface do serviço de provas, onde se encontra todos os métodos necessário para implementação dessa interface. */
type Service interface {
	//Create: Cria nova prova
	Create(ctx context.Context, pro model.Prova) (bool, error)
	//Alter: Altera uma prova existente
	Alter(ctx context.Context, pro model.Prova) (bool, error)
	//Get: Obtém uma prova específico
	Get(ctx context.Context, id uint64) (model.Prova, error)
	//GetProvaAluno: Obtém uma prova específico de um aluno
	GetProvaAluno(ctx context.Context, idProva uint64, raAluno string) (model.Prova, error)
	//GetAll: Obtém uma lista de 10 provas do range passado
	GetAll(ctx context.Context, page uint32) ([]model.Prova, error)
	//Delete: Deleta uma prova específico
	Delete(ctx context.Context, id uint64) (bool, error)
	//StatusService: status do serviço, serve para verificar se o serviço está funcionando
	StatusService(ctx context.Context) (bool, error)
}
