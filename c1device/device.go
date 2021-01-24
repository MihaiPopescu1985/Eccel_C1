package c1device

import (
	"encoding/json"
	"fmt"

	"example.com/c1/service"
	"golang.org/x/net/websocket"
)

// C1Device describes a C1 device
type C1Device struct {
	Name    string
	IP      string
	TCPPort string

	WsConnection *websocket.Conn
	WsChannel    chan []byte
}

// CardReading is a structure that resamble the reading from C1 device.
type CardReading struct {
	Type       string `json:"type"`
	UID        string `json:"uid"`
	Sak        int    `json:"sak"`
	CardType   string `json:"string"`
	DeviceName string `json:"device-name"`
	Memory     string `json:"memory"`
	IsTagKnown bool   `json:"known_tag"`
}

// ParseMessage returns device name and rfid card's uid when reading.
func (device *C1Device) ParseMessage(message []byte) (string, string) {

	var cardReading CardReading
	err := json.Unmarshal(message, &cardReading)

	var deviceName string
	var cardUID string

	if err == nil {
		deviceName = cardReading.DeviceName
		cardUID = cardReading.UID
	}
	return deviceName, cardUID
}

// UseDevice put at use a C1 device.
// The normal use is:
// - initiate a connection to device via websocket,
// - reads cards data
// - calls a dao to save data into database
func (device *C1Device) UseDevice(dao service.DAO) {

	device.WsConnect()
	device.WsRead()

	go func() {
		for msg := range device.WsChannel {
			deviceName, cardUID := device.ParseMessage(msg)

			command :=
				dao.InsertIntoWorkday(deviceName, cardUID)

			if dao.IsConnected() && deviceName != "" && cardUID != "" {
				fmt.Println(command) // TODO: must be replaced with a proper logger
				dao.Execute(command)
			}
			device.WsRead()
		}
	}()
}