package api

import (
	"context"
	ot "github.com/opentracing/opentracing-go"
	otl "github.com/opentracing/opentracing-go/log"
	"github.com/ozonva/ova-track-api/internal/kafka_client"
	"github.com/ozonva/ova-track-api/internal/utils"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)

func (s *ApiServer) CreateTrack(ctx context.Context, req *desc.TrackDescription) (*desc.Empty, error) {

	log.Info().
		Caller().
		Str("Track", req.Name).
		Str("Album", req.Album).
		Str("Artist", req.Artist).
		Msg("")

	if sendError := kafka_client.SendKafkaCreateEvent(s.kafka); sendError != nil {
		log.Error().Msgf("Can not send create event to kafka, error: %s", sendError)
	}

	span, ctx := ot.StartSpanFromContext(ctx, "CreateTrack")
	span.LogFields(otl.String("creating track", req.String()))
	defer span.Finish()

	addRes := s.repo.Add([]utils.Track{{utils.InitialTrackId, req.Name, req.Album, req.Album}})
	if addRes == nil {
		s.metrics.IncSuccessCreateTrackCounter()
	}
	return nil, addRes
}
