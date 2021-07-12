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

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"

	"google.golang.org/grpc"

	database "github.com/lcslucas/projeto-micro/database/postgresql"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/lcslucas/projeto-micro/config"
	"github.com/lcslucas/projeto-micro/services/prova"
	"github.com/lcslucas/projeto-micro/services/prova/endpoints"
	"github.com/lcslucas/projeto-micro/services/prova/instrumentation"
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
var portMetricProva int

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
	portMetricProva, _ = strconv.Atoi(os.Getenv("PRO_METRIC_PORT"))

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

func inicializeMetrics() (cMethods *prometheus.CounterVec, lMethods instrumentation.LatencyMethods) {

	cMethods = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "service_prova",
			Name:      "request_count",
			Help:      "Número de requisições recebidas no serviço Prova",
		},
		[]string{"method"},
	)

	lMethods = instrumentation.LatencyMethods{
		LatCreate: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_prova_create",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Create do serviço Prova",
			},
		),
		LatAlter: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_prova_alter",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Alter do serviço Prova",
			},
		),
		LatGet: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_prova_get",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Get do serviço Prova",
			},
		),
		LatGetProvaAluno: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_prova_getProvaAluno",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método GetProvaAluno do serviço Prova",
			},
		),
		LatGetAll: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_prova_getAll",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método GetAll do serviço Prova",
			},
		),
		LatDelete: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_prova_delete",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Delete do serviço Prova",
			},
		),
		LatStatusService: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_prova_status_service",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método StatusService do serviço Prova",
			},
		),
	}

	prometheus.MustRegister(cMethods)

	prometheus.MustRegister(lMethods.LatCreate)
	prometheus.MustRegister(lMethods.LatAlter)
	prometheus.MustRegister(lMethods.LatGet)
	prometheus.MustRegister(lMethods.LatGetProvaAluno)
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

	defer level.Info(logger).Log("msg", "service 'Prova' ended")
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
		return
	}

	sqlDB, _ := conn.DB()
	defer sqlDB.Close()
	//* Inicializando Conexão com o banco de dados *//

	//* Inicializando métricas do serviço Prova *//
	countsService, latencysService := inicializeMetrics()
	//* Inicializando métricas do serviço Prova *//

	//* Definindo o serviço Prova *//
	var service prova.Service
	{
		repository := repository.NewRepository(conn, logger, configDB)
		service = prova.NewService(repository, logger)
		service = instrumentation.NewInstrumentation(countsService, latencysService, service)
	}

	//* Inicializando Conexão gRPC do serviço Prova *//
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

	//* Inicializando Conexão http do serviço Prova *//
	go func() {

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
			Addr:    fmt.Sprintf("%s:%d", hostGrpcProva, portMetricProva),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		}

		r.Path("/metrics").Handler(promhttp.Handler())
		logger.Log("error", srv.ListenAndServe())

	}()
	//* Inicializando Conexão http do serviço Prova *//

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
