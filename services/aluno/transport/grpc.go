package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/lcslucas/projeto-micro/services/aluno/endpoints"
	"github.com/lcslucas/projeto-micro/services/aluno/model"
	"github.com/lcslucas/projeto-micro/services/aluno/proto"
)

type grpcServer struct {
	create        grpctransport.Handler
	alter         grpctransport.Handler
	get           grpctransport.Handler
	getAll        grpctransport.Handler
	delete        grpctransport.Handler
	statusService grpctransport.Handler
}

// NewGrpcServer inicializa um novo servidor gRPC
func NewGrpcServer(ep endpoints.Set) proto.ServiceAlunoServer {
	return &grpcServer{
		create: grpctransport.NewServer(
			ep.CreateEndpoint,
			decodeGrpcCreateRequest,
			decodeGrpcCreateResponse,
		),
		alter: grpctransport.NewServer(
			ep.AlterEndpoint,
			decodeGrpcAlterRequest,
			decodeGrpcAlterResponse,
		),
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGrpcGetRequest,
			decodeGrpcGetResponse,
		),
		getAll: grpctransport.NewServer(
			ep.GetAllEndpoint,
			decodeGrpcGetAllRequest,
			decodeGrpcGetAllResponse,
		),
		delete: grpctransport.NewServer(
			ep.DeleteEndpoint,
			decodeGrpcDeleteRequest,
			decodeGrpcDeleteResponse,
		),
		statusService: grpctransport.NewServer(
			ep.StatusServiceEndpoint,
			decodeGrpcStatusServiceRequest,
			decodeGrpcStatusServiceResponse,
		),
	}
}

func (g *grpcServer) Create(ctx context.Context, r *proto.CreateAlterRequest) (*proto.CreateAlterResponse, error) {
	_, res, err := g.get.ServeGRPC(ctx, r)
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

/* Request */

func decodeGrpcCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.CreateAlterRequest)
	return endpoints.CreateAlterRequest{
		Aluno: model.Aluno{
			RA:      req.Aluno.Ra,
			Nome:    req.Aluno.Nome,
			Email:   req.Aluno.Email,
			Celular: req.Aluno.Celular,
		},
	}, nil
}

func decodeGrpcAlterRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.CreateAlterRequest)
	return endpoints.CreateAlterRequest{
		Aluno: model.Aluno{
			RA:      req.Aluno.Ra,
			Nome:    req.Aluno.Nome,
			Email:   req.Aluno.Email,
			Celular: req.Aluno.Celular,
		},
	}, nil
}

func decodeGrpcGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.GetRequest)
	return endpoints.GetRequest{
		ID: req.Id,
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
		ID: req.Id,
	}, nil
}

func decodeGrpcStatusServiceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.StatusServiceRequest{}, nil
}

/* Responses */

func decodeGrpcCreateResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.CreateAlterResponse)
	return endpoints.CreateAlterResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func decodeGrpcAlterResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.CreateAlterResponse)
	return endpoints.CreateAlterResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func decodeGrpcGetResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.GetResponse)
	return endpoints.GetResponse{
		Aluno: model.Aluno{
			RA:      res.Aluno.Ra,
			Nome:    res.Aluno.Nome,
			Email:   res.Aluno.Email,
			Celular: res.Aluno.Celular,
		},
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func decodeGrpcGetAllResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.GetAllResponse)
	var alus []model.Aluno
	for _, a := range res.Alunos {
		alu := model.Aluno{
			RA:      a.Ra,
			Nome:    a.Nome,
			Email:   a.Email,
			Celular: a.Celular,
		}
		alus = append(alus, alu)
	}

	return endpoints.GetAllResponse{
		Alunos: alus,
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func decodeGrpcDeleteResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.DeleteResponse)
	return endpoints.DeleteResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func decodeGrpcStatusServiceResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.StatusServiceResponse)
	return endpoints.StatusServiceResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}
