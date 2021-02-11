package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lcslucas/projeto-micro/services/prova/proto_prova"
	"github.com/lcslucas/projeto-micro/utils"
	"google.golang.org/grpc"
)

func ProvaStatusHandler(w http.ResponseWriter, r *http.Request) {
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

	strConn := fmt.Sprintf("%s:%s", os.Getenv("PRO_GRPC_HOST"), os.Getenv("PRO_GRPC_PORT"))

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(strConn, grpc.WithInsecure())
	if err != nil {
		log.Printf("Não foi possível conectar: %s", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ResponseHTTP{
			Status: false,
			Erro:   "Não foi possível se conectar com o serviço de provas, serviço indisponível.",
		})
		return
	}
	defer conn.Close()

	c := proto_prova.NewServiceProvaClient(conn)

	req := proto_prova.StatusServiceRequest{}

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

func ProvaGetProvaAlunoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	paramProvaID, err := strconv.ParseUint(vars["prova_id"], 10, 64)
	if err != nil {
		log.Printf("Erro: não foi possível pegar o parametro [prova_id] da requisição: %v", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ResponseHTTP{
			Status: false,
			Erro:   "Erro inesperado",
		})

		return
	}

	paramAlunoRA := vars["aluno_ra"]
	if paramAlunoRA == "" {
		log.Printf("Erro: não foi possível pegar o parametro [aluno_ra] da requisição: %v", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ResponseHTTP{
			Status: false,
			Erro:   "Erro inesperado",
		})

		return
	}

	err = godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Erro: não foi possível ler o .env: %v", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ResponseHTTP{
			Status: false,
			Erro:   "Erro inesperado",
		})

		return
	}

	strConn := fmt.Sprintf("%s:%s", os.Getenv("PRO_GRPC_HOST"), os.Getenv("PRO_GRPC_PORT"))

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

	c := proto_prova.NewServiceProvaClient(conn)

	req := proto_prova.GetProvaAlunoRequest{
		IdProva: paramProvaID,
		RaAluno: paramAlunoRA,
	}

	response, err := c.GetProvaAluno(ctx, &req)
	if err != nil {
		log.Printf("Não foi possível chamar o método GetProvaAluno: %s", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ResponseHTTP{
			Status: false,
			Erro:   "Não foi possível se conectar com o serviço de provas, serviço indisponível.",
		})
		return
	}

	if response.Error != "" {
		log.Printf("Erro recebido do servidor: %s", response.Error)

		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.ResponseHTTP{
			Status: false,
			Erro:   "Não foi possível se conectar com o serviço de provas, serviço indisponível.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.ResponseHTTP{
		Status: true,
		Dados:  response.Prova,
	})

}
