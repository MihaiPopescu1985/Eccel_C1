package c1device

import (
	"fmt"
	"testing"
	"time"

	"example.com/c1/c1device"
)

/*
go test -run Method
*/

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
/*
Example of card reading
{
        "type": "uid",
        "uid":  "045D91B22C5E80",
        "sak":  0,
        "string":       "MIFARE Ultralight",
        "device-name":  "Pepper_C1-1A6318",
        "memory":       "045D9140B22C5E8040480000E110120005060708090A0B0C0000000000000000",
        "known_tag":    true
}
*/
func TestReceivingAllMessagesViaWebsocketInCertainDuration(t *testing.T) {
	const testDuration = 30
	const IP = "192.168.0.10"

	device := c1device.C1Device{
		IP:        IP,
		WsChannel: make(chan []byte),
	}

	device.WsConnect()
	device.WsRead()

	if device.WsChannel == nil {
		t.Error("Device channel is nil")
		t.FailNow()
	}

	go func() {
		fmt.Println("Scan card")

		for msg := range device.WsChannel {
			fmt.Println(string(msg))
			device.WsRead()
		}
	}()

	time.Sleep(time.Second * testDuration)
}

func TestParseMessageMethod(t *testing.T) {

	jsonMessage := []byte{123, 10, 9, 34, 116, 121, 112, 101, 34,
		58, 9, 34, 117, 105, 100, 34, 44, 10, 9, 34, 117, 105,
		100, 34, 58, 9, 34, 48, 52, 53, 68, 57, 49, 66, 50,
		50, 67, 53, 69, 56, 48, 34, 44, 10, 9, 34, 115, 97,
		107, 34, 58, 9, 48, 44, 10, 9, 34, 115, 116, 114, 105,
		110, 103, 34, 58, 9, 34, 77, 73, 70, 65, 82, 69, 32, 85,
		108, 116, 114, 97, 108, 105, 103, 104, 116, 34, 44, 10,
		9, 34, 100, 101, 118, 105, 99, 101, 45, 110, 97, 109, 101,
		34, 58, 9, 34, 80, 101, 112, 112, 101, 114, 95, 67, 49,
		45, 49, 65, 54, 51, 49, 56, 34, 44, 10, 9, 34, 109, 101,
		109, 111, 114, 121, 34, 58, 9, 34, 48, 52, 53, 68, 57, 49,
		52, 48, 66, 50, 50, 67, 53, 69, 56, 48, 52, 48, 52, 56, 48,
		48, 48, 48, 69, 49, 49, 48, 49, 50, 48, 48, 48, 53, 48, 54,
		48, 55, 48, 56, 48, 57, 48, 65, 48, 66, 48, 67, 48, 48, 48,
		48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 34, 44,
		10, 9, 34, 107, 110, 111, 119, 110, 95, 116, 97, 103, 34, 58,
		9, 116, 114, 117, 101, 10, 125}

	expectedDeviceName := "Pepper_C1-1A6318"
	expectedCardUID := "045D91B22C5E80"

	device := c1device.C1Device{}

	gotDeviceMessage := device.ParseMessage(jsonMessage)

	if expectedDeviceName != gotDeviceMessage.DeviceName {
		t.Error("got device name = " + gotDeviceMessage.DeviceName)
		t.Error("got device name != expected device name")
		t.FailNow()
	}
	if expectedCardUID != gotDeviceMessage.CardUID {
		t.Error("Got card UID != expected card UID")
		t.FailNow()
	}
}
