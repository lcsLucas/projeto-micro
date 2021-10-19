package prova

import (
	"context"
	"errors"
	"fmt"

	"os"
	"reflect"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	modelA "github.com/lcslucas/projeto-micro/services/aluno/model"
	protoAluno "github.com/lcslucas/projeto-micro/services/aluno/proto_aluno"
	modelE "github.com/lcslucas/projeto-micro/services/exercicio/model"
	"github.com/lcslucas/projeto-micro/services/exercicio/proto_exercicio"
	"github.com/lcslucas/projeto-micro/services/prova/model"
)

type provaService struct {
	repository model.Repository
	logger     log.Logger
}

type responseGrpc struct {
	response interface{}
	error    error
}

/*NewService cria um novo serviço de prova */
func NewService(rep model.Repository, logger log.Logger) Service {
	return provaService{
		repository: rep,
		logger:     logger,
	}
}

func (p provaService) Create(ctx context.Context, pro model.Prova) (err error) {
	return errors.New("not implemented")
}

func (p provaService) Alter(ctx context.Context, pro model.Prova) (err error) {
	return errors.New("not implemented")
}

func (p provaService) Get(ctx context.Context, id uint64) (model.Prova, error) {
	return model.Prova{}, errors.New("not implemented")
}

func (p provaService) GetProvaAluno(ctx context.Context, idProva uint64, raAluno string) (model.Prova, error) {
	logger := log.With(p.logger)

	prova, err := p.repository.GetProvaAluno(ctx, idProva, raAluno)
	if err != nil {
		return model.Prova{}, err
	}

	var exeIds []uint64

	for _, exe := range prova.ProvaExercicios {
		exeIds = append(exeIds, exe.ExercicioID)
	}

	done := make(chan interface{}, 2)

	// buscar aluno da prova
	go func() {

		strConn := fmt.Sprintf("%s:%s", os.Getenv("ALU_GRPC_HOST"), os.Getenv("ALU_GRPC_PORT"))

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(strConn, grpc.WithInsecure())
		if err != nil {
			//done <- &protoAluno.Aluno{}
			done <- responseGrpc{
				response: &protoAluno.Aluno{},
				error:    err,
			}
			return
		}
		defer conn.Close()

		c := protoAluno.NewServiceAlunoClient(conn)

		req := protoAluno.GetRequest{
			Ra: raAluno,
		}

		response, err := c.Get(ctx, &req)
		if err != nil {
			//level.Info(logger).Log("error", err)
			//done <- &protoAluno.Aluno{}
			done <- responseGrpc{
				response: &protoAluno.Aluno{},
				error:    err,
			}
			return
		}

		if response.Error != "" {
			//level.Info(logger).Log("error", response.Error)
			//done <- &protoAluno.Aluno{}
			done <- responseGrpc{
				response: &protoAluno.Aluno{},
				error:    errors.New(response.Error),
			}
			return
		}

		//done <- response.Aluno
		done <- responseGrpc{
			response: response.Aluno,
			error:    nil,
		}

	}()

	// buscar exercícios da prova
	go func() {

		strConn := fmt.Sprintf("%s:%s", os.Getenv("EXE_GRPC_HOST"), os.Getenv("EXE_GRPC_PORT"))

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(strConn, grpc.WithInsecure())
		if err != nil {
			//level.Info(logger).Log("error", err)
			//done <- []*proto_exercicio.Exercicio{}
			done <- responseGrpc{
				response: []*proto_exercicio.Exercicio{},
				error:    err,
			}
			return
		}
		defer conn.Close()

		c := proto_exercicio.NewServiceExercicioClient(conn)

		req := proto_exercicio.GetSomesRequest{
			Ids: exeIds,
		}

		response, err := c.GetSomes(ctx, &req)
		if err != nil {
			//level.Info(logger).Log("error", err)
			//done <- []*proto_exercicio.Exercicio{}
			done <- responseGrpc{
				response: []*proto_exercicio.Exercicio{},
				error:    err,
			}
			return
		}

		if response.Error != "" {
			//level.Info(logger).Log("error", response.Error)
			//done <- []*proto_exercicio.Exercicio{}
			done <- responseGrpc{
				response: []*proto_exercicio.Exercicio{},
				error:    errors.New(response.Error),
			}
			return
		}

		//done <- response.Exercicios
		done <- responseGrpc{
			response: response.Exercicios,
			error:    nil,
		}
	}()

	// esperando os resultados das goroutines chegarem
	for i := 0; i < 2; i++ {
		select {
		case respDone := <-done:
			var ptAluno *protoAluno.Aluno
			var ptExercicios []*proto_exercicio.Exercicio

			respGRPC := respDone.(responseGrpc)

			if reflect.TypeOf(respGRPC.response) == reflect.TypeOf(ptAluno) {
				prova.Aluno = modelA.Aluno{
					RA:      respGRPC.response.(*protoAluno.Aluno).Ra,
					Nome:    respGRPC.response.(*protoAluno.Aluno).Nome,
					Email:   respGRPC.response.(*protoAluno.Aluno).Email,
					Celular: respGRPC.response.(*protoAluno.Aluno).Celular,
				}
			}
			if reflect.TypeOf(respGRPC.response) == reflect.TypeOf(ptExercicios) {
				var exes []modelE.Exercicio
				for _, e := range respGRPC.response.([]*proto_exercicio.Exercicio) {
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

			if respGRPC.error != nil {
				level.Info(logger).Log("error", respGRPC.error)
			}

		}
	}

	return prova, nil
}

func (p provaService) GetAll(ctx context.Context, page uint32) ([]model.Prova, error) {
	return []model.Prova{}, errors.New("not implemented")
}

func (p provaService) Delete(ctx context.Context, id uint64) (err error) {
	return errors.New("not implemented")
}

func (p provaService) StatusService(ctx context.Context) (err error) {
	return nil
}
