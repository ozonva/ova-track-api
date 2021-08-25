package main

import (
	"github.com/ozonva/ova-track-api/internal/utils"
	"time"

	//	. "github.com/onsi/ginkgo"
	//	. "github.com/onsi/gomega"
)
import "github.com/ozonva/ova-track-api/internal/flusher"
import "github.com/ozonva/ova-track-api/internal/repo"
import "github.com/ozonva/ova-track-api/internal/saver"

var id = uint64(0)
func GenerateTrack (n int,  name string)  utils.Track {
	res := ""
	for i := 0; i < n; i++{
		res+=name
	}
	id++
	return utils.Track{TrackId: id, TrackName: res, Album: res, Artist: res}
}

//fmt.Println("Hi, i am ova-track-api!")
//if len(os.Args) != 2 {
//	fmt.Println("Path to config is strictly required")
//	return
//}
//path := os.Args[1]
//utils.InitLibraryFromFile(path)

func main() {

	rep := repo.PrintRepo{}
	fl := flusher.NewChunkFlusher(5, rep)
	svr := saver.NewBufferSaver(fl)
	bs := saver.NewTimelapseBufferedSaver(svr)
	go bs.Init(500)

	for i := 0; i < 10000; i++{
		var tracks = []utils.Track{
			GenerateTrack(i+1,"a"),
			GenerateTrack (i+1, "b"),
			GenerateTrack (i+1, "c"),
			GenerateTrack (i+1, "d"),
			GenerateTrack (i+1, "e"),
		}
		bs.Save(tracks)
		time.Sleep(time.Second)
	}
}