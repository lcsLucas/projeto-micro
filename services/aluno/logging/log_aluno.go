package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/lcslucas/projeto-micro/services/aluno"
	"github.com/lcslucas/projeto-micro/services/aluno/model"
)

type loggingMidleware struct {
	logger log.Logger
	next   aluno.Service
}

func NewLogging(logger log.Logger, nService aluno.Service) aluno.Service {
	return &loggingMidleware{
		logger: logger,
		next:   nService,
	}
}

func (lm loggingMidleware) Create(ctx context.Context, alu model.Aluno) (output bool, err error) {
	defer func(begin time.Time) {
		logger := log.With(lm.logger, "method", "Create")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", alu),
			"output", output,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.Create(ctx, alu)
	return
}

func (lm loggingMidleware) Alter(ctx context.Context, alu model.Aluno) (output bool, err error) {
	defer func(begin time.Time) {
		logger := log.With(lm.logger, "method", "Alter")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", alu),
			"output", output,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.Alter(ctx, alu)
	return
}

func (lm loggingMidleware) Get(ctx context.Context, ra string) (output model.Aluno, err error) {
	defer func(begin time.Time) {
		str_output, _ := json.Marshal(output)

		logger := log.With(lm.logger, "method", "Get")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", ra),
			"output", string(str_output),
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.Get(ctx, ra)
	return
}

func (lm loggingMidleware) GetAll(ctx context.Context, page uint32) (output []model.Aluno, err error) {
	defer func(begin time.Time) {
		str_output, _ := json.Marshal(output)

		logger := log.With(lm.logger, "method", "GetAll")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", page),
			"output", string(str_output),
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.GetAll(ctx, page)
	return
}

func (lm loggingMidleware) Delete(ctx context.Context, ra string) (output bool, err error) {
	defer func(begin time.Time) {
		logger := log.With(lm.logger, "method", "Delete")
		level.Info(logger).Log(
			"parameter", fmt.Sprintf("%v", ra),
			"output", output,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.Delete(ctx, ra)
	return
}

func (lm loggingMidleware) StatusService(ctx context.Context) (output bool, err error) {
	defer func(begin time.Time) {
		logger := log.With(lm.logger, "method", "StatusService")
		level.Info(logger).Log(
			"output", output,
			"error", err,
			"ended", time.Now(),
			"duration", fmt.Sprintf("%vs", time.Since(begin).Seconds()),
		)
	}(time.Now())

	output, err = lm.next.StatusService(ctx)
	return
}
