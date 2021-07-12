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

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"

	"google.golang.org/grpc"

	database "github.com/lcslucas/projeto-micro/database/postgresql"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/lcslucas/projeto-micro/config"
	"github.com/lcslucas/projeto-micro/services/prova"
	"github.com/lcslucas/projeto-micro/services/prova/endpoints"
	"github.com/lcslucas/projeto-micro/services/prova/migrations"
	proto "github.com/lcslucas/projeto-micro/services/prova/proto_prova"
	"github.com/lcslucas/projeto-micro/services/prova/repository"
	"github.com/lcslucas/projeto-micro/services/prova/transport"
	"gorm.io/gorm"
)

var conn *gorm.DB
var configDB config.ConfigDB

var logger log.Logger

var hostGrpcProva string
var portGrpcProva int

func inicializeLogger() {
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "prova",
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
		Driver:   os.Getenv("PRO_DB_DRIVER"),
		Host:     os.Getenv("PRO_DB_HOST"),
		User:     os.Getenv("PRO_DB_USER"),
		Password: os.Getenv("PRO_DB_PASSWORD"),
		Port:     os.Getenv("PRO_DB_PORT"),
		DBName:   os.Getenv("PRO_DB_NAME"),
	}

	hostGrpcProva = os.Getenv("PRO_GRPC_HOST")
	portGrpcProva, _ = strconv.Atoi(os.Getenv("PRO_GRPC_PORT"))

	return nil
}

func inicializeDB(ctx context.Context) error {
	newConn, err := database.Connect(configDB, false)
	if err != nil {
		return err
	}

	fmt.Println("eae entrou aqui!?")

	err = migrations.ExecCreateDatabaseProva(ctx, configDB.DBName, newConn)
	if err != nil {
		return err
	}
	sqlDB, _ := newConn.DB()
	sqlDB.Close()

	conn, err = database.Connect(configDB, true)
	if err != nil {
		return err
	}

	return migrations.ExecMigrationProva(ctx, configDB.DBName, conn)
}

func main() {
	var err error
	ctx := context.Background()

	//* Inicializando Logger *//
	inicializeLogger()

	defer level.Info(logger).Log("msg", "service 'Prova' ended")
	flag.Parse()
	//* Inicializando Logger *//

	//* Inicializando variáveis *//
	err = inicializeVars()
	if err != nil {
		level.Error(logger).Log("exit", err)
	}
	//* Inicializando variáveis *//

	//* Iniciando conexão com o banco*//
	err = inicializeDB(ctx)
	if err != nil {
		level.Error(logger).Log("exit", err)
		return
	}

	sqlDB, _ := conn.DB()
	defer sqlDB.Close()
	//*Iniciando conexão com o banco*//

	//* Inicializando Conexão gRPC do serviço Prova *//
	var service prova.Service
	{
		repository := repository.NewRepository(conn, logger, configDB)
		service = prova.NewService(repository, logger)
	}

	var (
		eps        = endpoints.NewEndpointSet(service)
		grpcServer = transport.NewGrpcServer(eps)
	)

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostGrpcProva, portGrpcProva))
	if err != nil {
		logger.Log("transport", "gRPC", "during", "listen", "err", err)
		os.Exit(1)
	}
	defer grpcListener.Close()

	level.Info(logger).Log("msg", fmt.Sprintf("serviço 'Prova' inicializado em: %s:%d", hostGrpcProva, portGrpcProva))

	go func() {
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
		proto.RegisterServiceProvaServer(baseServer, grpcServer)
		err = baseServer.Serve(grpcListener)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "listen", "err", "Fatal to serve:", hostGrpcProva, portGrpcProva, ":", err)
			os.Exit(1)
		}
	}()

	//* Inicializando Conexão gRPC do serviço Prova *//

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
