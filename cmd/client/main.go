package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/haivision/srtgo"
)

const (
	BufSize int32 = 1316
)

func main() {
	options := make(map[string]string)
	sck := srtgo.NewSrtSocket("localhost", 42069, options)
	for {
		err := sck.Connect()
		if err != nil {
			continue
		}
		break
	}
	reader := bufio.NewReader(sck)

	fmt.Println("receiving data from server...")
	file, err := os.OpenFile("sample_videos/recieved.h264", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	for {
		buf := make([]byte, BufSize)
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println("cannot read from server")
			panic(err)
		}
		file.Write(buf[:n])
	}
}
