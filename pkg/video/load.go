package video

import (
	"bytes"
)

func Load(filename string) *Reader {
	return NewReader(bytes.NewReader([]byte(filename)), 1000)
}
