package main

import (
	"fmt"
	"log"
	"time"

	"example.com/c1/dbcommunication"
	"example.com/c1/tcpcommunication"
	"example.com/c1/wscommunication"
)

const operational string = "ws://192.168.0.10/wscomm.cgi"

func main() {
	fmt.Println("muie psd")
	//c1commands.PrintCommands()

	operationalChannel := make(chan []byte)
	wscommunication.ReadFromC1(operational, operationalChannel)
	var msg string = ""

	go dbcommunication.DbConnect()

	go func() {
		for b := range operationalChannel {
			msg = string(b)
			fmt.Println(msg)
		}
	}()

	time.Sleep(time.Second * 1)
	log.Println("tcp connection:")
	go tcpcommunication.TCPControl("192.168.0.10:8080")

	fmt.Scanln()
}
