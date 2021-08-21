package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var artistsLibrary map[string]map[string]map[uint64]struct{}
var generalLibrary map[uint64]Track

func resetLibrary () {
	artistsLibrary = nil
	generalLibrary = nil
}

type Track struct
{
	TrackId uint64
	TrackName string
	Album string
	Artist string
}

func New(TrackId uint64, artistName string, albumName string, trackName string) (Track,error) {

	if generalLibrary == nil {
		generalLibrary = make (map[uint64]Track)
		artistsLibrary = make (map[string]map[string]map[uint64]struct{})
	}

	if _, exists := generalLibrary[TrackId]; exists{
		return Track{}, errors.New(fmt.Sprintf("Track with id %d already exists", TrackId))
	}

	if _, exists := artistsLibrary[artistName]; !exists{
		artistsLibrary[artistName] = make (map[string]map[uint64]struct{})
	}


	if _, exists := artistsLibrary[artistName][albumName]; !exists{
		artistsLibrary[artistName][albumName] = make (map[uint64]struct{})
	}
	artistsLibrary[artistName][albumName][TrackId] =struct{}{}
	res := Track{TrackId, trackName,albumName, artistName}
	generalLibrary[TrackId] = res
	return res, nil
}

func GetArtistTracks (artistName string)[] Track {
	res:= make([]Track, 0)
	if _, exists := artistsLibrary[artistName]; !exists{
		return res
	}

	for keyAlbum, _ := range artistsLibrary[artistName] {
		for keyTrack, _ := range artistsLibrary[artistName][keyAlbum] {
			res = append (res, generalLibrary[keyTrack])
		}
	}
	return res
}

func GetArtistAlbumTracks (artistName string, albumName string)[] Track {
	var res []Track
	if _, exists := artistsLibrary[artistName]; !exists{
		return res
	}

	if _, exists := artistsLibrary[artistName][albumName]; !exists{
		return res
	}
	for keyTrack, _ := range artistsLibrary[artistName][albumName] {
		res = append (res, generalLibrary[keyTrack])
	}
	return res
}

func GetTrackById(id uint64) (Track, bool) {
	track, exist := generalLibrary[id]
	if !exist{
		return Track{0, "","",""}, false
	}
	return track, true
}

func initLibraryFromArray (lines []string) []Track {
	res := make ([]Track, 0)
	if lines == nil {
		return res
	}
	nextId := uint64(0)
	for _, val := range lines {

		splitRiddenLine := strings.Split(val, ",")
		if len(splitRiddenLine)!=3 {
			panic (fmt.Sprintf("Can't parse line %s", val))
		}
		track, err := New(nextId, splitRiddenLine[0], splitRiddenLine[1], splitRiddenLine[2])
		nextId++
		if err !=nil {
			panic (fmt.Sprintf("Can't add track to library %s. Error %s", val, err))
		}
		res = append(res, track)
	}
	return res
}

func InitLibraryFromFile (filePath string) (tracks []Track){

	file, fileErr := os.Open(filePath)
	if fileErr != nil{
		fmt.Printf("Cant open file. Error %s", fileErr)
		return nil
	}

	defer func() {
		file.Close()
		resetLibrary()
		if libraryErr := recover(); libraryErr != nil{
			fmt.Printf("Cant init library from file %s. Error %s", filePath, libraryErr)
		}
	}()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	tracks = initLibraryFromArray (lines)
	return tracks
}


