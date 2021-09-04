package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ozonva/ova-track-api/internal/api"

	//"github.com/ozonva/ova-track-api/internal/api"
	"github.com/ozonva/ova-track-api/internal/repo"

	track_server "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"google.golang.org/grpc"

	"log"
	"net"
	"strconv"
)


func main() {

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
	track_server.RegisterTrackServer(grpcService, api.NewApiServer(rp))

	log.Print("Starting track service")
	if err := grpcService.Serve(listen); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}