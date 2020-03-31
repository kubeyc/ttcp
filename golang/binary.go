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
	Data   []byte
}

func NewPayloadMessage() *PayloadMessage {
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

func (p *PayloadMessage) ReadBySession(conn net.Conn, msg *SessionMessage) error {
	buf := make([]byte, 4)
	nr, err := conn.Read(buf)
	if err != nil {
		if err == io.EOF {
			return err
		}

		panic(err)
	}

	if nr == 0 {
		return fmt.Errorf("read PayloadMessage Length error, client closed")
	}

	p.Length = binary.BigEndian.Uint32(buf)
	if p.Length != msg.Length {
		return fmt.Errorf("read PayloadMessage Length error, payload message length %d != session message length %d", p.Length, msg.Length)
	}

	readn, needn := 0, int(p.Length)
	buf = make([]byte, p.Length)
	for readn < needn {
		nr, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return err
			}
		}

		readn += nr
	}

	if readn != needn {
		return fmt.Errorf("read PayloadMessage data error, need %d, give %d", needn, readn)
	}

	return nil
}

func (p *PayloadMessage) WriteAck(conn net.Conn) error {
	ack := make([]byte, 4)
	binary.BigEndian.PutUint32(ack, p.Length)
	nr, err := conn.Write(ack)
	if err != nil {
		panic(err)
	}

	if nr != 4 {
		return fmt.Errorf("write ack to client error, need write 4 != %d", nr)
	}

	return nil
}

func (p *PayloadMessage) PackBinary() []byte {
	binary.LittleEndian.PutUint32(p.Data, p.Length)
	return p.Data
}
