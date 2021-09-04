package api

import (
	"context"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)


func (s *ApiServer) RemoveTrackByID(ctx context.Context, dsc *desc.TrackID) (*desc.Empty, error){
	log.Info().Caller().Uint64("id", dsc.TrackId)
	return nil, s.rep.Remove(dsc.TrackId)
}
