package flusher

import (
	"fmt"
	"github.com/ozonva/ova-track-api/internal/repo"
	"github.com/ozonva/ova-track-api/internal/utils"
)

type Flusher interface {
	Flush ([]utils.Track)[]utils.Track
}

type ChunkFlusher struct {
	chunkSize int
	repo repo.TrackRepo
}

func (flusher ChunkFlusher) Flush (tracks []utils.Track)[]utils.Track {

	failedToAdd := make([]utils.Track, 0)
	curSlice := make([]utils.Track, 0, flusher.chunkSize)

	for i, _ := range tracks{
		curSlice = append(curSlice, tracks[i])
		if (((i + 1) % flusher.chunkSize) == 0) || (i + 1 == len (tracks)) {
			fmt.Println("!!! Call", curSlice, i, flusher.chunkSize )
			err := flusher.repo.Add(curSlice)
			if err != nil {
				failedToAdd = append(failedToAdd, curSlice...)
				fmt.Println("!!! Cur nit added", failedToAdd )
			}
			curSlice = curSlice[:0]
		}
	}
	fmt.Println("!!! Not added", failedToAdd)

	fmt.Println("===============================")

	if len (failedToAdd) == 0 {
		return nil
	}
	return failedToAdd
}

