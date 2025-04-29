package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

func main() {
	var keys = []map[string]interface{}{}

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

		go handleConnection(conn, &keys)
	}
}

func checkKeys(keys *[]map[string]interface{}) {
	if len(*keys) == 2 {
		conn1 := (*keys)[0]["Conn"].(net.Conn)
		conn2 := (*keys)[1]["Conn"].(net.Conn)

		map1 := mapWithoutConn((*keys)[0])
		map2 := mapWithoutConn((*keys)[1])

		var buf1 bytes.Buffer
		enc1 := gob.NewEncoder(&buf1)
		err1 := enc1.Encode(map1)
		if err1 != nil {
			fmt.Println(err1)
		}

		var buf2 bytes.Buffer
		enc2 := gob.NewEncoder(&buf2)
		err2 := enc2.Encode(map2)
		if err2 != nil {
			fmt.Println(err2)
		}

		conn1.Write(buf2.Bytes())
		conn2.Write(buf1.Bytes())

		conn1.Close()
		conn2.Close()

		for i := len(*keys); i > 0; i-- {
			*keys = (*keys)[:len(*keys)-1]
		}
		// if len(*keys) > 0 {
		// }
	}
}

func mapWithoutConn(keys map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, value := range keys {
		if key != "Conn" {
			newMap[key] = value
		}
	}
	return newMap
}

func handleConnection(conn net.Conn, keys *[]map[string]interface{}) {
	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	encodedData := buf[:n]

	var decodeMap map[string]interface{}
	dec := gob.NewDecoder(bytes.NewReader(encodedData))
	err = dec.Decode(&decodeMap)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	fmt.Printf("Decoded map: %v\n", decodeMap)

	name := decodeMap["Name"]
	g := decodeMap["G"]
	p := decodeMap["P"]
	A := decodeMap["A"]

	decodeMap["Conn"] = conn

	fmt.Printf("Name= %s, G = %.2f, P = %.2f, A = %.2f\n", name, g, p, A)

	*keys = append(*keys, decodeMap)
	checkKeys(keys)
}
