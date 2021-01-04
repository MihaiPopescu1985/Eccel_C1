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
		var deviceMessage c1device.DeviceMessage

		for msg := range device.WsChannel {

			deviceMessage = device.ParseMessage(msg)
			command :=
				dao.InsertIntoWorkday(deviceMessage.DeviceName, deviceMessage.CardUID)

			if dao.IsConnected() {
				fmt.Println(command)
				dao.Execute(command)
			}
			device.WsRead()
		}
	}()

	fmt.Scanln()
}
