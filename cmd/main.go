package main

import (
	"fmt"
	"time"

	_ "github.com/joho/godotenv/autoload"
	live_server "github.com/ring0-rootkit/video-streaming-in-go/pkg/udp_server/live_server"
)

func main() {
	time.Sleep(time.Second * 5)
	udpServer := live_server.New(42069)
	qc := udpServer.Start("12345678901234567890")
	defer udpServer.Close()
	fmt.Println("close the server")
	<-qc
}
