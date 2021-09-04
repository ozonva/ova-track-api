package api

import (
	"context"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)

func (s *ApiServer) GetRegisteredTracks (context.Context, *desc.Empty) (*desc.TracksDescriptions, error) {

	log.Info().Caller()
	tracks, err := s.rep.List(18446744073709551615,0)
	if err != nil {
		return &desc.TracksDescriptions{}, err
	}

	res := desc.TracksDescriptions {}
	for _, trk := range tracks{
		res.Item = append(res.Item, &desc.TrackDescription {Name: trk.TrackName, Artist: trk.Artist, Album: trk.Album})
	}

	return &res, nil
}
