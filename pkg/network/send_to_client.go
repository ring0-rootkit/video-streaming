package network

import (
	"fmt"
	"net"

	"github.com/ring0-rootkit/video-streaming-in-go/pkg/video"
)

// send 'r' from 'Conn.UDPConn' to 'Conn.UDPAddrs'
func Broadcast(udpConn *net.UDPConn, udpAddrs []*net.UDPAddr, r *video.Reader) {
	var buf [video.BufSize]byte
	_, err := r.Read(buf[:])
	if err != nil {
		Log.Error(fmt.Sprintf("error while sending to client: %s", err.Error()))
		return
	}
	for _, addr := range udpAddrs {
		udpConn.WriteToUDP(buf[:], addr)
	}
	Log.Info("wrote :" + string(buf[:]))
}
