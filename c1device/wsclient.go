package c1device

import (
	"example.com/c1/util"
	"golang.org/x/net/websocket"
)

const wsPrefix string = "ws://"
const wsSufix string = "/wscomm.cgi"
const port string = "8080"
const origin string = "http://localhost/"

// WsRead reads data sent by a C1 device via websocket.
// Data is passed to device's []byte channel.
// A connection to device must be established first.
func (dev *C1Device) WsRead() {

	msg := make([]byte, 1024)
	lenght, err := dev.WsConnection.Read(msg)

	if err != nil {
		util.Log.Panicln(err)
	}
	msg = msg[:lenght]
	go func() {
		dev.WsChannel <- msg
	}()
}

// WsConnect connects to a C1 device via websocket.
func (dev *C1Device) WsConnect() {

	url := wsPrefix + dev.IP + wsSufix
	var err error

	dev.WsConnection, err = websocket.Dial(url, "", origin)
	if err != nil {
		util.Log.Panicln(err)
	}
}
