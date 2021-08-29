package main

import (
	"github.com/ozonva/ova-track-api/internal/utils"
	track_server "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-conference-api/pkg/ova-conference-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

var id = uint64(0)
func GenerateTrack (n int,  name string)  utils.Track {
	res := ""
	for i := 0; i < n; i++{
		res+=name
	}
	id++
	return utils.Track{TrackId: id, TrackName: res, Album: res, Artist: res}
}

//fmt.Println("Hi, i am ova-track-api!")
//if len(os.Args) != 2 {
//	fmt.Println("Path to config is strictly required")
//	return
//}
//path := os.Args[1]
//utils.InitLibraryFromFile(path)


func main() {

	port := ":"+strconv.Itoa(8080)
	listen, err := net.Listen("tcp", port)
	log.Printf("TCP at %v is starting up ...", port)

	if err != nil {
		log.Printf("Failed to listen server %v", err)
	}

	log.Printf("TCP at %v started successfully", port)


	grpcService := grpc.NewServer()
	track_server.RegisterTrackServer(grpcService, track_server.UnimplementedTrackServer{})

	log.Print("Starting track service")
	if err := grpcService.Serve(listen); err != nil {
		log.Printf("failed to serve: %v", err)
	}
	log.Printf ("Track service started successfully")
}