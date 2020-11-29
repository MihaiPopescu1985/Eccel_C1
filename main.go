package main

import (
	"fmt"

	"example.com/c1/c1device"
)

func main() {

	operational := c1device.C1Device{
		Name:      "operational",
		IP:        "192.168.0.10",
		TCPPort:   "8080",
		WsChannel: make(chan []byte),
	}

	operational.WsRead()

	go func() {
		for b := range operational.WsChannel {
			msg := string(b)
			fmt.Println(msg)
		}
	}()

	//time.Sleep(time.Second * 1)
	//log.Println("tcp connection:")
	//go c1device.TCPControl(ipAddress + ":8080")

	//go c1device.DbConnect()

	fmt.Scanln()
}
