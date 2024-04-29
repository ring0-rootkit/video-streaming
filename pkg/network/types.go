package network

import (
	"net"

	"github.com/ring0-rootkit/video-streaming-in-go/pkg/logging"
)

const (
	BufSize int32 = 512
)

var Log logging.Log = *logging.New("[network]")

type Conn struct {
	UDPCon  *net.UDPConn
	UDPAddr *net.UDPAddr
}
