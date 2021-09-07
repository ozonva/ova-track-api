package api

import (
	"context"
	kafka "github.com/ozonva/ova-track-api/internal/kafka_client"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)

func (s *ApiServer) UpdateTrack(ctx context.Context, req *desc.TrackUpdateData) (*desc.Empty, error) {
	log.Info().Msgf("Receive new update request: %s", req.String())

	if sendError := kafka.SendKafkaUpdateEvent(s.kafka); sendError != nil {
		log.Error().Msgf("Can not send update event to kafka, error: %s", sendError)
	}

	_, errRemove := s.RemoveTrackByID(ctx, req.Id)
	if errRemove != nil {
		return nil, errRemove
	}

	_, errAdd := s.CreateTrack(ctx, req.Description)
	if errAdd != nil {
		return nil, errAdd
	}

	return &desc.Empty{}, nil
}
