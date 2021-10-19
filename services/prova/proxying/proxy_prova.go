package proxying

import (
	"context"

	"github.com/lcslucas/projeto-micro/services/prova"
	"github.com/lcslucas/projeto-micro/services/prova/model"
	"github.com/sony/gobreaker"
)

type proxyingMidleware struct {
	next prova.Service
	cb   *gobreaker.CircuitBreaker
}

func NewProxying(cb *gobreaker.CircuitBreaker, nService prova.Service) prova.Service {
	return &proxyingMidleware{
		cb:   cb,
		next: nService,
	}
}

func (pm proxyingMidleware) Create(ctx context.Context, pro model.Prova) (err error) {
	_, err = pm.cb.Execute(func() (interface{}, error) {
		return nil, pm.next.Create(ctx, pro)
	})

	return
}

func (pm proxyingMidleware) Alter(ctx context.Context, pro model.Prova) (err error) {
	_, err = pm.cb.Execute(func() (interface{}, error) {
		return nil, pm.next.Alter(ctx, pro)
	})

	return
}

func (pm proxyingMidleware) Get(ctx context.Context, id uint64) (output model.Prova, err error) {
	outputInterface, err := pm.cb.Execute(func() (interface{}, error) {
		return pm.next.Get(ctx, id)
	})

	output = outputInterface.(model.Prova)
	return
}

func (pm proxyingMidleware) GetProvaAluno(ctx context.Context, idProva uint64, raAluno string) (output model.Prova, err error) {
	outputInterface, err := pm.cb.Execute(func() (interface{}, error) {
		return pm.next.GetProvaAluno(ctx, idProva, raAluno)
	})

	output = outputInterface.(model.Prova)
	return
}

func (pm proxyingMidleware) GetAll(ctx context.Context, page uint32) (output []model.Prova, err error) {
	outputInterface, err := pm.cb.Execute(func() (interface{}, error) {
		return pm.next.GetAll(ctx, page)
	})

	output = outputInterface.([]model.Prova)
	return
}

func (pm proxyingMidleware) Delete(ctx context.Context, id uint64) (err error) {
	_, err = pm.cb.Execute(func() (interface{}, error) {
		return nil, pm.next.Delete(ctx, id)
	})

	return
}

func (pm proxyingMidleware) StatusService(ctx context.Context) (err error) {
	_, err = pm.cb.Execute(func() (interface{}, error) {
		return nil, pm.next.StatusService(ctx)
	})

	return
}
