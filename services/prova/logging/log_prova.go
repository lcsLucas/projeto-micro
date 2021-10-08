package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/lcslucas/projeto-micro/services/prova"
	"github.com/lcslucas/projeto-micro/services/prova/model"
)

type loggingMidleware struct {
	logger log.Logger
	next   prova.Service
}

func NewLogging(logger log.Logger, nService prova.Service) prova.Service {
	return &loggingMidleware{
		logger: logger,
		next:   nService,
	}
}

func (lm loggingMidleware) Create(ctx context.Context, prova model.Prova) (output bool, err error) {
	defer func(begin time.Time) {
		logger := log.With(lm.logger, "method", "Create")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", prova),
			//"output", output,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.Create(ctx, prova)
	return
}

func (lm loggingMidleware) Alter(ctx context.Context, prova model.Prova) (output bool, err error) {
	defer func(begin time.Time) {
		logger := log.With(lm.logger, "method", "Alter")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", prova),
			//"output", output,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.Alter(ctx, prova)
	return
}

func (lm loggingMidleware) Get(ctx context.Context, id uint64) (output model.Prova, err error) {
	defer func(begin time.Time) {
		//str_output, _ := json.Marshal(output)

		logger := log.With(lm.logger, "method", "Get")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", id),
			//"output", string(str_output),
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.Get(ctx, id)
	return
}

func (lm loggingMidleware) GetProvaAluno(ctx context.Context, idProva uint64, raAluno string) (output model.Prova, err error) {
	defer func(begin time.Time) {
		//str_output, _ := json.Marshal(output)

		logger := log.With(lm.logger, "method", "GetAll")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v, %v", idProva, raAluno),
			//"output", string(str_output),
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.GetProvaAluno(ctx, idProva, raAluno)
	return
}

func (lm loggingMidleware) GetAll(ctx context.Context, page uint32) (output []model.Prova, err error) {
	defer func(begin time.Time) {
		//str_output, _ := json.Marshal(output)

		logger := log.With(lm.logger, "method", "GetAll")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", page),
			//"output", string(str_output),
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.GetAll(ctx, page)
	return
}

func (lm loggingMidleware) Delete(ctx context.Context, id uint64) (output bool, err error) {
	defer func(begin time.Time) {
		logger := log.With(lm.logger, "method", "Delete")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", id),
			//"output", output,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.Delete(ctx, id)
	return
}

func (lm loggingMidleware) StatusService(ctx context.Context) (output bool, err error) {
	defer func(begin time.Time) {
		logger := log.With(lm.logger, "method", "StatusService")
		level.Info(logger).Log(
			//"output", output,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.StatusService(ctx)
	return
}
