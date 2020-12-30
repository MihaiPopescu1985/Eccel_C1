package main

import (
	"fmt"
)

func main() {

	msg := []byte("{ \"frame\":	\"version\", \"string\": \" 1.6 Oct 15 2020 14:19:44 \" }")
	fmt.Println(msg)
	/*
		operational := c1device.C1Device{
			Name:      "operational",
			IP:        "192.168.0.10",
			TCPPort:   "8080",
			WsChannel: make(chan []byte),
		}
		operational.WsConnect()
		operational.WsRead()

			go func() {
				for msg := range operational.WsChannel {
					var message interface{}
					err := json.Unmarshal(msg, &message)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(message)
					fmt.Println(time.Now())
				}
			}()

			log.Println("Establishing a TCP connection:")
	*/
	//	go operational.TCPControl()

	/*
		db.Connect()
		go func() {
			for {
				log.Println(db.IsConnected())
				time.Sleep(time.Minute)
			}
		}()
	*/
	fmt.Scanln()
}
