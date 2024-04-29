package main

import (
	"net"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/logging"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/network"
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/video"
)

func main() {
	log := logging.New("[main]")

	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		panic(err)
	}

	serv, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}

	log.Info("started the server")

	reader := video.Load()
	reader.StartVideoSync()

	for {
		var buf [network.BufSize]byte
		_, addr, err := serv.ReadFromUDP(buf[:])
		if err != nil {
			//TODO handle error gracefully
			panic(err)
		}

		con := &network.Conn{UDPCon: serv, UDPAddr: addr}
		go network.Send(con, reader)
	}
}
