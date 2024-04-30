package types

import (
	"net"

	"github.com/haivision/srtgo"
)

type SRTPeer struct {
	Socket   *srtgo.SrtSocket
	PeerAddr *net.UDPAddr
}
