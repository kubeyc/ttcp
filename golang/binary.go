package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)


type SessionMessage struct {
	Number uint32
	Length uint32
}

func NewSessionMessage() *SessionMessage {
	return new(SessionMessage)
}

func (s *SessionMessage) Read(conn net.Conn) error {
	bufsize := 8
	buf := make([]byte, bufsize)
	readn := 0
	for readn < bufsize {
		nr, err := conn.Read(buf[readn:bufsize])
		if err != nil {
			if err == io.EOF {
				return fmt.Errorf("read SessionMessage fail, client closed")
			}

			panic(err)
		}

		if nr > 0 {
			readn += nr
		}
	}

	s.Number = binary.BigEndian.Uint32(buf[:4])
	s.Length = binary.BigEndian.Uint32(buf[4:])
	return nil
}

func PackSessionMessage(buf []byte, msg *SessionMessage) {
	binary.LittleEndian.PutUint32(buf[:4], msg.Number)
	binary.LittleEndian.PutUint32(buf[4:], msg.Length)
}

func UnPackSessionMessage(msg *SessionMessage, data []byte) {
	msg.Number = binary.LittleEndian.Uint32(data[:4])
	msg.Length = binary.LittleEndian.Uint32(data[4:])
}

type PayloadMessage struct {
	Length uint32
	Data []byte
}

func NewPayloadMessage() *PayloadMessage{
	return new(PayloadMessage)
	// msg.Length = length
	// msg.Data = make([]byte, 4 + length)

	// // init data
	// var i uint32
	// for i = 0; i < msg.Length; i++ {
	// 	msg.Data[4+i] = "0123456789ABCDEF"[i%16]
	// }

	// return msg
}

func (p *PayloadMessage) ReadPayloadLength(conn net.Conn) error {
	data := make([]byte, 4)
	nr, err := conn.Read(data)
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("read PayloadMessage Length error, client closed")
		}

		panic(err)
	}

	if nr != 4 {
		return fmt.Errorf("read PayloadMessage Length error, need 4 length data, give %d", nr)
	}

	p.Length = binary.BigEndian.Uint32(data[:])

	return nil
}

func (p *PayloadMessage) PackBinary() []byte {
	binary.LittleEndian.PutUint32(p.Data, p.Length)
	return p.Data
}