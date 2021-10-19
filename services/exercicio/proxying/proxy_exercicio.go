package proxying

import (
	"context"

	"github.com/lcslucas/projeto-micro/services/exercicio"
	"github.com/lcslucas/projeto-micro/services/exercicio/model"
	"github.com/sony/gobreaker"
)

type proxyingMidleware struct {
	next exercicio.Service
	cb   *gobreaker.CircuitBreaker
}

func NewProxying(cb *gobreaker.CircuitBreaker, nService exercicio.Service) exercicio.Service {
	return &proxyingMidleware{
		cb:   cb,
		next: nService,
	}
}

func (pm proxyingMidleware) Create(ctx context.Context, exe model.Exercicio) (err error) {
	_, err = pm.cb.Execute(func() (interface{}, error) {
		return nil, pm.next.Create(ctx, exe)
	})

	return
}
func (pm proxyingMidleware) Alter(ctx context.Context, exe model.Exercicio) (err error) {
	_, err = pm.cb.Execute(func() (interface{}, error) {
		return nil, pm.next.Alter(ctx, exe)
	})

	return
}
func (pm proxyingMidleware) Get(ctx context.Context, id uint64) (output model.Exercicio, err error) {
	outputInterface, err := pm.cb.Execute(func() (interface{}, error) {
		return pm.next.Get(ctx, id)
	})

	output = outputInterface.(model.Exercicio)
	return
}
func (pm proxyingMidleware) GetSomes(ctx context.Context, ids []uint64) (output []model.Exercicio, err error) {
	outputInterface, err := pm.cb.Execute(func() (interface{}, error) {
		return pm.next.GetSomes(ctx, ids)
	})

	output = outputInterface.([]model.Exercicio)
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
