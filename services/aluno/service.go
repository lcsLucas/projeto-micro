package aluno

import (
	"context"

	"github.com/lcslucas/projeto-micro/services/aluno/model"
)

/*Service é uma interface do serviço de alunos, onde se encontra todos os métodos necessário para implementação dessa interface. */
type Service interface {
	//Create: Cria novo aluno
	Create(ctx context.Context, alu model.Aluno) (bool, error)
	//Alter: Altera um aluno existente
	Alter(ctx context.Context, alu model.Aluno) (bool, error)
	//Get: Obtém um aluno específico
	Get(ctx context.Context, ra string) (model.Aluno, error)
	//GetAll: Obtém uma lista de 10 alunos iniciado do parametro page passado
	GetAll(ctx context.Context, page uint32) ([]model.Aluno, error)
	//Delete: Deleta um aluno específico
	Delete(ctx context.Context, ra string) (bool, error)
	//StatusService: status do serviço, serve para verificar se o serviço está funcionando
	StatusService(ctx context.Context) (bool, error)
}
