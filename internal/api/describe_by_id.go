package api

import (
	"context"
	ot "github.com/opentracing/opentracing-go"
	otl "github.com/opentracing/opentracing-go/log"
	"github.com/ozonva/ova-track-api/internal/kafka_client"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)

func (s *ApiServer) DescribeTrackByID (ctx context.Context, trackId *desc.TrackID) (*desc.TrackDescription, error) {

	log.Info().Caller().Uint64("Id", trackId.TrackId)

	if sendError := kafka_client.SendKafkaGetEvent(s.kafka); sendError != nil {
		log.Error().Msgf("Can not send get event to kafka, error: %s", sendError)
	}

	span, ctx := ot.StartSpanFromContext(ctx, "DescribeTrackByID")
	span.LogFields(otl.Uint64("describing track", trackId.TrackId))
	defer span.Finish()

	trk, err := s.repo.Describe(trackId.TrackId)
	if err == nil {
		s.metrics.IncGetTrackCounter()
		return &desc.TrackDescription {Name: trk.TrackName, Artist: trk.Artist, Album: trk.Album}, nil
	}
	return &desc.TrackDescription {Name: "", Artist: "", Album: ""}, err
}
