package c1device

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

const wsPrefix string = "ws://"
const wsSufix string = "/wscomm.cgi"
const port string = "8080"
const origin string = "http://localhost/"

// WsRead connects to a given C1 device via websocket
// and sends the data through device's channel
// The message is a slice of bytes encoding json.
// The json has the fields:
// - device-name
// - known_tag
// - memory
// - sak
// - string
// - type
// - uid
func (dev *C1Device) WsRead() {
	msg := make([]byte, 1024)
	lenght, err := dev.WsConnection.Read(msg)
	if err != nil {
		fmt.Println(err)
	}
	msg = msg[:lenght]
	go func() {
		dev.WsChannel <- msg
	}()

	/*
		go func() {
			defer close(dev.WsChannel)
			defer dev.WsConnection.Close()

			for {
				msg := make([]byte, 1024)
				lenght, err := dev.WsConnection.Read(msg)
				if err != nil {
					log.Println(err)
					break
				}
				if lenght != 0 {
					msg = msg[:lenght]
					dev.WsChannel <- msg
				}
			}
		}()
	*/
}

// WsConnect connects to a C1 device via websocket.
func (dev *C1Device) WsConnect() {
	url := wsPrefix + dev.IP + wsSufix
	var error error
	dev.WsConnection, error = websocket.Dial(url, "", origin)
	if error != nil {
		log.Println("Error while connecting to " + dev.Name)
		log.Println(error)
	}
}
