package types

import (
	"net"

	"github.com/haivision/srtgo"
)

type SRTClient struct {
	Socket     *srtgo.SrtSocket
	ClientAddr *net.UDPAddr

	HasStreamHeader bool
	Connected       bool
}
