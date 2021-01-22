package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/lcslucas/projeto-micro/services/exercicio/endpoints"
	"github.com/lcslucas/projeto-micro/services/exercicio/model"
	"github.com/lcslucas/projeto-micro/services/exercicio/proto"
)

type grpcServer struct {
	create        grpctransport.Handler
	alter         grpctransport.Handler
	get           grpctransport.Handler
	getSomes      grpctransport.Handler
	delete        grpctransport.Handler
	statusService grpctransport.Handler
}

// NewGrpcServer inicializa um novo servidor gRPC
func NewGrpcServer(ep endpoints.Set) proto.ServiceExercicioServer {
	return &grpcServer{
		create: grpctransport.NewServer(
			ep.CreateEndpoint,
			decodeGrpcCreateAlterRequest,
			decodeGrpcCreateAlterResponse,
		),
		alter: grpctransport.NewServer(
			ep.AlterEndpoint,
			decodeGrpcCreateAlterRequest,
			decodeGrpcCreateAlterResponse,
		),
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGrpcGetRequest,
			decodeGrpcGetResponse,
		),
		getSomes: grpctransport.NewServer(
			ep.GetSomesEndpoint,
			decodeGrpcGetSomesRequest,
			decodeGrpcGetSomesResponse,
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

func (g *grpcServer) GetSomes(ctx context.Context, r *proto.GetSomesRequest) (*proto.GetSomesResponse, error) {
	_, res, err := g.getSomes.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return res.(*proto.GetSomesResponse), nil
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
	return endpoints.CreateAlterRequest{
		Exercicio: model.Exercicio{
			ID:        req.Exercicio.Id,
			Nome:      req.Exercicio.Nome,
			Descricao: req.Exercicio.Descricao,
			Materia:   req.Exercicio.Materia,
			Ativo:     req.Exercicio.Ativo,
		},
	}, nil
}

func decodeGrpcGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.GetRequest)
	return endpoints.GetRequest{
		ID: req.Id,
	}, nil
}

func decodeGrpcGetSomesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.GetSomesRequest)
	return endpoints.GetSomesRequest{
		Ids: req.Ids,
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

func decodeGrpcCreateAlterResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.CreateAlterResponse)
	return &proto.CreateAlterResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func decodeGrpcGetResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.GetResponse)
	return &proto.GetResponse{
		Exercicio: &proto.Exercicio{
			Id:        res.Exercicio.ID,
			Nome:      res.Exercicio.Nome,
			Descricao: res.Exercicio.Descricao,
			Materia:   res.Exercicio.Materia,
			Ativo:     res.Exercicio.Ativo,
		},
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func decodeGrpcGetSomesResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.GetSomesResponse)
	var exes []*proto.Exercicio
	for _, e := range res.Exercicios {
		exes = append(exes, &proto.Exercicio{
			Id:        e.ID,
			Nome:      e.Nome,
			Descricao: e.Descricao,
			Materia:   e.Materia,
			Ativo:     e.Ativo,
		})
	}

	return &proto.GetSomesResponse{
		Exercicios: exes,
		Status:     res.Status,
		Error:      res.Error,
	}, nil
}

func decodeGrpcDeleteResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.DeleteResponse)
	return &proto.DeleteResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func decodeGrpcStatusServiceResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.StatusServiceResponse)
	return &proto.StatusServiceResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}
