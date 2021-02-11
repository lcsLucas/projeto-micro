package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lcslucas/projeto-micro/api/router"
)

func Run(addr string) {
	r := router.NewRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	fmt.Printf("Servidor escutando em: %s\n", addr)
	log.Fatal(srv.ListenAndServe())

}
