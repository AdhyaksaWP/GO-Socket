package main

import (
	"fmt"
	"net"
)

func main() {
	var keys [2]float64

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn, keys)
	}
}

func handleConnection(conn net.Conn, keys [2]float64) {
	defer conn.Close()

	i := 0
	buf := make([]byte, 1024)
	for i < 2 {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Received: %s\n", buf)
			i++
		}

	}
}
