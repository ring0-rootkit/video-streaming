package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	BufSize int32 = 512
)

func main() {
	servAddp, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		fmt.Println("cannot resolve serv addr")
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, servAddp)
	if err != nil {
		fmt.Println("cannot connect to the server")
		panic(err)
	}

	fmt.Println("connected to server, listening...")

	_, err = conn.Write([]byte("i wanna connect"))
	if err != nil {
		fmt.Println("cannot write to server")
		panic(err)
	}

	reader := bufio.NewReader(conn)

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
