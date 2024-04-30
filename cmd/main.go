package main

import (
	"fmt"
	"time"

	_ "github.com/joho/godotenv/autoload"
	live_server "github.com/ring0-rootkit/video-streaming-in-go/pkg/srt_server/live_server"
)

func main() {
	time.Sleep(time.Second * 2)
	srtServer := live_server.New(42069)
	qc := srtServer.Start("12345678901234567890")
	defer srtServer.Close()
	<-qc
	fmt.Println("close the server")
}
