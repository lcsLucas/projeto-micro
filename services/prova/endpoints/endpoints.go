package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/lcslucas/projeto-micro/services/prova"
	"github.com/lcslucas/projeto-micro/services/prova/model"
	"github.com/lcslucas/projeto-micro/utils"
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
		CreateEndpoint:        utils.MakeRateLimit(MakeCreateEndpoint(s), time.Second, 100),
		AlterEndpoint:         utils.MakeRateLimit(MakeAlterEndpoint(s), time.Second, 100),
		GetEndpoint:           utils.MakeRateLimit(MakeGetEndpoint(s), time.Second, 100),
		GetProvaAlunoEndpoint: utils.MakeRateLimit(MakeGetProvaAlunoEndpoint(s), time.Second, 100),
		GetAllEndpoint:        utils.MakeRateLimit(MakeGetAllEndpoint(s), time.Second, 100),
		DeleteEndpoint:        utils.MakeRateLimit(MakeDeleteEndpoint(s), time.Second, 100),
		StatusServiceEndpoint: utils.MakeRateLimit(MakeStatusServiceEndpoint(s), time.Second, 100),
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
