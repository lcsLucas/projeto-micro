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

	"github.com/joho/godotenv"
	"github.com/lcslucas/projeto-micro/config"
	"go.mongodb.org/mongo-driver/mongo"

	"google.golang.org/grpc"

	database "github.com/lcslucas/projeto-micro/database/mongodb"

	"github.com/lcslucas/projeto-micro/services/exercicio"
	"github.com/lcslucas/projeto-micro/services/exercicio/endpoints"
	"github.com/lcslucas/projeto-micro/services/exercicio/migrations"
	proto "github.com/lcslucas/projeto-micro/services/exercicio/proto_exercicio"
	"github.com/lcslucas/projeto-micro/services/exercicio/repository"
	"github.com/lcslucas/projeto-micro/services/exercicio/transport"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

var conn *mongo.Client
var configDB config.ConfigDB

var logger log.Logger

var hostGrpcExe string
var portGrpcExe int

func inicializeLogger() {
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "exercicio",
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
		Driver:   os.Getenv("EXE_DB_DRIVER"),
		Host:     os.Getenv("EXE_DB_HOST"),
		User:     os.Getenv("EXE_DB_USER"),
		Password: os.Getenv("EXE_DB_PASSWORD"),
		Port:     os.Getenv("EXE_DB_PORT"),
		DBName:   os.Getenv("EXE_DB_NAME"),
	}

	hostGrpcExe = os.Getenv("EXE_GRPC_HOST")
	portGrpcExe, _ = strconv.Atoi(os.Getenv("EXE_GRPC_PORT"))

	return nil
}

func inicializeDB(ctx context.Context) error {
	newConn, err := database.Connect(ctx, configDB)
	if err != nil {
		return err
	}

	conn = newConn
	return migrations.ExecMigrationExercicios(ctx, configDB.DBName, conn)
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

	//* Inicializando Conexão gRPC do serviço Exercicio *//
	var service exercicio.Service
	{
		repository := repository.NewRepository(conn, logger, configDB)
		service = exercicio.NewService(repository, logger)
	}

	var (
		eps        = endpoints.NewEndpointSet(service)
		grpcServer = transport.NewGrpcServer(eps)
	)

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostGrpcExe, portGrpcExe))
	if err != nil {
		logger.Log("transport", "gRPC", "during", "listen", "err", err)
		os.Exit(1)
	}
	defer grpcListener.Close()

	level.Info(logger).Log("msg", fmt.Sprintf("serviço 'Exercicio' inicializado em: %s:%d", hostGrpcExe, portGrpcExe))

	go func() {
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
		proto.RegisterServiceExercicioServer(baseServer, grpcServer)
		err = baseServer.Serve(grpcListener)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "listen", "err", "Fatal to serve:", hostGrpcExe, portGrpcExe, ":", err)
			os.Exit(1)
		}
	}()
	//* Inicializando Conexão gRPC do serviço Exercicio *//

	//* Notifica o programa quando for encerrado *//
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Error(logger).Log("exit", <-errs)
	//* Notifica o programa quando for encerrado *//

}
