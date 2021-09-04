package api

import (
	"github.com/ozonva/ova-track-api/internal/repo"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
)

type ApiServer struct {
	desc.UnimplementedTrackServer
	rep repo.TrackRepo
}

func NewApiServer(rep repo.TrackRepo) desc.TrackServer {
	return &ApiServer{rep: rep}
}