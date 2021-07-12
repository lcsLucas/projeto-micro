package instrumentation

import (
	"context"
	"time"

	"github.com/lcslucas/projeto-micro/services/prova"
	"github.com/lcslucas/projeto-micro/services/prova/model"
	"github.com/prometheus/client_golang/prometheus"
)

type LatencyMethods struct {
	LatCreate        prometheus.Histogram
	LatAlter         prometheus.Histogram
	LatGet           prometheus.Histogram
	LatGetProvaAluno prometheus.Histogram
	LatGetAll        prometheus.Histogram
	LatDelete        prometheus.Histogram
	LatStatusService prometheus.Histogram
}

type instrumentationMiddleware struct {
	countMethods   *prometheus.CounterVec
	latencyMethods LatencyMethods
	next           prova.Service
}

func NewInstrumentation(cMethods *prometheus.CounterVec, lMethods LatencyMethods, nService prova.Service) prova.Service {
	return &instrumentationMiddleware{
		countMethods:   cMethods,
		latencyMethods: lMethods,
		next:           nService,
	}
}

func (im instrumentationMiddleware) Create(ctx context.Context, pro model.Prova) (output bool, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("create").Inc()
		im.latencyMethods.LatCreate.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Create(ctx, pro)
	return
}
func (im instrumentationMiddleware) Alter(ctx context.Context, pro model.Prova) (output bool, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("alter").Inc()
		im.latencyMethods.LatAlter.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Alter(ctx, pro)
	return
}
func (im instrumentationMiddleware) Get(ctx context.Context, id uint64) (output model.Prova, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("get").Inc()
		im.latencyMethods.LatGet.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.Get(ctx, id)
	return
}
func (im instrumentationMiddleware) GetProvaAluno(ctx context.Context, idProva uint64, raAluno string) (output model.Prova, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("getProvaAluno").Inc()
		im.latencyMethods.LatGetProvaAluno.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.GetProvaAluno(ctx, idProva, raAluno)
	return
}
func (im instrumentationMiddleware) GetAll(ctx context.Context, page uint32) (output []model.Prova, err error) {
	defer func(begin time.Time) {
		im.countMethods.WithLabelValues("total").Inc()
		im.countMethods.WithLabelValues("getAll").Inc()
		im.latencyMethods.LatGetAll.Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = im.next.GetAll(ctx, page)
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
