package flusher

import (
	"github.com/ozonva/ova-track-api/internal/repo"
	"github.com/ozonva/ova-track-api/internal/utils"
	"log"
)

type Flusher interface {
	Flush ([]utils.Track)[]utils.Track
}

type ChunkFlusher struct {
	chunkSize int
	repo repo.TrackRepo
}

func NewChunkFlusher (chunkSize int, repo repo.TrackRepo) ChunkFlusher{
	return ChunkFlusher{chunkSize, repo}
}

func (flusher ChunkFlusher) Flush (tracks []utils.Track)[]utils.Track {

	failedToAdd := make([]utils.Track, 0)
	curSlice := make([]utils.Track, 0, flusher.chunkSize)

	for i, _ := range tracks{
		curSlice = append(curSlice, tracks[i])
		if (((i + 1) % flusher.chunkSize) == 0) || (i + 1 == len (tracks)) {
			err := flusher.repo.Add(curSlice)
			if err != nil {
				failedToAdd = append(failedToAdd, curSlice...)
			}
			curSlice = curSlice[:0]
		}
	}

	if len (failedToAdd) == 0 {
		return nil
	}
	log.Println("can't flush tracks ", failedToAdd)
	return failedToAdd
}

