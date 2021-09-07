package main

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	ot "github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-track-api/internal/api"
	kafka "github.com/ozonva/ova-track-api/internal/kafka_client"
	"github.com/ozonva/ova-track-api/internal/repo"
	track_server "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	jgrc "github.com/uber/jaeger-client-go/config"
	jgrl "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"google.golang.org/grpc"
	"io"
	"net/http"

	"log"
	"net"
	"strconv"
)


func main() {

	log.Println("Hi, i am ova-track-api!")
	go runPrometheus()

	tracer, closer := initTracer()
	ot.SetGlobalTracer(tracer)
	defer closer.Close()

	dsn :=  "postgres://admin:admin@localhost:5434/db?sslmode=disable"

	pdb, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Println("failed to load driver", err)
	}

	err = pdb.Ping()
	if err != nil {
		log.Fatalln("failed to connect to db:", err)
	}

	port := ":"+strconv.Itoa(8080)
	listen, err := net.Listen("tcp", port)
	log.Println("TCP is starting up at port ", port)

	if err != nil {
		log.Fatalln("Failed to listen server  ", err)
	}

	log.Println("TCP started successfully", port)

	grpcService := grpc.NewServer()
	rp :=repo.NewSQLTrackRepo(pdb)
	kafkaClient := kafka.NewKafkaClient()
	kafkaConn := "kafka:9092"
	if kafkaConnErr := kafkaClient.Connect(context.Background(), kafkaConn, "CUDEvents",0); kafkaConnErr != nil{
		log.Fatalf("Can not connect to kafka, %s", kafkaConnErr)
	}

	track_server.RegisterTrackServer(grpcService, api.NewApiServer(rp,api.NewApiMetrics(),kafkaClient))

	log.Println("Starting track service")
	if err := grpcService.Serve(listen); err != nil {
		log.Fatalln("failed to serve", err)
	}
}

func runPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalln("failed to start listen to metric requests, error", err)
	}
}


func initTracer() (ot.Tracer, io.Closer) {
	cfg := jgrc.Configuration{
		ServiceName: "ova-recipe-api",
		Sampler: &jgrc.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jgrc.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jgrl.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jgrc.Logger(jLogger),
		jgrc.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Fatalln("Can not create tracer", err)
	}
	return tracer, closer
}
