package instrumentation

import (
	"context"
	"time"

	"github.com/lcslucas/projeto-micro/services/aluno"
	"github.com/lcslucas/projeto-micro/services/aluno/model"
	"github.com/prometheus/client_golang/prometheus"
)

type instrumentationMiddleware struct {
	requestCount   prometheus.Counter
	requestLatency prometheus.Histogram
	next           aluno.Service
}

func NewInstrumentation(rCount prometheus.Counter, rLatency prometheus.Histogram, nService aluno.Service) aluno.Service {
	return &instrumentationMiddleware{
		requestCount:   rCount,
		requestLatency: rLatency,
		next:           nService,
	}
}

func (im instrumentationMiddleware) Create(ctx context.Context, alu model.Aluno) (output bool, err error) {
	defer func(begin time.Time) {
		//lvs := []string{"method", "create", "error", fmt.Sprint(err != nil)}
		im.requestCount.Add(1)
		im.requestLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Create(ctx, alu)
	return
}

func (im instrumentationMiddleware) Alter(ctx context.Context, alu model.Aluno) (output bool, err error) {
	defer func(begin time.Time) {
		//lvs := []string{"method", "alter", "error", fmt.Sprint(err != nil)}
		im.requestCount.Add(1)
		im.requestLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Alter(ctx, alu)
	return
}

func (im instrumentationMiddleware) Get(ctx context.Context, ra string) (output model.Aluno, err error) {
	defer func(begin time.Time) {
		//lvs := []string{"method", "get", "error", fmt.Sprint(err != nil)}
		im.requestCount.Add(1)
		im.requestLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Get(ctx, ra)
	return
}

func (im instrumentationMiddleware) GetAll(ctx context.Context, page uint32) (output []model.Aluno, err error) {
	defer func(begin time.Time) {
		//lvs := []string{"method", "getAll", "error", fmt.Sprint(err != nil)}
		im.requestCount.Add(1)
		im.requestLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.GetAll(ctx, page)
	return
}

func (im instrumentationMiddleware) Delete(ctx context.Context, ra string) (output bool, err error) {
	defer func(begin time.Time) {
		//lvs := []string{"method", "delete", "error", fmt.Sprint(err != nil)}
		im.requestCount.Add(1)
		im.requestLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Delete(ctx, ra)
	return
}

func (im instrumentationMiddleware) StatusService(ctx context.Context) (output bool, err error) {
	defer func(begin time.Time) {
		//lvs := []string{"method", "statusService", "error", fmt.Sprint(err != nil)}
		im.requestCount.Add(1)
		im.requestLatency.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.StatusService(ctx)
	return
}
