package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math"
	"net"
)

func power(g float64, p float64, a int) float64 {
	return float64(int(math.Pow(g, p)) % a)
}

func main() {
	var g, p, A float64
	var a int
	var name string

	fmt.Print("Set the name, public, and private key key (Name, G, P, a): ")
	fmt.Scanf("%s %f %f %d", &name, &g, &p, &a)

	A = power(g, p, a)

	m := make(map[string]interface{})
	m["Name"] = name
	m["G"] = g
	m["P"] = p
	m["A"] = A

	// Encode the map
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	err := e.Encode(m)
	if err != nil {
		fmt.Println("encode error:", err)
	}

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	encodedData := b.Bytes()

	_, err = conn.Write([]byte(encodedData))
	if err != nil {
		fmt.Println(err)
		return
	}

	var decodeMap map[string]interface{}
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err == nil {
			dec := gob.NewDecoder(bytes.NewReader(buf[:n]))
			err = dec.Decode(&decodeMap)
			if err == nil {
				fmt.Println("Receive decoded map: ", decodeMap)
				break
			} else {
				continue
			}
		}
	}

	conn.Close()
}
