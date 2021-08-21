package flusher

import (
	"fmt"
	"github.com/ozonva/ova-task-api/internal/utils"
	"github.com/ozonva/ova-task-api/internal/repo"
)

type Flusher interface {
	Flush ([]utils.Track)[]utils.Track
}

type ChunkFlusher struct {
	chunkSize uint32
	repo repo.TrackRepo
}

func (flusher * ChunkFlusher) Flush (tracks []utils.Track)[]utils.Track {
	err := flusher.repo.Add (tracks)
	if err == nil{
		fmt.Printf("repository can't add tracks. reason %s", err)
		return []utils.Track{}
	}
	return tracks
}

