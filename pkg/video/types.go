package video

import (
	"bytes"
	"errors"

	"github.com/ring0-rootkit/video-streaming-in-go/pkg/logging"
)

const (
	BufSize int32 = 1
)

var Log *logging.Log = logging.New("[video]")

type Reader struct {
	reader *bytes.Reader
	curBuf [BufSize]byte

	// seconds per packet
	mspp int64
}

// mspp - milliseconds per packet
func NewReader(r *bytes.Reader, mspp int64) *Reader {
	return &Reader{reader: r, mspp: mspp}
}

// reads 'video.BufSize' bytes to buf, if len(buf) less than video.BufSize
// the error is returned and nothing is written to buf
func (vr *Reader) Read(buf []byte) (int, error) {
	if len(buf) < int(BufSize) {
		return -1, errors.New("len(buf) should be equal or more than video.BufSize")
	}
	n := copy(buf, vr.curBuf[:])
	return n, nil
}

func (vr *Reader) MillisecondsPerPack() int64 {
	return vr.mspp
}
