package saver

import (
	"fmt"
	"github.com/ozonva/ova-track-api/internal/flusher"
	"github.com/ozonva/ova-track-api/internal/utils"
	"time"
)

type BufferedSaver struct {
	flusher flusher.Flusher
	bufferedTracks []utils.Track
}

func (bs * BufferedSaver) SaveToBuffer (tracks []utils.Track){
	bs.FlushBuffer()
	bs.bufferedTracks = append(bs.bufferedTracks, tracks...)
}

func (bs * BufferedSaver) FlushBuffer (){
	if bs.bufferedTracks == nil {
		return
	}
	notFlushed := bs.flusher.Flush(bs.bufferedTracks)
	if notFlushed != nil   {
		fmt.Errorf("cannot flush tracks %v", notFlushed)
	}
	bs.bufferedTracks = notFlushed
}

func NewBufferSaver(flusher flusher.Flusher) BufferedSaver {
	return BufferedSaver {flusher, nil}
}

// ====================================================================================

type Saver interface {
	Save(tracks []utils.Track)
	Close()
	Init ()
}

type TimelapseBufferedSaver struct {
	bs BufferedSaver
	timer time.Timer
}

func (tls * TimelapseBufferedSaver) Init (msc int64)  {
	timer := time.NewTimer(time.Duration(msc) * time.Millisecond)
	for true {
		select {
		case <-timer.C:
			tls.bs.FlushBuffer()
		}
	}
}

func (tls * TimelapseBufferedSaver) Save (tracks []utils.Track)  {
	tls.bs.SaveToBuffer(tracks)
}

func (tls * TimelapseBufferedSaver) Close ()  {
	tls.timer.Stop()
	tls.bs.FlushBuffer()
}

func NewTimelapseBufferedSaver (bufferSaver BufferedSaver) TimelapseBufferedSaver{
	return TimelapseBufferedSaver {bufferSaver, time.Timer{}}
}

// ====================================================================================
