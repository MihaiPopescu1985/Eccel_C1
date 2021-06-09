package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSaveTimeController(t *testing.T) {
	t.Run("should run into an error", func(t *testing.T) {

		message := `{
"type": "uid",
"uid": "045D91B22C5E80",
"sak": 0,
"string": "MIFARE Ultralight",
"device_name": "Pepper_C1-1A631C",
"known_tag": false
`
		r := httptest.NewRequest(http.MethodPost, "/save-time", strings.NewReader(message))

		_, _, err := parseDeviceReading(r)
		if err == nil {
			t.Error("want error but did not got one")
		}
	})

	t.Run("should parse device message", func(t *testing.T) {
		message := `{
"type": "uid",
"uid": "045D91B22C5E80",
"sak": 0,
"string": "MIFARE Ultralight",
"device_name": "Pepper_C1-1A631C",
"known_tag": false
}	
`
		r := httptest.NewRequest(http.MethodPost, "/save-time", strings.NewReader(message))
		wantDevName := "Pepper_C1-1A631C"
		wantTagUid := "045D91B22C5E80"

		gotDevName, gotTagUid, err := parseDeviceReading(r)
		if err != nil {
			t.Error("got unwanted error: ", err)
		}
		if gotDevName != wantDevName {
			t.Errorf("want device name (%q) != got device name (%q)", wantDevName, gotDevName)
		}
		if gotTagUid != wantTagUid {
			t.Errorf("want tag uid (%q) != got tag uid (%q)", wantTagUid, gotTagUid)
		}
	})

}
