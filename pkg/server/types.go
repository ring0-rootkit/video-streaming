package udp_server

import (
	"fmt"
	"net"

	"github.com/ring0-rootkit/video-streaming-in-go/pkg/logging"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/network"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/video"
)

const (
	ClientConnBufSize int = 512
)

type UDPServer struct {
	conn      *net.UDPConn
	reader    *video.Reader
	port      int
	isRunning bool

	hasNextCh chan bool
	peers     []*net.UDPAddr

	log *logging.Log
	qc  chan bool
}

func New(port int) *UDPServer {
	log := logging.New(fmt.Sprintf("[server:%d]", port))

	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		//TODO handle error
		panic(err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		//TODO handle error
		panic(err)
	}
	return &UDPServer{
		conn:      conn,
		port:      port,
		isRunning: false,
		peers:     make([]*net.UDPAddr, 8),
		log:       log,
	}
}

// returns 'quit channel'
func (s *UDPServer) Start(filename string) chan bool {
	s.log.Info("started the server")

	s.isRunning = true

	s.reader = video.Load(filename)
	s.hasNextCh = s.reader.StartVideoSync()

	go s.startListener()
	go s.startBroadcaster()
	s.qc = make(chan bool)
	return s.qc
}

func (s *UDPServer) Close() {
	s.conn.Close()
}

func (s *UDPServer) startListener() {
	for s.isRunning {
		var buf [ClientConnBufSize]byte
		_, addr, err := s.conn.ReadFromUDP(buf[:])
		if err != nil {
			continue
		}
		s.peers = append(s.peers, addr)
	}
}

func (s *UDPServer) startBroadcaster() {
	for <-s.hasNextCh {
		network.Broadcast(s.conn, s.peers, s.reader)
	}
	s.isRunning = false
	s.qc <- true
}
