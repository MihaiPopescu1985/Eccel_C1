package main

import (
	"fmt"

	"example.com/c1/c1device"
	"example.com/c1/service"
)

func main() {
	fmt.Println("Hello world!")

	go func() {
		var dao service.DAO
		dao.Connect()

		if !dao.IsConnected() {
			fmt.Println("No database connection")
		}

		device := c1device.C1Device{
			IP:        "192.168.0.10",
			WsChannel: make(chan []byte),
		}
		device.WsConnect()
		device.WsRead()

		for msg := range device.WsChannel {

			deviceName, cardUID := device.ParseMessage(msg)

			command :=
				dao.InsertIntoWorkday(deviceName, cardUID)

			if dao.IsConnected() && deviceName != "" && cardUID != "" {
				fmt.Println(command)
				dao.Execute(command)
			}
			device.WsRead()
		}
	}()

	fmt.Scanln()
}
