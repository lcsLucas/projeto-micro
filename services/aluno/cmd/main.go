package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	proto "github.com/lcslucas/projeto-micro/services/aluno/proto_aluno"
	"github.com/lcslucas/projeto-micro/services/aluno/repository"

	"github.com/lcslucas/projeto-micro/config"
	"github.com/lcslucas/projeto-micro/services/aluno"
	"github.com/lcslucas/projeto-micro/services/aluno/endpoints"
	"github.com/lcslucas/projeto-micro/services/aluno/transport"
	"google.golang.org/grpc"

	"github.com/lcslucas/projeto-micro/services/aluno/migrations"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	database "github.com/lcslucas/projeto-micro/database/mongodb"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
)

var conn *mongo.Client
var configDB config.ConfigDB

var logger log.Logger

var hostGrpcAlu string
var portGrpcAlu int

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
	err := godotenv.Load("/bin/projeto-service/.env")
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

	hostGrpcAlu = os.Getenv("ALU_GRPC_HOST")
	portGrpcAlu, _ = strconv.Atoi(os.Getenv("ALU_GRPC_PORT"))

	return nil
}

func inicializeDB(ctx context.Context) error {
	newConn, err := database.Connect(ctx, configDB)
	if err != nil {
		return err
	}

	conn = newConn
	return migrations.ExecMigrationAlunos(ctx, configDB.DBName, conn)
}

func main() {
	var err error
	ctx := context.Background()

	//* Inicializando Logger *//
	inicializeLogger()

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
	}
	defer conn.Disconnect(ctx)
	//* Inicializando Conexão com o banco de dados *//

	//* Inicializando Conexão gRPC do serviço Aluno *//
	var service aluno.Service
	{
		repository := repository.NewRepository(conn, logger, configDB)
		service = aluno.NewService(repository, logger)
	}

	var (
		eps        = endpoints.NewEndpointSet(service)
		grpcServer = transport.NewGrpcServer(eps)
	)

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostGrpcAlu, portGrpcAlu))
	if err != nil {
		logger.Log("transport", "gRPC", "during", "listen", "err", err)
		os.Exit(1)
	}
	defer grpcListener.Close()

	level.Info(logger).Log("msg", fmt.Sprintf("serviço 'Aluno' inicializado em: %s:%d", hostGrpcAlu, portGrpcAlu))

	go func() {
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
		proto.RegisterServiceAlunoServer(baseServer, grpcServer)
		err = baseServer.Serve(grpcListener)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "listen", "err", "Fatal to serve:", hostGrpcAlu, portGrpcAlu, ":", err)
			os.Exit(1)
		}
	}()
	//* Inicializando Conexão gRPC do serviço Aluno *//

	//* Notifica o programa quando for encerrado *//
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Error(logger).Log("exit", <-errs)
}
