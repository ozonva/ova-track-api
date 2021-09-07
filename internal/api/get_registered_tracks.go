package api

import (
	"context"
	ot "github.com/opentracing/opentracing-go"
	otl "github.com/opentracing/opentracing-go/log"
	"github.com/ozonva/ova-track-api/internal/kafka_client"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)

func (s *ApiServer) GetRegisteredTracks(ctx context.Context, em *desc.Empty) (*desc.TracksDescriptions, error) {

	if sendError := kafka_client.SendKafkaGetEvent(s.kafka); sendError != nil {
		log.Error().Msgf("Can not send get event to kafka, error: %s", sendError)
	}

	span, ctx := ot.StartSpanFromContext(ctx, "GetRegisteredTracks")
	span.LogFields(otl.Int("Get registered tracks", 1))
	defer span.Finish()

	log.Info().Caller()
	tracks, err := s.repo.List(18446744073709551615, 0)
	if err != nil {
		return &desc.TracksDescriptions{}, err
	}
	s.metrics.IncGetTrackCounter()
	res := desc.TracksDescriptions{}
	for _, trk := range tracks {
		res.Item = append(res.Item, &desc.TrackDescription{Name: trk.TrackName, Artist: trk.Artist, Album: trk.Album})
	}

	return &res, nil
}
