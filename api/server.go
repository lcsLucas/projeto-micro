package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/log/level"
	"github.com/lcslucas/projeto-micro/api/router"
	"github.com/lcslucas/projeto-micro/api/routes"

	"github.com/rs/cors"
)

func Run(addr string) {
	r := router.NewRouter()

	handler := cors.Default().Handler(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Credentials", "Access-Control-Allow-Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            false,
	})

	handler = c.Handler(handler)

	srv := &http.Server{
		Handler: handler,
		Addr:    addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	routes.InicializeLogger()

	level.Info(routes.Logger).Log("msg", fmt.Sprintf("Servidor escutando em: %s", addr))
	level.Error(routes.Logger).Log("exit", srv.ListenAndServe())

}
