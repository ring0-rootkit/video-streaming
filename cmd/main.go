package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	live_server "github.com/ring0-rootkit/video-streaming-in-go/pkg/srt_server/live_server"
)

func main() {
	srtServer := live_server.New(42069)
	qc := srtServer.Start()
	<-qc
	fmt.Println("server closed")
}
