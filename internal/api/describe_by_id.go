package api

import (
	"context"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)

func (s *ApiServer) DescribeTrackByID (ctx context.Context, trackId *desc.TrackID) (*desc.TrackDescription, error) {

	log.Info().Caller().Uint64("Id", trackId.TrackId)
	trk, err := s.rep.Describe(trackId.TrackId)
	if err == nil {
		return &desc.TrackDescription {Name: trk.TrackName, Artist: trk.Artist, Album: trk.Album}, nil
	}
	return &desc.TrackDescription {Name: "", Artist: "", Album: ""}, err
}
