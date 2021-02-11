package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
		log.Println("> Rota acessada...")
		log.Println("Host:", r.Host, "URI:", r.RequestURI, "Method:", r.Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}

	r.Use(loggingMiddleware)

	return r
}
