package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	message := "Hello, I am a server"
	listener, err := net.Listen("tcp", ":2874")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		messageLength := len(message)

		fmt.Println("Client connected")
		bytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(bytes, uint32(messageLength))
		conn.Write(bytes)
		conn.Write([]byte(message))
	}
}
