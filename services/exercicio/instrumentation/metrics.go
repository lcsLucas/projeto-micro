package instrumentation

import (
	"context"
	"time"

	"github.com/lcslucas/projeto-micro/services/exercicio"
	"github.com/lcslucas/projeto-micro/services/exercicio/model"
	"github.com/prometheus/client_golang/prometheus"
)

type LatencyMethods struct {
	LatCreate        prometheus.Histogram
	LatAlter         prometheus.Histogram
	LatGet           prometheus.Histogram
	LatGetSomes      prometheus.Histogram
	LatDelete        prometheus.Histogram
	LatStatusService prometheus.Histogram
}

type instrumentationMiddleware struct {
	countMethods   *prometheus.CounterVec
	latencyMethods LatencyMethods
	next           exercicio.Service
}

func NewInstrumentation(cMethods *prometheus.CounterVec, lMethods LatencyMethods, nService exercicio.Service) exercicio.Service {
	return &instrumentationMiddleware{
		countMethods:   cMethods,
		latencyMethods: lMethods,
		next:           nService,
	}
}

func (im instrumentationMiddleware) Create(ctx context.Context, exe model.Exercicio) (output bool, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("create").Inc()
		im.latencyMethods.LatCreate.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Create(ctx, exe)
	return
}

func (im instrumentationMiddleware) Alter(ctx context.Context, exe model.Exercicio) (output bool, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("alter").Inc()
		im.latencyMethods.LatAlter.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Alter(ctx, exe)
	return
}

func (im instrumentationMiddleware) Get(ctx context.Context, id uint64) (output model.Exercicio, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("get").Inc()
		im.latencyMethods.LatGet.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Get(ctx, id)
	return
}

func (im instrumentationMiddleware) GetSomes(ctx context.Context, ids []uint64) (output []model.Exercicio, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("getSomes").Inc()
		im.latencyMethods.LatGetSomes.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.GetSomes(ctx, ids)
	return
}

func (im instrumentationMiddleware) Delete(ctx context.Context, id uint64) (output bool, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("delete").Inc()
		im.latencyMethods.LatDelete.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Delete(ctx, id)
	return
}

func (im instrumentationMiddleware) StatusService(ctx context.Context) (output bool, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("statusService").Inc()
		im.latencyMethods.LatStatusService.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.StatusService(ctx)
	return
}
