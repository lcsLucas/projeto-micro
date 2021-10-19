package instrumentation

import (
	"context"
	"time"

	"github.com/lcslucas/projeto-micro/services/aluno"
	"github.com/lcslucas/projeto-micro/services/aluno/model"
	"github.com/prometheus/client_golang/prometheus"
)

/*
type CountMethods struct {
	CountCreate        prometheus.Counter
	CountAlter         prometheus.Counter
	CountGet           prometheus.Counter
	CountGetAll        prometheus.Counter
	CountDelete        prometheus.Counter
	CountStatusService prometheus.Counter
}
*/

type LatencyMethods struct {
	LatCreate        prometheus.Histogram
	LatAlter         prometheus.Histogram
	LatGet           prometheus.Histogram
	LatGetAll        prometheus.Histogram
	LatDelete        prometheus.Histogram
	LatStatusService prometheus.Histogram
}

type instrumentationMiddleware struct {
	countMethods   *prometheus.CounterVec
	latencyMethods LatencyMethods
	next           aluno.Service
}

func NewInstrumentation(cMethods *prometheus.CounterVec, lMethods LatencyMethods, nService aluno.Service) aluno.Service {
	return &instrumentationMiddleware{
		countMethods:   cMethods,
		latencyMethods: lMethods,
		next:           nService,
	}
}

func (im instrumentationMiddleware) Create(ctx context.Context, alu model.Aluno) (err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("create").Inc()
		im.latencyMethods.LatCreate.Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = im.next.Create(ctx, alu)
	return
}

func (im instrumentationMiddleware) Alter(ctx context.Context, alu model.Aluno) (err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("alter").Inc()
		im.latencyMethods.LatAlter.Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = im.next.Alter(ctx, alu)
	return
}

func (im instrumentationMiddleware) Get(ctx context.Context, ra string) (output model.Aluno, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("get").Inc()
		im.latencyMethods.LatGet.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Get(ctx, ra)
	return
}

func (im instrumentationMiddleware) GetAll(ctx context.Context, page uint32) (output []model.Aluno, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("getAll").Inc()
		im.latencyMethods.LatGetAll.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.GetAll(ctx, page)
	return
}

func (im instrumentationMiddleware) Delete(ctx context.Context, ra string) (err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("delete").Inc()
		im.latencyMethods.LatDelete.Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = im.next.Delete(ctx, ra)
	return
}

func (im instrumentationMiddleware) StatusService(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("statusService").Inc()
		im.latencyMethods.LatStatusService.Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = im.next.StatusService(ctx)
	return
}
