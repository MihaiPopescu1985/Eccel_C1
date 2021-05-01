package c1device

import (
	"golang.org/x/net/websocket"
)

const wsPrefix string = "ws://"
const wsSufix string = "/wscomm.cgi"
const port string = "8080"
const origin string = "http://localhost/"

// WsRead reads data sent by a C1 device via websocket.
// Data is passed to device's []byte channel.
// A connection to device must be established first.
func (dev *C1Device) WsRead() error {

	msg := make([]byte, 1024)
	lenght, err := dev.WsConnection.Read(msg)

	if err != nil {
		return err
	}
	msg = msg[:lenght]
	go func() {
		dev.WsChannel <- msg
	}()
	return nil
}

// WsConnect connects to a C1 device via websocket.
func (dev *C1Device) WsConnect() error {

	url := wsPrefix + dev.IP + wsSufix
	var err error

	dev.WsConnection, err = websocket.Dial(url, "", origin)
	return err
}
