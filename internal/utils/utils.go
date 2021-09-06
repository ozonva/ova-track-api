package utils

import (
	"fmt"
)

func min (a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func DivideIntoBatches(slice []uint8, batchSize int) [][]uint8 {

	if batchSize <= 0 {
		panic("Incorrect batch size")
	}
	result := make([][]uint8, 0)
	lastBegin := 0
	for i, _ := range slice {
		if (i+1)%batchSize == 0 || i == len (slice) - 1 {
			result = append(result, slice[lastBegin: i + 1])
			lastBegin = i + 1
		}
	}
	return result
}

func DivideTracksIntoBatches(tracks []Track, batchSize int) [][]Track {

	dividedIntoBatches := make([][]Track, 0, 0)
	//curBatch := make([]Track, 0, 0)

	for i := 0; i < len(tracks); i += batchSize {
		next:=i+min(batchSize, len(tracks) - i)
		dividedIntoBatches = append(dividedIntoBatches, []Track{})
		for j := i; j < next; j++ {
			dividedIntoBatches[len(dividedIntoBatches)-1] = append(dividedIntoBatches[len(dividedIntoBatches)-1],
				Track{tracks[j].TrackId,tracks[j].TrackName, tracks[j].Album, tracks[j].Artist })
		}
	}

	return dividedIntoBatches
}

func InverseMap(data map[string]int) map[int]string {

	result := make(map[int]string)
	for key, value := range data {
		_, contains := result[value]
		if contains {
			panic(fmt.Sprintf( "key %d contains twice", value))
		}
		result[value] = key
	}
	return result
}

func FilterMap(data map[string]int, exclude []string) map[string]int {
	for _, val := range exclude {
		_, contains := data[val]
		if contains {
			delete(data, val)
		}
	}
	return data
}
