package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	ot "github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-track-api/internal/api"
	"github.com/ozonva/ova-track-api/internal/repo"
	track_server "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	kafka "github.com/ozonva/ova-track-api/internal/kafka_client"
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

	fmt.Println("Hi, i am ova-track-api!")
	go runPrometheus()

	tracer, closer := initTracer()
	ot.SetGlobalTracer(tracer)
	defer closer.Close()

	dsn :=  "postgres://admin:admin@localhost:5434/db?sslmode=disable"

	pdb, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to load driver: %v", err)
	}

	err = pdb.Ping()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	port := ":"+strconv.Itoa(8080)
	listen, err := net.Listen("tcp", port)
	log.Printf("TCP at %v is starting up ...", port)

	if err != nil {
		log.Printf("Failed to listen server %v", err)
	}

	log.Printf("TCP at %v started successfully", port)

	grpcService := grpc.NewServer()
	rp :=repo.NewSQLTrackRepo(pdb)
	fmt.Println("Hi, i am ova-recipe-api!")
	kafkaClient := kafka.NewKafkaClient()
	kafkaConn := "kafka:9092"
	if kafkaConnErr := kafkaClient.Connect(context.Background(), kafkaConn, "CUDEvents",0); kafkaConnErr != nil{
		log.Fatalf("Can not connect to kafka, %s", kafkaConnErr)
	}

	track_server.RegisterTrackServer(grpcService, api.NewApiServer(rp,api.NewApiMetrics(),kafkaClient))

	log.Print("Starting track service")
	if err := grpcService.Serve(listen); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func runPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start listen to metric requests, error %s", err)
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
		log.Fatalf("Can not create tracer, %s", err)
	}
	return tracer, closer
}
