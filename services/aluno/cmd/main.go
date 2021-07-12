package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/lcslucas/projeto-micro/services/aluno/instrumentation"
	proto "github.com/lcslucas/projeto-micro/services/aluno/proto_aluno"
	"github.com/lcslucas/projeto-micro/services/aluno/repository"
	"github.com/rs/cors"

	"github.com/lcslucas/projeto-micro/config"
	"github.com/lcslucas/projeto-micro/services/aluno"
	"github.com/lcslucas/projeto-micro/services/aluno/endpoints"
	"github.com/lcslucas/projeto-micro/services/aluno/transport"
	"google.golang.org/grpc"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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
var portMetricAlu int

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
	portMetricAlu, _ = strconv.Atoi(os.Getenv("ALU_METRIC_PORT"))

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

func inicializeMetrics() (cMethods *prometheus.CounterVec, lMethods instrumentation.LatencyMethods) {

	cMethods = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "service_aluno",
			Name:      "request_count",
			Help:      "Número de requisições recebidas no serviço Aluno",
		},
		[]string{"method"},
	)

	lMethods = instrumentation.LatencyMethods{
		LatCreate: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_aluno_create",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Create do serviço Aluno",
			},
		),
		LatAlter: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_aluno_alter",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Alter do serviço Aluno",
			},
		),
		LatGet: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_aluno_get",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Get do serviço Aluno",
			},
		),
		LatGetAll: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_aluno_getAll",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método GetAll do serviço Aluno",
			},
		),
		LatDelete: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_aluno_delete",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Delete do serviço Aluno",
			},
		),
		LatStatusService: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_aluno_status_service",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método StatusService do serviço Aluno",
			},
		),
	}

	prometheus.MustRegister(cMethods)

	prometheus.MustRegister(lMethods.LatCreate)
	prometheus.MustRegister(lMethods.LatAlter)
	prometheus.MustRegister(lMethods.LatGet)
	prometheus.MustRegister(lMethods.LatGetAll)
	prometheus.MustRegister(lMethods.LatDelete)
	prometheus.MustRegister(lMethods.LatStatusService)

	return
}

func main() {
	var err error
	ctx := context.TODO()

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

	//* Inicializando métricas do serviço Aluno *//
	countsService, latencysService := inicializeMetrics()
	//* Inicializando métricas do serviço Aluno *//

	//* Definindo o serviço Aluno *//
	var service aluno.Service
	{
		repository := repository.NewRepository(conn, logger, configDB)
		service = aluno.NewService(repository, logger)
		service = instrumentation.NewInstrumentation(countsService, latencysService, service)
	}
	//* Definindo o serviço Aluno *//

	//* Inicializando Conexão gRPC do serviço Aluno *//
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

	//* Inicializando Conexão http do serviço Aluno *//
	go func() {
		/*
			http.Handle("/metrics", promhttp.Handler())
			logger.Log("err", http.ListenAndServe(":9999", nil))
		*/

		r := mux.NewRouter().StrictSlash(true)

		handler := cors.Default().Handler(r)

		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"},
			AllowedHeaders:   []string{"Access-Control-Allow-Credentials", "Access-Control-Allow-Origin", "Authorization", "Content-Type"},
			AllowCredentials: true,
			Debug:            false,
		})

		handler = c.Handler(handler)

		srv := &http.Server{
			Handler: handler,
			Addr:    fmt.Sprintf("%s:%d", hostGrpcAlu, portMetricAlu),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		}

		r.Path("/metrics").Handler(promhttp.Handler())
		logger.Log("error", srv.ListenAndServe())

	}()
	//* Inicializando Conexão http do serviço Aluno *//

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
