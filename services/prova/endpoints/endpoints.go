package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lcslucas/projeto-micro/services/prova"
	"github.com/lcslucas/projeto-micro/services/prova/model"
)

type Set struct {
	CreateEndpoint        endpoint.Endpoint
	AlterEndpoint         endpoint.Endpoint
	GetEndpoint           endpoint.Endpoint
	GetProvaAlunoEndpoint endpoint.Endpoint
	GetAllEndpoint        endpoint.Endpoint
	DeleteEndpoint        endpoint.Endpoint
	StatusServiceEndpoint endpoint.Endpoint
}

func NewEndpointSet(s prova.Service) Set {
	return Set{
		CreateEndpoint:        MakeCreateEndpoint(s),
		AlterEndpoint:         MakeAlterEndpoint(s),
		GetEndpoint:           MakeGetEndpoint(s),
		GetProvaAlunoEndpoint: MakeGetProvaAlunoEndpoint(s),
		GetAllEndpoint:        MakeGetAllEndpoint(s),
		DeleteEndpoint:        MakeDeleteEndpoint(s),
		StatusServiceEndpoint: MakeStatusServiceEndpoint(s),
	}
}

func MakeCreateEndpoint(s prova.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAlterRequest)
		err := s.Create(ctx, req.Prova)
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

func MakeAlterEndpoint(s prova.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAlterRequest)
		err := s.Alter(ctx, req.Prova)
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

func MakeGetEndpoint(s prova.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		p, err := s.Get(ctx, req.ID)
		if err != nil {
			return GetResponse{
				Prova:  model.Prova{},
				Status: false,
				Error:  err.Error(),
			}, nil
		}

		return GetResponse{
			Prova:  p,
			Status: true,
			Error:  "",
		}, nil
	}
}

func MakeGetProvaAlunoEndpoint(s prova.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProvaAlunoRequest)
		p, err := s.GetProvaAluno(ctx, req.IDProva, req.RaAluno)
		if err != nil {
			return GetResponse{
				Prova:  model.Prova{},
				Status: false,
				Error:  err.Error(),
			}, nil
		}

		return GetResponse{
			Prova:  p,
			Status: true,
			Error:  "",
		}, nil
	}
}

func MakeGetAllEndpoint(s prova.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllRequest)
		ps, err := s.GetAll(ctx, req.Page)
		if err != nil {
			return GetAllResponse{
				Provas: []model.Prova{},
				Status: false,
				Error:  err.Error(),
			}, nil
		}

		return GetAllResponse{
			Provas: ps,
			Status: true,
			Error:  "",
		}, nil
	}
}

func MakeDeleteEndpoint(s prova.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(ctx, req.ID)
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

func MakeStatusServiceEndpoint(s prova.Service) endpoint.Endpoint {
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
