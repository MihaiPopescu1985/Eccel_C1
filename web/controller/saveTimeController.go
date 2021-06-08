package controller

import (
	"net/http"

	"example.com/c1/util"
)

// contains message in json format as the device is reading
// type deviceMsg struct {
// }

func SaveTimeHandler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, 0)
	r.Body.Read(body)
	util.Log.Println(string(body))

}
