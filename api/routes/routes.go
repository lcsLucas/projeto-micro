package routes

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Logger log.Logger

func InicializeLogger() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.NewSyncLogger(Logger)
	Logger = log.With(Logger,
		"service", "api",
		"hour", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
}

type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

/*Load carrega todas as rotas (Home, Pedidos)*/
func Load() []Route {
	var allRoutes []Route

	allRoutes = append(allRoutes, homeRoutes...)      // rotas de home
	allRoutes = append(allRoutes, alunoRoutes...)     // rotas de aluno
	allRoutes = append(allRoutes, exercicioRoutes...) // rotas de exercicio
	allRoutes = append(allRoutes, provaRoutes...)     // rotas de prova

	return allRoutes
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level.Info(Logger).Log("msg", "Requisitando...", "host", r.Host, "uri", r.RequestURI, "method", r.Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}

	r.Use(loggingMiddleware)

	r.Path("/metrics").Handler(promhttp.Handler())

	return r
}
