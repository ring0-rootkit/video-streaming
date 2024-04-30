package network

import (
	"fmt"

	types "github.com/ring0-rootkit/video-streaming-in-go/pkg"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/live"
)

// send 'r' to every peer in 'peers'
func Broadcast(peers []*types.SRTClient, r *live.ReadWriter) error {
	buf := make([]byte, live.BufferSize)
	n, err := r.Read(buf)
	if err == nil {
		for _, peer := range peers {
			stat, _ := peer.Socket.Stats()
			Log.Info(fmt.Sprintf("bandwitdh:%f, dropped-packets:%d, byte-loss:%d",
				stat.MbpsBandwidth,
				stat.PktSndDrop,
				stat.ByteRcvLossTotal))

			peer.Socket.Write(buf[:n])
		}
		return nil
	}
	if err.Error() == "EOS" {
		return err
	}
	Log.Error(fmt.Sprintf("error while sending to client: %s", err.Error()))
	return err
}
