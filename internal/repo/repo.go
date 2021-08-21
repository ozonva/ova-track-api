package repo

import "github.com/ozonva/ova-task-api/internal/utils"

type TrackRepo interface {
	Add ([]utils.Track) error
	List (limit, offset uint64) ([]utils.Track, error)
	Describe (id uint64) (*utils.Track, error)
}