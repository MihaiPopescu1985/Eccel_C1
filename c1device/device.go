package c1device

import (
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

// DeviceMessage stores data from C1 device
type DeviceMessage struct {
	CardUID    string
	DeviceName string
}

// ParseMessage returns device name and rfid card's uid when reading.
func (device *C1Device) ParseMessage(message []byte) DeviceMessage {

	const (
		deviceNameStartIndex = 27
		deviceNameEndIndex   = 41
		cardUIDStartIndex    = 104
		cardUIDEndIndex      = 120
	)
	return DeviceMessage{
		CardUID:    string(message[deviceNameStartIndex:deviceNameEndIndex]),
		DeviceName: string(message[cardUIDStartIndex:cardUIDEndIndex]),
	}
}
