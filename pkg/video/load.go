package video

import (
	"bytes"
)

func Load() *Reader {
	return NewReader(bytes.NewReader([]byte("Hello world")), 1000)
}
