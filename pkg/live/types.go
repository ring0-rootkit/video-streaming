package live

import (
	"errors"
	"sync"

	"github.com/ring0-rootkit/video-streaming-in-go/pkg/logging"
)

const (
	BufferSize int16 = 1316
	MaxCached  uint8 = 255
)

var Log *logging.Log = logging.New("[video]")

type VideoChunk struct {
	id    int64
	chunk []byte
}

type ReadWriter struct {
	HasNextCh chan bool

	rwMut  *sync.RWMutex
	cache  []VideoChunk
	curBuf []byte

	currCached uint8

	lastId int64
}

// mspp - milliseconds per packet
func NewReadWriter() *ReadWriter {
	rw := &ReadWriter{
		rwMut:     &sync.RWMutex{},
		HasNextCh: make(chan bool),
		cache:     make([]VideoChunk, 0),
	}
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
	vr.rwMut.RLock()
	defer vr.rwMut.RUnlock()
	n := copy(buf, vr.curBuf)
	return n, nil
}

// func (vr *ReadWriter) Write(buf []byte) (int, error) {
// 	// set flag so everyone knows that streamer is trying to write
// 	if vr.currCached < MaxCached {
// 		tmp := make([]byte, len(buf))
// 		n := copy(tmp, buf)
// 		vr.cache = append(vr.cache, VideoChunk{chunk: tmp, id: vr.lastId})
// 		vr.lastId++
// 		vr.currCached++
// 		return n, nil
// 	}
//
// 	vr.curBuf = make([]byte, len(vr.cache[0].chunk))
// 	vr.rwMut.Lock()
// 	n := copy(vr.curBuf, vr.cache[0].chunk)
// 	vr.rwMut.Unlock()
// 	vr.HasNextCh <- true
//
// 	vr.cache = slices.Delete(vr.cache, 0, 1)
// 	tmp := make([]byte, len(buf))
// 	copy(tmp, buf)
// 	vr.cache = append(vr.cache, VideoChunk{chunk: tmp, id: vr.lastId})
// 	vr.lastId++
// 	return n, nil
// }

func (vr *ReadWriter) Write(buf []byte) (int, error) {
	// set flag so everyone knows that streamer is trying to write
	// vr.rwMut.Lock()
	vr.curBuf = make([]byte, len(buf))
	n := copy(vr.curBuf, buf)
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
