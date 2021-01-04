package c1device

import (
	"encoding/json"

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
