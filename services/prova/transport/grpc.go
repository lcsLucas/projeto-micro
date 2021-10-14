package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	modelA "github.com/lcslucas/projeto-micro/services/aluno/model"
	modelE "github.com/lcslucas/projeto-micro/services/exercicio/model"
	"github.com/lcslucas/projeto-micro/services/prova/endpoints"
	"github.com/lcslucas/projeto-micro/services/prova/model"
	proto "github.com/lcslucas/projeto-micro/services/prova/proto_prova"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type grpcServer struct {
	create        grpctransport.Handler
	alter         grpctransport.Handler
	get           grpctransport.Handler
	getProvaAluno grpctransport.Handler
	getAll        grpctransport.Handler
	delete        grpctransport.Handler
	statusService grpctransport.Handler
}

// NewGrpcServer inicializa um novo servidor gRPC
func NewGrpcServer(ep endpoints.Set) proto.ServiceProvaServer {
	return &grpcServer{
		create: grpctransport.NewServer(
			ep.CreateEndpoint,
			decodeGrpcCreateAlterRequest,
			encodeGrpcCreateAlterResponse,
		),
		alter: grpctransport.NewServer(
			ep.AlterEndpoint,
			decodeGrpcCreateAlterRequest,
			encodeGrpcCreateAlterResponse,
		),
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGrpcGetRequest,
			encodeGrpcGetResponse,
		),
		getProvaAluno: grpctransport.NewServer(
			ep.GetProvaAlunoEndpoint,
			decodeGrpcGetProvaAlunoRequest,
			encodeGrpcGetResponse,
		),
		getAll: grpctransport.NewServer(
			ep.GetAllEndpoint,
			decodeGrpcGetAllRequest,
			encodeGrpcGetAllResponse,
		),
		delete: grpctransport.NewServer(
			ep.DeleteEndpoint,
			decodeGrpcDeleteRequest,
			encodeGrpcDeleteResponse,
		),
		statusService: grpctransport.NewServer(
			ep.StatusServiceEndpoint,
			decodeGrpcStatusServiceRequest,
			encodeGrpcStatusServiceResponse,
		),
	}
}

/* Implementations interfaces methods */

func (g *grpcServer) Create(ctx context.Context, r *proto.CreateAlterRequest) (*proto.CreateAlterResponse, error) {
	_, res, err := g.create.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*proto.CreateAlterResponse), nil
}

func (g *grpcServer) Alter(ctx context.Context, r *proto.CreateAlterRequest) (*proto.CreateAlterResponse, error) {
	_, res, err := g.alter.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*proto.CreateAlterResponse), nil
}

func (g *grpcServer) Get(ctx context.Context, r *proto.GetRequest) (*proto.GetResponse, error) {
	_, res, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*proto.GetResponse), nil
}

func (g *grpcServer) GetProvaAluno(ctx context.Context, r *proto.GetProvaAlunoRequest) (*proto.GetResponse, error) {
	_, res, err := g.getProvaAluno.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*proto.GetResponse), nil
}

func (g *grpcServer) GetAll(ctx context.Context, r *proto.GetAllRequest) (*proto.GetAllResponse, error) {
	_, res, err := g.getAll.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*proto.GetAllResponse), nil
}

func (g *grpcServer) Delete(ctx context.Context, r *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	_, res, err := g.delete.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*proto.DeleteResponse), nil
}

func (g *grpcServer) StatusService(ctx context.Context, r *proto.StatusServiceRequest) (*proto.StatusServiceResponse, error) {
	_, res, err := g.statusService.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*proto.StatusServiceResponse), nil
}

/* Requests */

func decodeGrpcCreateAlterRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.CreateAlterRequest)

	var exes []modelE.Exercicio
	for _, e := range req.Prova.Exercicios {
		exes = append(exes, modelE.Exercicio{
			ID:        e.Id,
			Nome:      e.Nome,
			Descricao: e.Descricao,
			Materia:   e.Materia,
			Ativo:     e.Ativo,
		})
	}

	aluRes := modelA.Aluno{}

	if req.Prova.Aluno != nil {
		aluRes = modelA.Aluno{
			RA:      req.Prova.Aluno.Ra,
			Nome:    req.Prova.Aluno.Nome,
			Email:   req.Prova.Aluno.Email,
			Celular: req.Prova.Aluno.Celular,
		}
	}

	return endpoints.CreateAlterRequest{
		Prova: model.Prova{
			ID:           req.Prova.Id,
			Nome:         req.Prova.Nome,
			DataCadastro: req.Prova.DataCadastro.AsTime(),
			DataInicio:   req.Prova.DataInicio.AsTime(),
			DataFinal:    req.Prova.DataFinal.AsTime(),
			Serie:        req.Prova.Serie,
			Materia:      req.Prova.Materia,
			Bimestre:     uint16(req.Prova.Bimestre),
			Finalizada:   req.Prova.Finalizada,
			Aluno:        aluRes,
			Exercicios:   exes,
		},
	}, nil
}

func decodeGrpcGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.GetRequest)
	return endpoints.GetRequest{
		ID: req.Id,
	}, nil
}

func decodeGrpcGetProvaAlunoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.GetProvaAlunoRequest)
	return endpoints.GetProvaAlunoRequest{
		IDProva: req.IdProva,
		RaAluno: req.RaAluno,
	}, nil
}

func decodeGrpcGetAllRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.GetAllRequest)
	return endpoints.GetAllRequest{
		Page: req.Page,
	}, nil
}

func decodeGrpcDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.DeleteRequest)
	return endpoints.DeleteRequest{
		ID: uint64(req.Id),
	}, nil
}

func decodeGrpcStatusServiceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.StatusServiceRequest{}, nil
}

/* Responses */

func encodeGrpcCreateAlterResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.CreateAlterResponse)
	return &proto.CreateAlterResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func encodeGrpcGetResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.GetResponse)

	tsCad := timestamppb.New(res.Prova.DataCadastro)
	tsIni := timestamppb.New(res.Prova.DataInicio)
	tsFin := timestamppb.New(res.Prova.DataFinal)

	var exes []*proto.Exercicio
	for _, e := range res.Prova.Exercicios {
		exes = append(exes, &proto.Exercicio{
			Id:        e.ID,
			Nome:      e.Nome,
			Descricao: e.Descricao,
			Materia:   e.Materia,
			Ativo:     e.Ativo,
		})
	}

	return &proto.GetResponse{
		Prova: &proto.Prova{
			Id:           res.Prova.ID,
			Nome:         res.Prova.Nome,
			DataCadastro: tsCad,
			DataInicio:   tsIni,
			DataFinal:    tsFin,
			Serie:        res.Prova.Serie,
			Materia:      res.Prova.Materia,
			Bimestre:     uint32(res.Prova.Bimestre),
			Finalizada:   res.Prova.Finalizada,
			Aluno: &proto.Aluno{
				Ra:      res.Prova.Aluno.RA,
				Nome:    res.Prova.Aluno.Nome,
				Email:   res.Prova.Aluno.Email,
				Celular: res.Prova.Aluno.Celular,
			},
			Exercicios: exes,
		},
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func encodeGrpcGetAllResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.GetAllResponse)

	var provas []*proto.Prova

	for _, p := range res.Provas {

		tsCad := timestamppb.New(p.DataCadastro)
		tsIni := timestamppb.New(p.DataInicio)
		tsFin := timestamppb.New(p.DataFinal)

		var exes []*proto.Exercicio
		for _, e := range p.Exercicios {
			exes = append(exes, &proto.Exercicio{
				Id:        e.ID,
				Nome:      e.Nome,
				Descricao: e.Descricao,
				Materia:   e.Materia,
				Ativo:     e.Ativo,
			})
		}

		provas = append(provas, &proto.Prova{
			Id:           p.ID,
			Nome:         p.Nome,
			DataCadastro: tsCad,
			DataInicio:   tsIni,
			DataFinal:    tsFin,
			Serie:        p.Serie,
			Materia:      p.Materia,
			Bimestre:     uint32(p.Bimestre),
			Finalizada:   p.Finalizada,
			Aluno: &proto.Aluno{
				Ra:      p.Aluno.RA,
				Nome:    p.Aluno.Nome,
				Email:   p.Aluno.Email,
				Celular: p.Aluno.Celular,
			},
			Exercicios: exes,
		})

	}

	return &proto.GetAllResponse{
		Provas: provas,
		Status: res.Status,
		Error:  res.Error,
	}, nil

}

func encodeGrpcDeleteResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.DeleteResponse)
	return &proto.DeleteResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func encodeGrpcStatusServiceResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.StatusServiceResponse)
	return &proto.StatusServiceResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}
