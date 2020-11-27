package tcpcommunication

import (
	"log"
	"net"
	"time"
)

var dummyCommand = []byte{0xF5, 0x03, 0x00, 0xFC, 0xFF, 0x01, 0xD1, 0xF1}

// TCPControl establishes a tcp connection to a C1 device.
// The address must be of type 192.168.0.10:8080
func TCPControl(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
	}
	msg := make([]byte, 9)
	duration, err := time.ParseDuration("10s")
	if err != nil {
		log.Println(err)
	}
	for {
		time.Sleep(duration)
		conn.Write(dummyCommand)

		_, err := conn.Read(msg)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("%X \n", msg)
		}
	}
}
