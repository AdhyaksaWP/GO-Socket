package main

import (
	"fmt"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Sends data 10 times
	for x := 0; x<10; x++ {
		_, err = conn.Write([]byte("Hello, server!"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	conn.Close()
}