package middlewares

import (
	"net/http"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func Metrics(next http.HandlerFunc) http.HandlerFunc {
	var dur metrics.Histogram = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "myservice",
		Subsystem: "api",
		Name:      "request_duration_seconds",
		Help:      "Total time spent serving requests.",
	}, []string{})

	defer func(begin time.Time) {
		dur.Observe(time.Since(begin).Seconds())
	}(time.Now())

	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
