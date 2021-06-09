package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
)

// contains message in json format as the device is reading
type deviceMsg struct {
	Model      string `json:"type"`
	Uid        string `json:"uid"`
	Sak        int    `json:"sak"`
	CardType   string `json:"string"`
	DeviceName string `json:"device_name"`
	TagKnown   bool   `json:"known_tag"`
}

// SaveTimeHandler handles requests from devices sent in json format.
// Requests are parsed and data is saved in database.
func SaveTimeHandler(w http.ResponseWriter, r *http.Request) {

	devName, tagUid, err := parseDeviceReading(r)
	if err != nil {
		util.Log.Println(err)
		return
	}
	model.Db.InsertIntoWorkday(devName, tagUid)
}

func parseDeviceReading(r *http.Request) (string, string, error) {
	body := make([]byte, 1024)
	bytes, err := r.Body.Read(body)
	if err != nil && err != io.EOF {
		return "", "", err
	}
	body = body[:bytes]
	var message deviceMsg

	if err := json.Unmarshal(body, &message); err != nil {
		return "", "", err
	}
	return message.DeviceName, message.Uid, nil
}
