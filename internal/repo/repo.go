package repo

import "github.com/ozonva/ova-track-api/internal/utils"

type TrackRepo interface {
	Add ([]utils.Track) error
	List (limit, offset uint64) ([]utils.Track, error)
	Describe (id uint64) (*utils.Track, error)
}