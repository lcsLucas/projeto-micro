package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/lcslucas/projeto-micro/services/exercicio"
	"github.com/lcslucas/projeto-micro/services/exercicio/model"
	"github.com/lcslucas/projeto-micro/utils"
)

type Set struct {
	CreateEndpoint        endpoint.Endpoint
	AlterEndpoint         endpoint.Endpoint
	GetEndpoint           endpoint.Endpoint
	GetSomesEndpoint      endpoint.Endpoint
	DeleteEndpoint        endpoint.Endpoint
	StatusServiceEndpoint endpoint.Endpoint
}

func NewEndpointSet(s exercicio.Service) Set {
	return Set{
		CreateEndpoint:        utils.MakeRateLimit(MakeCreateEndpoint(s), time.Second, 100),
		AlterEndpoint:         utils.MakeRateLimit(MakeAlterEndpoint(s), time.Second, 100),
		GetEndpoint:           utils.MakeRateLimit(MakeGetEndpoint(s), time.Second, 100),
		GetSomesEndpoint:      utils.MakeRateLimit(MakeGetSomesEndpoint(s), time.Second, 100),
		DeleteEndpoint:        utils.MakeRateLimit(MakeDeleteEndpoint(s), time.Second, 100),
		StatusServiceEndpoint: utils.MakeRateLimit(MakeStatusServiceEndpoint(s), time.Second, 100),
	}
}

func MakeCreateEndpoint(s exercicio.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAlterRequest)
		err := s.Create(ctx, req.Exercicio)
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

func MakeAlterEndpoint(s exercicio.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAlterRequest)
		err := s.Alter(ctx, req.Exercicio)
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

func MakeGetEndpoint(s exercicio.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		exe, err := s.Get(ctx, req.ID)
		if err != nil {
			return GetResponse{
				Exercicio: model.Exercicio{},
				Status:    false,
				Error:     err.Error(),
			}, nil
		}

		return GetResponse{
			Exercicio: exe,
			Status:    true,
			Error:     "",
		}, nil
	}
}

func MakeGetSomesEndpoint(s exercicio.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSomesRequest)
		exes, err := s.GetSomes(ctx, req.Ids)
		if err != nil {
			return GetSomesResponse{
				Exercicios: []model.Exercicio{},
				Status:     false,
				Error:      err.Error(),
			}, nil
		}

		return GetSomesResponse{
			Exercicios: exes,
			Status:     true,
			Error:      "",
		}, nil
	}
}

func MakeDeleteEndpoint(s exercicio.Service) endpoint.Endpoint {
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

func MakeStatusServiceEndpoint(s exercicio.Service) endpoint.Endpoint {
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
