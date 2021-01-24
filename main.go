package main

import (
	"fmt"

	"example.com/c1/c1device"
	"example.com/c1/service"
)

func main() {
	fmt.Println("Hello world!")

	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("No database connection")
	}

	device1 := c1device.C1Device{
		IP:        "192.168.0.10",
		WsChannel: make(chan []byte),
	}
	device1.WsConnect()
	device1.WsRead()

	go func() {
		for msg := range device1.WsChannel {

			deviceName, cardUID := device1.ParseMessage(msg)

			command :=
				dao.InsertIntoWorkday(deviceName, cardUID)

			if dao.IsConnected() && deviceName != "" && cardUID != "" {
				fmt.Println(command)
				dao.Execute(command)
			}
			device1.WsRead()
		}
	}()

	device2 := c1device.C1Device{
		IP:        "192.168.0.20",
		WsChannel: make(chan []byte),
	}
	device2.WsConnect()
	device2.WsRead()

	go func() {
		for msg := range device2.WsChannel {

			deviceName, cardUID := device2.ParseMessage(msg)

			command :=
				dao.InsertIntoWorkday(deviceName, cardUID)

			if dao.IsConnected() && deviceName != "" && cardUID != "" {
				fmt.Println(command)
				dao.Execute(command)
			}
			device2.WsRead()
		}
	}()

	device3 := c1device.C1Device{
		IP:        "192.168.0.30",
		WsChannel: make(chan []byte),
	}
	device3.WsConnect()
	device3.WsRead()

	go func() {
		for msg := range device3.WsChannel {

			deviceName, cardUID := device3.ParseMessage(msg)

			command :=
				dao.InsertIntoWorkday(deviceName, cardUID)

			if dao.IsConnected() && deviceName != "" && cardUID != "" {
				fmt.Println(command)
				dao.Execute(command)
			}
			device3.WsRead()
		}
	}()

	device4 := c1device.C1Device{
		IP:        "192.168.0.40",
		WsChannel: make(chan []byte),
	}
	device4.WsConnect()
	device4.WsRead()

	go func() {
		for msg := range device4.WsChannel {

			deviceName, cardUID := device4.ParseMessage(msg)

			command :=
				dao.InsertIntoWorkday(deviceName, cardUID)

			if dao.IsConnected() && deviceName != "" && cardUID != "" {
				fmt.Println(command)
				dao.Execute(command)
			}
			device4.WsRead()
		}
	}()

	fmt.Scanln()
}
