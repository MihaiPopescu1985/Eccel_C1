package c1device

import (
	"testing"

	"example.com/c1/c1device"
)

const WsConnectionResponse string = "{ \"frame\":	\"version\", \"string\": \" 1.6 Oct 15 2020 14:19:44 \" }"

func TestWebsocketConnection(t *testing.T) {

	device := c1device.C1Device{
		IP: "192.168.0.10",
	}
	device.WsConnect()

	if device.WsConnection == nil {
		t.Error("Nil websocket connection.")
		t.FailNow()
	}
}

func TestReceivingMessageViaWebsocket(t *testing.T) {
	device := c1device.C1Device{
		IP:        "192.168.0.10",
		WsChannel: make(chan []byte),
	}

	device.WsConnect()
	device.WsRead()

	got := <-device.WsChannel

	if len(got) == 0 {
		t.Error("Received no message from device.")
		t.FailNow()
	}
}
