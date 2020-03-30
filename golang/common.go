package main

import (
	"io"
	"net"
)

func ReadSessionMessage(conn net.Conn) (*SessionMessage) {
	bufsize := 8
	buf := make([]byte, bufsize)
	readn := 0
	for readn < bufsize {
		nr, err := conn.Read(buf[readn:bufsize])
		if err != nil {
			if err == io.EOF {
				conn.Close()
				return nil	
			}

			panic(err)
		}

		if nr > 0 {
			readn += nr
		}
	}

	msg := new(SessionMessage)
	PackSessionMessage(buf, msg)

	return msg
}