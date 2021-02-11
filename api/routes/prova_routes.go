package routes

import (
	"net/http"

	"github.com/lcslucas/projeto-micro/api/handlers"
	"github.com/lcslucas/projeto-micro/api/middlewares"
)

var provaRoutes = []Route{
	{
		Uri:     "/provas/status",
		Method:  http.MethodGet,
		Handler: middlewares.SetMiddlewareJSON(handlers.ProvaStatusHandler),
	},
	{
		Uri:     "/provas/get/{prova_id}/aluno/{aluno_ra}",
		Method:  http.MethodGet,
		Handler: middlewares.SetMiddlewareJSON(handlers.ProvaGetProvaAlunoHandler),
	},
}
