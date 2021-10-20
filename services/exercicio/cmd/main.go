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
	"github.com/joho/godotenv"
	"github.com/lcslucas/projeto-micro/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"github.com/sony/gobreaker"
	"go.mongodb.org/mongo-driver/mongo"

	"google.golang.org/grpc"

	database "github.com/lcslucas/projeto-micro/database/mongodb"

	"github.com/lcslucas/projeto-micro/services/exercicio"
	"github.com/lcslucas/projeto-micro/services/exercicio/endpoints"
	"github.com/lcslucas/projeto-micro/services/exercicio/instrumentation"
	"github.com/lcslucas/projeto-micro/services/exercicio/logging"
	"github.com/lcslucas/projeto-micro/services/exercicio/migrations"
	proto "github.com/lcslucas/projeto-micro/services/exercicio/proto_exercicio"
	"github.com/lcslucas/projeto-micro/services/exercicio/proxying"
	"github.com/lcslucas/projeto-micro/services/exercicio/repository"
	"github.com/lcslucas/projeto-micro/services/exercicio/transport"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

var conn *mongo.Client
var configDB config.ConfigDB

var logger log.Logger

var cb *gobreaker.CircuitBreaker

var hostGrpcExe string
var portGrpcExe int
var portMetricExe int

func inicializeLogger() {

	file, err := os.OpenFile("/bin/projeto-service/temp/log_exercicio.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	logger = log.NewLogfmtLogger(file) //os.Stderr
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "exercicio",
		"caller", log.DefaultCaller,
		"hour", log.DefaultTimestampUTC,
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
	portMetricExe, _ = strconv.Atoi(os.Getenv("EXE_METRIC_PORT"))

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

func inicializeCircuitBreaker() {
	var st gobreaker.Settings
	st.Name = "Service Exercicio"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

func inicializeMetrics() (cMethods *prometheus.CounterVec, lMethods instrumentation.LatencyMethods) {

	cMethods = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "service_exercicio",
			Name:      "request_count",
			Help:      "Número de requisições recebidas no serviço Exercício",
		},
		[]string{"method"},
	)

	lMethods = instrumentation.LatencyMethods{
		LatCreate: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_exercicio_create",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Create do serviço Exercício",
			},
		),
		LatAlter: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_exercicio_alter",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Alter do serviço Exercício",
			},
		),
		LatGet: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_exercicio_get",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Get do serviço Exercício",
			},
		),
		LatGetSomes: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_exercicio_getSomes",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método GetSomes do serviço Exercício",
			},
		),
		LatDelete: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_exercicio_delete",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método Delete do serviço Exercício",
			},
		),
		LatStatusService: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "service_exercicio_status_service",
				Name:      "latency_seconds",
				Help:      "Durações de requisições recebidas no método StatusService do serviço Exercício",
			},
		),
	}

	prometheus.MustRegister(cMethods)

	prometheus.MustRegister(lMethods.LatCreate)
	prometheus.MustRegister(lMethods.LatAlter)
	prometheus.MustRegister(lMethods.LatGet)
	prometheus.MustRegister(lMethods.LatGetSomes)
	prometheus.MustRegister(lMethods.LatDelete)
	prometheus.MustRegister(lMethods.LatStatusService)

	return

}

func main() {
	var err error
	ctx := context.TODO()

	//* Inicializando Logger *//
	inicializeLogger()

	defer level.Info(logger).Log("msg", "service 'Exercicio' ended")
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

	//* Inicializando circuit breaker do serviço Aluno *//
	inicializeCircuitBreaker()
	//* Inicializando circuit breaker do serviço Aluno *//

	//* Inicializando métricas do serviço Exercicio *//
	countsService, latencysService := inicializeMetrics()
	//* Inicializando métricas do serviço Exercicio *//

	//* Definindo o serviço Exercício *//
	var service exercicio.Service
	{
		repository := repository.NewRepository(conn, configDB)
		service = exercicio.NewService(repository)
		service = logging.NewLogging(logger, service)
		service = proxying.NewProxying(cb, service)
		service = instrumentation.NewInstrumentation(countsService, latencysService, service)
	}
	//* Definindo o serviço Exercício *//

	//* Inicializando Conexão gRPC do serviço Exercicio *//
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

	//* Inicializando Conexão http do serviço Exercicio *//
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
			Addr:    fmt.Sprintf("%s:%d", hostGrpcExe, portMetricExe),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		}

		r.Path("/metrics").Handler(promhttp.Handler())
		logger.Log("error", srv.ListenAndServe())

	}()
	//* Inicializando Conexão http do serviço Exercicio *//

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
