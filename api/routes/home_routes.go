package routes

import (
	"net/http"

	"github.com/lcslucas/projeto-micro/api/handlers"
)

var homeRoutes = []Route{
	{
		Uri:     "/",
		Method:  http.MethodGet,
		Handler: handlers.HomeHandler,
	},
}
