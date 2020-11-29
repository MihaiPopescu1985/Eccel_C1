package c1device

import (
	"log"

	"golang.org/x/net/websocket"
)

// WsRead connects to a given C1 device via websocket
// and sends the data through device's channel
func (dev C1Device) WsRead() {

	url := wsPrefix + dev.IP + wsSufix
	con, error := websocket.Dial(url, "ws", localhost)
	if error != nil {
		log.Println(error)
	}
	go func() {
		defer close(dev.WsChannel)
		defer con.Close()

		for {
			msg := make([]byte, 1024)
			lenght, err := con.Read(msg)
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
}
