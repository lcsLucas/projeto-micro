package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/lcslucas/projeto-micro/services/exercicio/proto"
	"google.golang.org/grpc"
)

var (
	grpcHost = "localhost"
	grpcPort = "8082"
)

func TestStatusService(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceExercicioClient(conn)

	req := proto.StatusServiceRequest{}

	response, err := c.StatusService(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método StatusService: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}

	t.Log(response)
}

func TestCreate(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceExercicioClient(conn)

	req := proto.CreateAlterRequest{
		Exercicio: &proto.Exercicio{
			Id:        0,
			Nome:      "Exercicio de Teste",
			Descricao: "Descrição do exercicio de teste",
			Materia:   "Teste",
			Ativo:     false,
		},
	}

	response, err := c.Create(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método Create: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}

	t.Log(response)
}

func TestAlter(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceExercicioClient(conn)

	req := proto.CreateAlterRequest{
		Exercicio: &proto.Exercicio{
			Id:        0,
			Nome:      "Exercicio de Teste",
			Descricao: "Descrição do exercicio de teste",
			Materia:   "Teste",
			Ativo:     false,
		},
	}

	response, err := c.Alter(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método Alter: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}

	t.Log(response)
}

func TestGet(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceExercicioClient(conn)

	req := proto.GetRequest{
		Id: 1,
	}

	response, err := c.Get(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método Get: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}

	t.Log(response)

}

func TestGetSomes(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceExercicioClient(conn)

	req := proto.GetSomesRequest{
		Ids: []uint64{1, 3, 5},
	}

	response, err := c.GetSomes(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método GetSomes: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}

	t.Log(response)

}

func TestDelete(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceExercicioClient(conn)

	req := proto.DeleteRequest{
		Id: 1,
	}

	response, err := c.Delete(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método Delete: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}

	t.Log(response)

}
