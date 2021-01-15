package main

import (
	"context"
	"testing"

	"github.com/lcslucas/projeto-micro/services/aluno/proto"
	"google.golang.org/grpc"
)

func TestServiceStatusService(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceAlunoClient(conn)

	req := proto.StatusServiceRequest{}

	response, err := c.StatusService(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método StatusService: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}
}

func TestServiceGet(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceAlunoClient(conn)

	req := proto.GetRequest{
		Ra: "10.846.2074-5",
	}

	response, err := c.Get(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método Get: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}

	t.Log(response.Aluno)
}
func TestServiceGetAll(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceAlunoClient(conn)

	req := proto.GetAllRequest{
		Page: 1,
	}

	response, err := c.GetAll(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método GetAll: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}
}
func TestServiceCreate(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceAlunoClient(conn)

	req := proto.CreateAlterRequest{
		Aluno: &proto.Aluno{
			Ra:      "111",
			Nome:    "Teste de Aluno",
			Email:   "teste@email.com",
			Celular: "(18) 9999-9999",
		},
	}

	response, err := c.Create(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método Create: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}
}

func TestServiceAlter(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceAlunoClient(conn)

	req := proto.CreateAlterRequest{
		Aluno: &proto.Aluno{
			Ra:      "111",
			Nome:    "Teste de Aluno",
			Email:   "teste@email.com",
			Celular: "(18) 9999-9999",
		},
	}

	response, err := c.Alter(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método Alter: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}
}

func TestServiceDelete(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceAlunoClient(conn)

	req := proto.DeleteRequest{
		Ra: "40.738.017-6",
	}

	response, err := c.Delete(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método Delete: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}
}
