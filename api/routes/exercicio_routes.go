package routes

import (
	"net/http"

	"github.com/lcslucas/projeto-micro/api/handlers"
	"github.com/lcslucas/projeto-micro/api/middlewares"
)

var exercicioRoutes = []Route{
	{
		Uri:     "/exercicios/status",
		Method:  http.MethodGet,
		Handler: middlewares.SetMiddlewareJSON(handlers.ExercicioStatusHandler),
	},
}
