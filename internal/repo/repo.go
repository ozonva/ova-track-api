package repo

import (
	"errors"
	"fmt"
	"github.com/ozonva/ova-track-api/internal/utils"
)

type TrackRepo interface {
	Add ([]utils.Track) error
	List (limit, offset uint64) ([]utils.Track, error)
	Describe (id uint64) (*utils.Track, error)
}

type PrintRepo struct {}

func (pr PrintRepo) Add (tracks []utils.Track) error{
	fmt.Println("Add called", tracks)
	return errors.New ("")
}
func (pr PrintRepo) List (limit, offset uint64) ([]utils.Track, error){
	fmt.Printf("")
	return nil,nil
}
func (pr PrintRepo) Describe (id uint64) (*utils.Track, error){
	fmt.Printf("")
	return nil,nil
}
