package api

import (
	"github.com/ozonva/ova-track-api/internal/kafka_client"
	"github.com/ozonva/ova-track-api/internal/repo"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
)

type ApiServer struct {
	desc.UnimplementedTrackServer
	rep repo.TrackRepo
	metrics Metrics
	kafka kafka_client.IKafkaClient
}

func NewApiServer(rep repo.TrackRepo, met Metrics, kfk kafka_client.IKafkaClient) desc.TrackServer {
	return &ApiServer{rep: rep, metrics: met, kafka: kfk}
}