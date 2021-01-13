package main

import (
	"context"
	"flag"
	"os"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	database "github.com/lcslucas/micro-service/database/mongodb"

	"github.com/joho/godotenv"
	"github.com/lcslucas/projeto-micro/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var conn *mongo.Client
var configDB config.ConfigDB

var logger log.Logger

var host_grpc_alu string
var port_grpc_alu int

func inicializeLogger() {
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "aluno",
		"hour", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
}

func inicializeVars() error {
	err := godotenv.Load("../../../.env")
	if err != nil {
		return err
	}

	configDB = config.ConfigDB{
		Driver:   os.Getenv("ALU_DB_DRIVER"),
		Host:     os.Getenv("ALU_DB_HOST"),
		User:     os.Getenv("ALU_DB_USER"),
		Password: os.Getenv("ALU_DB_PASSWORD"),
		Port:     os.Getenv("ALU_DB_PORT"),
		DBName:   os.Getenv("ALU_DB_NAME"),
	}

	host_grpc_alu = os.Getenv("ALU_GRPC_HOST")
	port_grpc_alu, _ = strconv.Atoi(os.Getenv("ALU_GRPC_PORT"))

	return nil
}

func inicializeDB(ctx context.Context) error {
	newConn, err := database.Connect(ctx, configDB)
	if err != nil {
		return err
	}

	conn = newConn
	return nil
}

func main() {
	var err error
	ctx := context.Background()

	//* Inicializando Logger *//
	inicializeLogger()

	level.Info(logger).Log("msg", "service 'Aluno' started")
	defer level.Info(logger).Log("msg", "service 'Aluno' ended")
	flag.Parse()
	//* Inicializando Logger *//

	//* Inicializando variáveis *//
	err = inicializeVars()
	if err != nil {
		level.Error(logger).Log("exit", err)
	}
	//* Inicializando variáveis *//

	//* Inicializando Conexão com o banco de dados *//
	err = inicializeDB(ctx)
	if err != nil {
		level.Error(logger).Log("exit", err)
		defer conn.Disconnect(ctx)
	}
	//* Inicializando Conexão com o banco de dados *//

}
