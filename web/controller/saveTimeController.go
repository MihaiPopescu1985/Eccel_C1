package controller

import (
	"encoding/json"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
)

// contains message in json format as the device is reading
type deviceMsg struct {
	Model      string `json:"type"`
	Uid        string `json:"uid"`
	Sak        string `json:"sak"`
	CardType   string `json:"string"`
	DeviceName string `json:"device_name"`
	TagKnown   bool   `json:"known_tag"`
}

func SaveTimeHandler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, 1)
	r.Body.Read(body)
	util.Log.Println(string(body))

	var message deviceMsg

	if err := json.Unmarshal(body, &message); err != nil {
		util.Log.Println(err)
	}
	model.Db.InsertIntoWorkday(message.DeviceName, message.Uid)
}
