package c1device

import "golang.org/x/net/websocket"

// C1Device describes a C1 device
type C1Device struct {
	Name    string
	IP      string
	TCPPort string

	WsConnection *websocket.Conn
	WsChannel    chan []byte
}
