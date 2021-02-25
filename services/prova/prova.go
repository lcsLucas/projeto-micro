package prova

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"

	modelA "github.com/lcslucas/projeto-micro/services/aluno/model"
	protoAluno "github.com/lcslucas/projeto-micro/services/aluno/proto_aluno"
	modelE "github.com/lcslucas/projeto-micro/services/exercicio/model"
	"github.com/lcslucas/projeto-micro/services/exercicio/proto_exercicio"
	protoExercicio "github.com/lcslucas/projeto-micro/services/exercicio/proto_exercicio"
	"github.com/lcslucas/projeto-micro/services/prova/model"
)

type provaService struct {
	repository model.Repository
	logger     log.Logger
}

/*NewService cria um novo serviço de prova */
func NewService(rep model.Repository, logger log.Logger) Service {
	return provaService{
		repository: rep,
		logger:     logger,
	}
}

func (p provaService) Create(ctx context.Context, pro model.Prova) (bool, error) {
	logger := log.With(p.logger, "method", "Create")
	level.Info(logger).Log("msg", fmt.Sprintf("Criando o registro: %v", pro))

	return false, errors.New("Not implemented")
}

func (p provaService) Alter(ctx context.Context, pro model.Prova) (bool, error) {
	logger := log.With(p.logger, "method", "Alter")
	level.Info(logger).Log("msg", fmt.Sprintf("Alterando o registro: %v", pro))

	return false, errors.New("Not implemented")
}

func (p provaService) Get(ctx context.Context, id uint64) (model.Prova, error) {
	logger := log.With(p.logger, "method", "Get")
	level.Info(logger).Log("msg", fmt.Sprintf("Buscando o registro com ID: %d", id))

	return model.Prova{}, errors.New("Not implemented")
}

func (p provaService) GetProvaAluno(ctx context.Context, idProva uint64, raAluno string) (model.Prova, error) {
	logger := log.With(p.logger, "method", "GetProvaAluno")
	level.Info(logger).Log("msg", fmt.Sprintf("Buscando o registro com ID: %d do RA: %s", idProva, raAluno))

	prova, err := p.repository.GetProvaAluno(ctx, idProva, raAluno)
	if err != nil {
		level.Info(logger).Log("error", err)
		return model.Prova{}, err
	}

	var exeIds []uint64

	for _, exe := range prova.ProvaExercicios {
		exeIds = append(exeIds, exe.ExercicioID)
	}

	done := make(chan interface{}, 2)

	// buscar aluno da prova
	go func() {
		level.Info(logger).Log("msg", "Iniciando a comunicação com o serviço 'Aluno'...")
		defer level.Info(logger).Log("msg", "Finalizada a comunicação com o serviço 'Aluno'")

		strConn := fmt.Sprintf("%s:%s", os.Getenv("ALU_GRPC_HOST"), os.Getenv("ALU_GRPC_PORT"))

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(strConn, grpc.WithInsecure())
		if err != nil {
			level.Info(logger).Log("error", err)
			done <- &protoAluno.Aluno{}
			return
		}
		defer conn.Close()

		c := protoAluno.NewServiceAlunoClient(conn)

		req := protoAluno.GetRequest{
			Ra: raAluno,
		}

		response, err := c.Get(ctx, &req)
		if err != nil {
			level.Info(logger).Log("error", err)
			done <- &protoAluno.Aluno{}
			return
		}

		if response.Error != "" {
			level.Info(logger).Log("error", response.Error)
			done <- &protoAluno.Aluno{}
			return
		}

		done <- response.Aluno

	}()

	// buscar exercícios da prova
	go func() {
		level.Info(logger).Log("msg", "Iniciando a comunicação com o serviço 'Exercício'...")
		defer level.Info(logger).Log("msg", "Finalizada a comunicação com o serviço 'Exercício'")

		strConn := fmt.Sprintf("%s:%s", os.Getenv("EXE_GRPC_HOST"), os.Getenv("EXE_GRPC_PORT"))

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(strConn, grpc.WithInsecure())
		if err != nil {
			level.Info(logger).Log("error", err)
			done <- []*proto_exercicio.Exercicio{}
			return
		}
		defer conn.Close()

		c := protoExercicio.NewServiceExercicioClient(conn)

		req := protoExercicio.GetSomesRequest{
			Ids: exeIds,
		}

		response, err := c.GetSomes(ctx, &req)
		if err != nil {
			level.Info(logger).Log("error", err)
			done <- []*proto_exercicio.Exercicio{}
			return
		}

		if response.Error != "" {
			level.Info(logger).Log("error", response.Error)
			done <- []*proto_exercicio.Exercicio{}
			return
		}

		done <- response.Exercicios
	}()

	// esperando os resultados das goroutines chegarem
	for i := 0; i < 2; i++ {
		select {
		case respDone := <-done:
			var ptAluno *protoAluno.Aluno
			var ptExercicios []*proto_exercicio.Exercicio
			if reflect.TypeOf(respDone) == reflect.TypeOf(ptAluno) {
				prova.Aluno = modelA.Aluno{
					RA:      respDone.(*protoAluno.Aluno).Ra,
					Nome:    respDone.(*protoAluno.Aluno).Nome,
					Email:   respDone.(*protoAluno.Aluno).Email,
					Celular: respDone.(*protoAluno.Aluno).Celular,
				}
			}
			if reflect.TypeOf(respDone) == reflect.TypeOf(ptExercicios) {
				var exes []modelE.Exercicio
				for _, e := range respDone.([]*proto_exercicio.Exercicio) {
					exes = append(exes, modelE.Exercicio{
						ID:        e.Id,
						Nome:      e.Nome,
						Descricao: e.Descricao,
						Materia:   e.Materia,
						Ativo:     e.Ativo,
					})
				}
				prova.Exercicios = exes
			}
		}
	}

	return prova, nil
}

func (p provaService) GetAll(ctx context.Context, page uint32) ([]model.Prova, error) {
	logger := log.With(p.logger, "method", "GetAll")
	level.Info(logger).Log("msg", fmt.Sprintf("Buscando registros da página %d", page))

	/*
		allProvas, err := p.repository.GetAll(ctx, 1)
		if err != nil {
			level.Error(logger).Log("error", err)
			return []model.Prova{}, err
		}

		fmt.Println()
		fmt.Println(allProvas)
		fmt.Println()
	*/

	return []model.Prova{}, errors.New("Not implemented")
}

func (p provaService) Delete(ctx context.Context, id uint64) (bool, error) {
	logger := log.With(p.logger, "method", "Delete")
	level.Info(logger).Log("msg", fmt.Sprintf("Deletando o registro ID: %d", id))

	return false, errors.New("Not implemented")
}

func (p provaService) StatusService(ctx context.Context) (bool, error) {
	logger := log.With(p.logger, "method", "StatusService")
	level.Info(logger).Log("msg", fmt.Sprint("Status do serviço: OK"))

	return true, nil
}
