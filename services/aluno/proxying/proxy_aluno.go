package proxying

import (
	"context"

	"github.com/lcslucas/projeto-micro/services/aluno"
	"github.com/lcslucas/projeto-micro/services/aluno/model"
	"github.com/sony/gobreaker"
)

type proxyingMidleware struct {
	next aluno.Service
	cb   *gobreaker.CircuitBreaker
}

func NewProxying(cb *gobreaker.CircuitBreaker, nService aluno.Service) aluno.Service {
	return &proxyingMidleware{
		cb:   cb,
		next: nService,
	}
}

func (pm proxyingMidleware) Create(ctx context.Context, alu model.Aluno) (err error) {
	_, err = pm.cb.Execute(func() (interface{}, error) {
		return nil, pm.next.Create(ctx, alu)
	})

	return
}
func (pm proxyingMidleware) Alter(ctx context.Context, alu model.Aluno) (err error) {
	_, err = pm.cb.Execute(func() (interface{}, error) {
		return nil, pm.next.Alter(ctx, alu)
	})

	return
}
func (pm proxyingMidleware) Get(ctx context.Context, ra string) (output model.Aluno, err error) {
	outputInterface, err := pm.cb.Execute(func() (interface{}, error) {
		return pm.next.Get(ctx, ra)
	})

	output = outputInterface.(model.Aluno)
	return
}
func (pm proxyingMidleware) GetAll(ctx context.Context, page uint32) (output []model.Aluno, err error) {
	outputInterface, err := pm.cb.Execute(func() (interface{}, error) {
		return pm.next.GetAll(ctx, page)
	})

	output = outputInterface.([]model.Aluno)
	return
}
func (pm proxyingMidleware) Delete(ctx context.Context, ra string) (err error) {
	_, err = pm.cb.Execute(func() (interface{}, error) {
		return nil, pm.next.Delete(ctx, ra)
	})

	return
}
func (pm proxyingMidleware) StatusService(ctx context.Context) (err error) {
	_, err = pm.cb.Execute(func() (interface{}, error) {
		return nil, pm.next.StatusService(ctx)
	})

	return
}
