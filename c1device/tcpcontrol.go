package c1device

import (
	"net"
	"time"

	"example.com/c1/util"
)

var dummyCommand = []byte{0xF5, 0x03, 0x00, 0xFC, 0xFF, 0x01, 0xD1, 0xF1}

// TCPControl establishes a tcp connection to a C1 device.
func (dev C1Device) TCPControl() {

	address := dev.IP + ":" + dev.TCPPort
	conn, err := net.Dial("tcp", address)

	if err != nil {
		util.Log.Panicln(err)
	}

	msg := make([]byte, 1024)
	duration, err := time.ParseDuration("14s")

	if err != nil {
		util.Log.Panicln(err)
	}
	for {
		time.Sleep(duration)
		conn.Write(dummyCommand)

		lenght, err := conn.Read(msg)

		if err != nil {
			util.Log.Panicln(err)

		} else {
			msg = msg[:lenght]
			util.Log.Panicln(err)
		}
	}
}
