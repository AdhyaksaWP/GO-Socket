package main

import (
	"fmt"
	"math"
	"net"
)

func power(g float64, p float64, a int) int {
	return int(math.Pow(g, p)) % a
}

func main() {
	var g, p float64
	var a, A int

	fmt.Print("Set the public and private key key (G, P, a): ")
	fmt.Scanf("%f %f %d", &g, &p, &a)

	A = power(g, p, a)

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	exchange_keys := fmt.Sprintf("%d, %d, %d", g, p, A)
	_, err = conn.Write([]byte(exchange_keys))
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Close()
}
