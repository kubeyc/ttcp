package main

import (
	"fmt"
	"log"
	"net"
	"time"
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

	totalMib := float64(sessionMsg.Number * sessionMsg.Length * 1.0 / 1024 / 1024)
	fmt.Printf("receive buffer length = %d\nreceive number of buffers = %d\n", sessionMsg.Length, sessionMsg.Number)
	fmt.Printf("%.3f MiB in total\n", totalMib)

	payloadMsg := NewPayloadMessage()
	var i uint32
	start := time.Now()
	for i = 0; i < sessionMsg.Number; i++ {
		if err = payloadMsg.ReadBySession(conn, sessionMsg); err != nil {
			log.Fatal(err)
		}

		if err = payloadMsg.WriteAck(conn); err != nil {
			log.Fatal(err)
		}
	}
	elasped := time.Since(start)

	fmt.Printf("%.3f seconds\n%.3f MiB/s\n", elasped.Seconds(), totalMib/elasped.Seconds())
}
