package main

import (
	"fmt"
	"log"
	"time"

	"example.com/c1/c1device"
)

const operational string = "ws://192.168.0.10/wscomm.cgi"

func main() {
	fmt.Println("muie psd")
	//c1commands.PrintCommands()

	operationalChannel := make(chan []byte)
	c1device.ReadFromC1(operational, operationalChannel)
	var msg string = ""

	go c1device.DbConnect()

	go func() {
		for b := range operationalChannel {
			msg = string(b)
			fmt.Println(msg)
		}
	}()

	time.Sleep(time.Second * 1)
	log.Println("tcp connection:")
	go c1device.TCPControl("192.168.0.10:8080")

	fmt.Scanln()
}
