package routes

import (
	"net/http"

	"github.com/lcslucas/projeto-micro/api/handlers"
	"github.com/lcslucas/projeto-micro/api/middlewares"
)

var alunoRoutes = []Route{
	{
		Uri:     "/alunos/status",
		Method:  http.MethodGet,
		Handler: middlewares.SetMiddlewareJSON(handlers.AlunoStatusHandler),
	},
}
