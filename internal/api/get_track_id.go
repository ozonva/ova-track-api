package api

import (
	"context"
	"errors"
	ot "github.com/opentracing/opentracing-go"
	otl "github.com/opentracing/opentracing-go/log"
	"github.com/ozonva/ova-track-api/internal/kafka_client"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)

func (s *ApiServer) GetTrackID (ctx context.Context, td *desc.TrackDescription) (*desc.TrackID, error) {

	if sendError := kafka_client.SendKafkaGetEvent(s.kafka); sendError != nil {
		log.Error().Msgf("Can not send get event to kafka, error: %s", sendError)
	}

	log.Info().Caller().Str("name", td.Name).Str("album", td.GetAlbum()).Str("artist", td.GetArtist())
	tracks, err := s.rep.List(18446744073709551615,0)
	span, ctx := ot.StartSpanFromContext(ctx, "DescribeTrackByID")
	span.LogFields(otl.String("looking for a id of track", td.GetName()))
	defer span.Finish()


	if err != nil {
		return &desc.TrackID{TrackId: 0}, err
	}
	s.metrics.IncGetTrackCounter()


	for _, trk := range tracks{
		if trk.TrackName == td.GetName() && trk.Album == td.GetAlbum() && trk.Artist == td.GetArtist(){
			return &desc.TrackID{TrackId: trk.TrackId}, nil
		}
	}

	return &desc.TrackID{TrackId: 0}, errors.New("cant find track with such id")
}
