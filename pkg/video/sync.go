package video

import (
	"time"
)

func (r *Reader) StartVideoSync() {
	for {
		var buf [BufSize]byte
		_, err := r.Read(buf[:])
		if err == nil {
			copy(r.curBuf[:], buf[:])
			time.Sleep(time.Microsecond * time.Duration(r.MillisecondsPerPack()))
			continue
		}
		if err.Error() == "EOF" {
			return
		}
		// TODO error handling
		panic(err)
	}
}
