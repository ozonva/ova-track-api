
package api

import (
	"context"
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

	return nil, s.rep.Add([]utils.Track{{0, req.Name, req.Artist, req.Album}})
}