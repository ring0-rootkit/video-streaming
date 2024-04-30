package network

import (
	"fmt"

	types "github.com/ring0-rootkit/video-streaming-in-go/pkg"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/video"
)

// send 'r' to every peer in 'peers'
func Broadcast(peers []*types.SRTPeer, r *video.Reader) {
	var buf [video.BufSize]byte
	_, err := r.Read(buf[:])
	if err != nil {
		Log.Error(fmt.Sprintf("error while sending to client: %s", err.Error()))
		return
	}
	for _, peer := range peers {
		peer.Socket.Write(buf[:])
	}
	Log.Info("wrote :" + string(buf[:]))
}
