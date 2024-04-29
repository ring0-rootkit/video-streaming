package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	udp_server "github.com/ring0-rootkit/video-streaming-in-go/pkg/server"
)

func main() {
	udpServer := udp_server.New(42069)
	qc := udpServer.Start("Hello world")
	defer udpServer.Close()
	fmt.Println("close the server")
	<-qc
}
