package live

// ffmpeg -re -i srt://localhost:42069 -c:v h264 -c:a copy -vf setpts='PTS-STARTPTS'  sample_videos/output.mp4
import (
	"errors"
	// "slices"
	"sync"

	"github.com/ring0-rootkit/video-streaming-in-go/pkg/logging"
)

const (
	BufSize   int32 = 1316
	MaxCached uint8 = 255
)

var Log *logging.Log = logging.New("[video]")

type VideoChunk struct {
	chunk [BufSize]byte
	id    int64
}

type ReadWriter struct {
	HasNextCh chan bool

	rwMut  *sync.RWMutex
	cache  []VideoChunk
	curBuf [BufSize]byte

	currCached uint8

	lastId int64
}

// mspp - milliseconds per packet
func NewReadWriter() *ReadWriter {
	rw := &ReadWriter{rwMut: &sync.RWMutex{}}
	rw.HasNextCh = make(chan bool)
	rw.cache = make([]VideoChunk, MaxCached)
	return rw
}

// reads 'video.BufSize' bytes to buf, if len(buf) less than video.BufSize
// the error is returned and nothing is written to buf
//
// returns EOS (end of stream) when streamer stops sending packets
func (vr *ReadWriter) Read(buf []byte) (int, error) {
	if hasNext := <-vr.HasNextCh; !hasNext {
		return -1, errors.New("EOS")
	}
	if len(buf) < int(BufSize) {
		return -1, errors.New("len(buf) should be equal or more than video.BufSize")
	}
	// vr.rwMut.RLock()
	// defer vr.rwMut.RUnlock()
	n := copy(buf, vr.curBuf[:])
	return n, nil
}

// func (vr *ReadWriter) Write(buf []byte) (int, error) {
// 	// set flag so everyone knows that streamer is trying to write
// 	if vr.currCached < MaxCached {
// 		var tmp [BufSize]byte
// 		n := copy(tmp[:], buf)
// 		vr.cache = append(vr.cache, VideoChunk{chunk: tmp, id: vr.lastId})
// 		vr.lastId++
// 		vr.currCached++
// 		return n, nil
// 	}
//
// 	vr.rwMut.Lock()
// 	n := copy(vr.curBuf[:], vr.cache[0].chunk[:])
// 	vr.rwMut.Unlock()
// 	vr.HasNextCh <- true
//
// 	vr.cache = slices.Delete(vr.cache, 0, 1)
// 	var tmp [BufSize]byte
// 	copy(tmp[:], buf)
// 	vr.cache = append(vr.cache, VideoChunk{chunk: tmp, id: vr.lastId})
// 	vr.lastId++
// 	return n, nil
// }

func (vr *ReadWriter) Write(buf []byte) (int, error) {
	// set flag so everyone knows that streamer is trying to write
	// vr.rwMut.Lock()
	n := copy(vr.curBuf[:], buf)
	// vr.rwMut.Unlock()
	vr.HasNextCh <- true
	return n, nil
}

func (vr *ReadWriter) Close() {
	vr.HasNextCh <- false
	Log.Info("ReadWriter closed")
}

// func (vr *ReadWriter) MillisecondsPerPack() int64 {
// 	return vr.mspp
// }
