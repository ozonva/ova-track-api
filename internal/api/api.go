package api

import (
	"github.com/ozonva/ova-track-api/internal/kafka_client"
	"github.com/ozonva/ova-track-api/internal/repo"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
)

type ApiServer struct {
	desc.UnimplementedTrackServer
	repo repo.TrackRepo
	metrics Metrics
	kafka kafka_client.KafkaClientInterface
}

func NewApiServer(repo repo.TrackRepo, metrics Metrics, kafka kafka_client.KafkaClientInterface) desc.TrackServer {
	return &ApiServer{repo: repo, metrics: metrics, kafka: kafka}
}