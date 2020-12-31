package c1device

import (
	"fmt"
	"testing"
	"time"

	"example.com/c1/c1device"
)

//go test -run Method

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

// go test -run ReceivingAllMessagesViaWebsocket
func TestReceivingAllMessagesViaWebsocket(t *testing.T) {

	device := c1device.C1Device{
		IP:        "192.168.0.10",
		WsChannel: make(chan []byte),
	}

	device.WsConnect()
	device.WsRead()

	go func() {
		fmt.Println("Scan card")

		for msg := range device.WsChannel {
			fmt.Println(string(msg))
			device.WsRead()
		}
	}()

	time.Sleep(time.Second * 10)
}
