package api

import (
	"context"
	"errors"
	desc "github.com/ozonva/ova-track-api/pkg/api/github.com/ova-track-api/pkg/ova-track-api"
	"github.com/rs/zerolog/log"
)

func (s *ApiServer) GetTrackID (ctx context.Context, td *desc.TrackDescription) (*desc.TrackID, error) {

	log.Info().Caller().Str("name", td.Name).Str("album", td.GetAlbum()).Str("artist", td.GetArtist())
	tracks, err := s.rep.List(18446744073709551615,0)
	if err != nil {
		return &desc.TrackID{TrackId: 0}, err
	}

	for _, trk := range tracks{
		if trk.TrackName == td.GetName() && trk.Album == td.GetAlbum() && trk.Artist == td.GetArtist(){
			return &desc.TrackID{TrackId: trk.TrackId}, nil
		}
	}

	return &desc.TrackID{TrackId: 0}, errors.New("cant find track with such id")
}
