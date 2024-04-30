package live_server

import (
	"fmt"
	"net"

	"github.com/haivision/srtgo"
	types "github.com/ring0-rootkit/video-streaming-in-go/pkg"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/live"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/logging"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/network"
)

const (
	ClientConnBufSize int = 512
)

var allowedStreamIDs = map[string]bool{
	"foo": true,
	"bar": true,
}

type SRTServer struct {
	sck        *srtgo.SrtSocket
	readWriter *live.ReadWriter
	port       uint16
	isRunning  bool

	streamer *types.SRTClient

	peers []*types.SRTClient

	log *logging.Log
	qc  chan bool
}

func New(port uint16) *SRTServer {
	log := logging.New(fmt.Sprintf("[server:%d]", port))

	options := make(map[string]string)
	options["blocking"] = "0"
	options["transtype"] = "live"
	options["latency"] = "300"

	sck := srtgo.NewSrtSocket("localhost", port, options)

	srv := &SRTServer{
		sck:       sck,
		port:      port,
		isRunning: false,
		peers:     make([]*types.SRTClient, 0),
		log:       log,
	}

	srv.sck.SetListenCallback(srv.listenCallback)
	return srv
}

// returns 'quit channel'
func (s *SRTServer) Start() chan bool {
	err := s.sck.Listen(1)
	if err != nil {
		//TODO close the srtserver gracefully
		panic(err)
	}
	s.log.Info("started the server")

	s.isRunning = true

	s.readWriter = live.NewReadWriter()

	go s.startListener()
	s.log.Info("waiting for streamer to join")
	s.qc = make(chan bool)
	return s.qc
}

func (s *SRTServer) Close() {
	s.log.Warn("stopping the server")
	s.qc <- true
	s.sck.Close()
	s.readWriter.Close()
	s.isRunning = false
	if s.streamer != nil {
		s.streamer.Socket.Close()
	}
	for _, peer := range s.peers {
		if peer != nil {
			peer.Socket.Close()
		}
	}
}

func (s *SRTServer) listenCallback(socket *srtgo.SrtSocket, version int, addr *net.UDPAddr, streamid string) bool {
	s.log.Info(fmt.Sprintf("socket will connect, hsVersion: %d, streamid: %s\n", version, streamid))

	if _, found := allowedStreamIDs[streamid]; found {
		s.streamer = &types.SRTClient{ClientAddr: addr, Connected: false}
		go s.recieveStream()
		go s.startBroadcaster()
	}
	return true
}

func (s *SRTServer) startListener() {
	for s.isRunning {
		socket, addr, err := s.sck.Accept()
		if err != nil {
			continue
		}
		s.log.Info("new client connected!")
		defer socket.Close()
		if s.streamer != nil && addr.AddrPort() == s.streamer.ClientAddr.AddrPort() && addr.IP.Equal(s.streamer.ClientAddr.IP) {
			s.log.Info("streamer  connected")
			s.streamer.Socket = socket
			s.streamer.Connected = true
		} else {
			s.peers = append(s.peers, &types.SRTClient{Socket: socket, ClientAddr: addr})
		}
	}
}

func (s *SRTServer) startBroadcaster() {
	for s.isRunning {
		err := network.Broadcast(s.peers, s.readWriter)
		if err != nil {
			break
		}
	}
}
