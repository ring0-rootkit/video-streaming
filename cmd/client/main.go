package main

import (
	"bufio"
	"fmt"

	"github.com/haivision/srtgo"
)

const (
	BufSize int32 = 512
)

func main() {
	options := make(map[string]string)
	options["streamid"] = "foo"
	sck := srtgo.NewSrtSocket("localhost", 42069, options)
	err := sck.Connect()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(sck)

	fmt.Println("receiving data from server...")

	for {
		var buf [BufSize]byte
		_, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("cannot read from server")
			panic(err)
		}
		fmt.Println(string(buf[:]))
	}
}
