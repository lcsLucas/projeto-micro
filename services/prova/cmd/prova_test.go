package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	proto "github.com/lcslucas/projeto-micro/services/prova/proto_prova"
	"github.com/lcslucas/projeto-micro/utils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	grpcHost = "localhost"
	grpcPort = "8083"
)

func TestStatusService(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceProvaClient(conn)

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

	c := proto.NewServiceProvaClient(conn)

	req := proto.CreateAlterRequest{
		Prova: &proto.Prova{
			Nome:         "Prova de Teste",
			DataCadastro: timestamppb.New(time.Now()),
			DataInicio:   timestamppb.New(time.Now()),
			DataFinal:    timestamppb.New(time.Now()),
			Serie:        "3º Ano Ensino Médio",
			Materia:      "Biologia",
			Bimestre:     1,
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

	c := proto.NewServiceProvaClient(conn)

	req := proto.CreateAlterRequest{
		Prova: &proto.Prova{
			Id:           1,
			Nome:         "Prova de Teste",
			DataCadastro: timestamppb.New(time.Now()),
			DataInicio:   timestamppb.New(time.Now()),
			DataFinal:    timestamppb.New(time.Now()),
			Serie:        "3º Ano Ensino Médio",
			Materia:      "Biologia",
			Bimestre:     2,
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

	c := proto.NewServiceProvaClient(conn)

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

func TestGetProvaAluno(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceProvaClient(conn)

	req := proto.GetProvaAlunoRequest{
		IdProva: 5,
		RaAluno: "35.329.394-5",
	}

	response, err := c.GetProvaAluno(context.Background(), &req)
	if err != nil {
		t.Fatalf("Não foi possível chamar o método Get: %s", err)
	}

	if response.Error != "" {
		t.Fatalf("Erro recebido do servidor: %s", response.Error)
	}

	utils.Pretty(response)

}

func TestGetAll(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceProvaClient(conn)

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

	t.Log(response)

}

func TestDelete(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcHost, grpcPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Não foi possível conectar: %s", err)
	}
	defer conn.Close()

	c := proto.NewServiceProvaClient(conn)

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
