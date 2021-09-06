package api

import (
	"context"
	"github.com/ozonva/ova-track-api/internal/kafka_client"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)


func (s *ApiServer) RemoveTrackByID(ctx context.Context, dsc *desc.TrackID) (*desc.Empty, error){

	if sendError := kafka_client.SendKafkaDeleteEvent(s.kafka); sendError != nil {
		log.Error().Msgf("Can not send get event to kafka, error: %s", sendError)
	}

	log.Info().Caller().Uint64("id", dsc.TrackId)
	removeRes := s.rep.Remove(dsc.TrackId)
	if removeRes == nil{
		s.metrics.IncSuccessRemoveTrackCounter()
	}
	return nil, removeRes
}
