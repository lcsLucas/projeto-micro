package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/lcslucas/projeto-micro/services/exercicio"
	"github.com/lcslucas/projeto-micro/services/exercicio/model"
)

type loggingMidleware struct {
	logger log.Logger
	next   exercicio.Service
}

func NewLogging(logger log.Logger, nService exercicio.Service) exercicio.Service {
	return &loggingMidleware{
		logger: logger,
		next:   nService,
	}
}

func (lm loggingMidleware) Create(ctx context.Context, exe model.Exercicio) (output bool, err error) {
	defer func(begin time.Time) {
		logger := log.With(lm.logger, "method", "Create")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", exe),
			//"output", output,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
		)
	}(time.Now())

	output, err = lm.next.Create(ctx, exe)
	return
}

func (lm loggingMidleware) Alter(ctx context.Context, exe model.Exercicio) (output bool, err error) {
	defer func(begin time.Time) {
		logger := log.With(lm.logger, "method", "Alter")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", exe),
			//"output", output,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
		)
	}(time.Now())

	output, err = lm.next.Alter(ctx, exe)
	return
}

func (lm loggingMidleware) Get(ctx context.Context, id uint64) (output model.Exercicio, err error) {
	defer func(begin time.Time) {
		//str_output, _ := json.Marshal(output)

		logger := log.With(lm.logger, "method", "Get")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", id),
			//"output", string(str_output),
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
		)
	}(time.Now())

	output, err = lm.next.Get(ctx, id)
	return
}

func (lm loggingMidleware) GetSomes(ctx context.Context, ids []uint64) (output []model.Exercicio, err error) {
	defer func(begin time.Time) {
		//str_output, _ := json.Marshal(output)

		logger := log.With(lm.logger, "method", "GetAll")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", ids),
			//"output", string(str_output),
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
		)
	}(time.Now())

	output, err = lm.next.GetSomes(ctx, ids)
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
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
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
			"duration", fmt.Sprintf("%vms", time.Since(begin).Milliseconds()),
		)
	}(time.Now())

	output, err = lm.next.StatusService(ctx)
	return
}
