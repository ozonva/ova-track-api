package api

import (
	"context"
	"errors"
	ot "github.com/opentracing/opentracing-go"
	otl "github.com/opentracing/opentracing-go/log"
	"github.com/ozonva/ova-track-api/internal/utils"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)

func (s *ApiServer) MultiCreateTrack(ctx context.Context, req *desc.TracksDescriptions) (*desc.Empty, error) {

	log.Info().Msgf("Multi create track call with parameters: %s", req.String())
	if len(req.Item) == 0{
		return nil, errors.New("multi create call with no items")
	}

	multiTracks := make([]utils.Track, 0, len(req.Item))
	for _, track := range req.Item {
		multiTracks = append(multiTracks, utils.Track{TrackId: 0, TrackName: track.Name, Album: track.Album, Artist: track.Artist})
	}

	span, ctx := ot.StartSpanFromContext(ctx, "MultiCreateTrack")
	span.LogFields(otl.Int("Total recipes count", len(req.Item)))
	defer span.Finish()

	splitIntoBatches := utils.DivideTracksIntoBatches(multiTracks, 2)
	childSpan := ot.StartSpan("MultiCreateTrackBatching", ot.ChildOf(span.Context()))
	for _, batch := range splitIntoBatches {
		childSpan.LogFields(otl.Int("Inserting tracks", len(batch)))
		err := s.repo.Add(batch)
		if err != nil {
			log.Error().Msgf("Can't add new track, error: %s", err)
			return nil, err
		}
	}
	return &desc.Empty{}, nil
}