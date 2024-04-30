package live_server

import (
	"fmt"
	"net"

	"github.com/haivision/srtgo"
	types "github.com/ring0-rootkit/video-streaming-in-go/pkg"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/logging"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/network"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/video"
)

const (
	ClientConnBufSize int = 512
)

var allowedStreamIDs = map[string]bool{
	"foo":    true,
	"foobar": true,
}

type SRTServer struct {
	sck       *srtgo.SrtSocket
	reader    *video.Reader
	port      uint16
	isRunning bool

	hasNextCh chan bool
	peers     []*types.SRTPeer

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
		peers:     make([]*types.SRTPeer, 0),
		log:       log,
	}

	srv.sck.SetListenCallback(srv.listenCallback)
	return srv
}

// returns 'quit channel'
func (s *SRTServer) Start(filename string) chan bool {
	err := s.sck.Listen(1)
	if err != nil {
		//TODO close the srtserver gracefully
		panic(err)
	}
	s.log.Info("started the server")

	s.isRunning = true

	s.reader = video.Load(filename)
	s.hasNextCh = s.reader.StartVideoSync()

	go s.startListener()
	go s.startBroadcaster()
	s.qc = make(chan bool)
	return s.qc
}

func (s *SRTServer) Close() {
	s.sck.Close()
}

func (s *SRTServer) listenCallback(socket *srtgo.SrtSocket, version int, addr *net.UDPAddr, streamid string) bool {
	s.log.Info(fmt.Sprintf("socket will connect, hsVersion: %d, streamid: %s\n", version, streamid))

	// socket not in allowed ids -> reject
	if _, found := allowedStreamIDs[streamid]; !found {
		// set custom reject reason
		socket.SetRejectReason(srtgo.RejectionReasonUnauthorized)
		return false
	}

	// allow connection
	return true
}

func (s *SRTServer) startListener() {
	for s.isRunning {
		socket, addr, err := s.sck.Accept()
		s.log.Info("new client connected!")
		defer socket.Close()
		if err != nil {
			continue
		}
		s.peers = append(s.peers, &types.SRTPeer{Socket: socket, PeerAddr: addr})
	}
}

func (s *SRTServer) startBroadcaster() {
	for <-s.hasNextCh {
		network.Broadcast(s.peers, s.reader)
	}
	s.isRunning = false
	s.qc <- true
}
