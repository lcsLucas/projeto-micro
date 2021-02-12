package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lcslucas/projeto-micro/api"
)

func main() {
	err := godotenv.Load("/bin/projeto-service/.env")
	if err != nil {
		log.Fatalf("Erro ao ler o .env: %v", err)
	}

	urlAPI := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	api.Run(urlAPI)

}
