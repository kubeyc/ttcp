package main

import (
	"fmt"
	"log"
	"net"
)



func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:19999")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	receive(conn)
	// data := make([]byte, 8)
	// PackSessionMessage(data, &SessionMessage{1000, 1024})
	// log.Println(data)

	// sessionMsg := &SessionMessage {0, 0}
	// UnPackSessionMessage(sessionMsg, data)
	// log.Println(sessionMsg)
}

func receive(conn net.Conn) {
	defer conn.Close()
	sessionMsg := NewSessionMessage()
	err := sessionMsg.Read(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("receive buffer length = %d\nreceive number of buffers = %d\n", sessionMsg.Length, sessionMsg.Number)
	fmt.Printf("%.3f MiB in total\n", float64(sessionMsg.Number * sessionMsg.Length * 1.0 / 1024 / 1024))
	
	payloadMsg := NewPayloadMessage()
	var i uint32
	for i = 0; i < sessionMsg.Number; i++ {
		if err = payloadMsg.ReadPayloadLength(conn); err != nil {
			log.Fatal(err)
		}

		// fmt.Println(payloadMsg.Length)
	}
}