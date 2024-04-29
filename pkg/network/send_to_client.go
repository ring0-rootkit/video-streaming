package network

import (
	"fmt"
	"time"

	"github.com/ring0-rootkit/video-streaming-in-go/pkg/video"
)

// send 'r' from 'Conn.UDPConn' to 'Conn.UDPAddr' by chunks
func Send(con *Conn, r *video.Reader) {
	for {
		var buf [video.BufSize]byte
		_, err := r.Read(buf[:])
		if err == nil {
			con.UDPCon.WriteToUDP(buf[:], con.UDPAddr)
			continue
		}
		if err.Error() == "EOF" {
			return
		}
		Log.Error(fmt.Sprintf("error while sending to client: %s", err.Error()))
		time.Sleep(time.Microsecond * time.Duration(r.MillisecondsPerPack()))
	}
}
