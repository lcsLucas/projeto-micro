package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/lcslucas/projeto-micro/services/aluno/endpoints"
	"github.com/lcslucas/projeto-micro/services/aluno/model"
	proto "github.com/lcslucas/projeto-micro/services/aluno/proto_aluno"
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
			encodeGrpcCreateResponse,
		),
		alter: grpctransport.NewServer(
			ep.AlterEndpoint,
			decodeGrpcAlterRequest,
			encodeGrpcAlterResponse,
		),
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGrpcGetRequest,
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
			decodeGrpcStatusServiceResponse,
		),
	}
}

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
		RA: req.Ra,
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
		RA: req.Ra,
	}, nil
}

func decodeGrpcStatusServiceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.StatusServiceRequest{}, nil
}

/* Responses */

func encodeGrpcCreateResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.CreateAlterResponse)
	return &proto.CreateAlterResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func encodeGrpcAlterResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.CreateAlterResponse)
	return &proto.CreateAlterResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func encodeGrpcGetResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.GetResponse)
	return &proto.GetResponse{
		Aluno: &proto.Aluno{
			Ra:      res.Aluno.RA,
			Nome:    res.Aluno.Nome,
			Email:   res.Aluno.Email,
			Celular: res.Aluno.Celular,
		},
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func encodeGrpcGetAllResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.GetAllResponse)
	var alus []*proto.Aluno
	for _, a := range res.Alunos {
		alus = append(alus, &proto.Aluno{
			Ra:      a.RA,
			Nome:    a.Nome,
			Email:   a.Email,
			Celular: a.Celular,
		})
	}
	return &proto.GetAllResponse{
		Alunos: alus,
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

func decodeGrpcStatusServiceResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(endpoints.StatusServiceResponse)
	return &proto.StatusServiceResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}
