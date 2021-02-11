package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/lcslucas/projeto-micro/services/exercicio/proto_exercicio"
	"github.com/lcslucas/projeto-micro/utils"
	"google.golang.org/grpc"
)

func ExercicioStatusHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Erro: não foi possível ler o .env: %v", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ResponseHTTP{
			Status: false,
			Erro:   "Erro inesperado",
		})

		return
	}

	strConn := fmt.Sprintf("%s:%s", os.Getenv("EXE_GRPC_HOST"), os.Getenv("EXE_GRPC_PORT"))

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(strConn, grpc.WithInsecure())
	if err != nil {
		log.Printf("Não foi possível conectar: %s", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ResponseHTTP{
			Status: false,
			Erro:   "Não foi possível se conectar com o serviço de exercícios, serviço indisponível.",
		})
		return
	}
	defer conn.Close()

	c := proto_exercicio.NewServiceExercicioClient(conn)

	req := proto_exercicio.StatusServiceRequest{}

	response, err := c.StatusService(ctx, &req)
	if err != nil {
		log.Printf("Não foi possível chamar o método StatusService: %s", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ResponseHTTP{
			Status: false,
			Erro:   "Não foi possível se conectar com o serviço de alunos, serviço indisponível.",
		})
		return
	}

	if response.Error != "" {
		log.Printf("Erro recebido do servidor: %s", response.Error)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ResponseHTTP{
			Status: false,
			Erro:   "Não foi possível se conectar com o serviço de alunos, serviço indisponível.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.ResponseHTTP{
		Status: true,
	})

}
