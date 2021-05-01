package c1device

import (
	"encoding/json"

	"example.com/c1/model"
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
func (device *C1Device) ParseMessage(message []byte) (string, string, error) {

	var cardReading CardReading
	var deviceName string
	var cardUID string

	err := json.Unmarshal(message, &cardReading)

	if err != nil {
		return "", "", err
	} else {
		deviceName = cardReading.DeviceName
		cardUID = cardReading.UID
	}
	return deviceName, cardUID, nil
}

// UseDevice put at use a C1 device.
// The normal use is:
// - initiate a connection to device via websocket,
// - reads cards data
// - calls a database connection to save data into database
func (device *C1Device) UseDevice() error {
	var err error
	err = device.WsConnect()

	if err != nil {
		return err
	}

	err = device.WsRead()
	if err != nil {
		return err
	}

	go func() error {
		for msg := range device.WsChannel {
			deviceName, cardUID, err := device.ParseMessage(msg)
			if err != nil {
				return err
			}

			err = model.Db.IsConnected()
			if err != nil {
				return err
			} else if deviceName != "" && cardUID != "" {
				err = model.Db.InsertIntoWorkday(deviceName, cardUID)
				if err != nil {
					return err
				}
			}
			err = device.WsRead()
			if err != nil {
				return err
			}
		}
		return err
	}()
	return err
}
