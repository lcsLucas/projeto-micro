package endpoints

import (
	"context"

	"github.com/lcslucas/projeto-micro/services/aluno/model"

	"github.com/go-kit/kit/endpoint"
	"github.com/lcslucas/projeto-micro/services/aluno"
)

type Set struct {
	CreateEndpoint        endpoint.Endpoint
	AlterEndpoint         endpoint.Endpoint
	GetEndpoint           endpoint.Endpoint
	GetAllEndpoint        endpoint.Endpoint
	DeleteEndpoint        endpoint.Endpoint
	StatusServiceEndpoint endpoint.Endpoint
}

func NewEndpointSet(s aluno.Service) Set {
	return Set{
		CreateEndpoint:        MakeCreateEndpoint(s),
		AlterEndpoint:         MakeAlterEndpoint(s),
		GetEndpoint:           MakeGetEndpoint(s),
		GetAllEndpoint:        MakeGetAllEndpoint(s),
		DeleteEndpoint:        MakeDeleteEndpoint(s),
		StatusServiceEndpoint: MakeStatusServiceEndpoint(s),
	}
}

func MakeCreateEndpoint(s aluno.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAlterRequest)
		err := s.Create(ctx, req.Aluno)
		if err != nil {
			return CreateAlterResponse{
				Status: false,
				Error:  err.Error(),
			}, nil
		}

		return CreateAlterResponse{
			Status: true,
			Error:  "",
		}, nil
	}
}

func MakeAlterEndpoint(s aluno.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAlterRequest)
		err := s.Alter(ctx, req.Aluno)
		if err != nil {
			return CreateAlterResponse{
				Status: false,
				Error:  err.Error(),
			}, nil
		}

		return CreateAlterResponse{
			Status: true,
			Error:  "",
		}, nil
	}
}

func MakeGetEndpoint(s aluno.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		alu, err := s.Get(ctx, req.RA)
		if err != nil {
			return GetResponse{
				Aluno:  model.Aluno{},
				Status: false,
				Error:  err.Error(),
			}, nil
		}

		return GetResponse{
			Aluno:  alu,
			Status: true,
			Error:  "",
		}, nil
	}
}

func MakeGetAllEndpoint(s aluno.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllRequest)
		alus, err := s.GetAll(ctx, req.Page)
		if err != nil {
			return GetAllResponse{
				Alunos: []model.Aluno{},
				Status: false,
				Error:  err.Error(),
			}, nil
		}

		return GetAllResponse{
			Alunos: alus,
			Status: true,
			Error:  "",
		}, nil
	}
}

func MakeDeleteEndpoint(s aluno.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(ctx, req.RA)
		if err != nil {
			return DeleteResponse{
				Status: false,
				Error:  err.Error(),
			}, nil
		}

		return DeleteResponse{
			Status: true,
			Error:  "",
		}, nil
	}
}

func MakeStatusServiceEndpoint(s aluno.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := s.StatusService(ctx)
		if err != nil {
			return StatusServiceResponse{
				Status: false,
				Error:  err.Error(),
			}, nil
		}

		return StatusServiceResponse{
			Status: true,
			Error:  "",
		}, nil
	}
}
