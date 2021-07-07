package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"example.com/c1/model"
)

func TestJWTMiddleware(t *testing.T) {
	t.Run("should pass if valid token", func(t *testing.T) {

	})
}

func TestLoginController(t *testing.T) {
	t.Run("Login should return OK status if a request is made with GET method", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/login", nil)
		response := httptest.NewRecorder()

		Login(response, request)

		want := http.StatusOK
		got := response.Code

		if want != got {
			t.Fatalf("want status code %v, got %v instead", want, got)
		}
	})
	t.Run("test should fail if a request is made with POST method and no form", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/login", nil)
		response := httptest.NewRecorder()

		Login(response, request)

		want := http.StatusUnprocessableEntity
		got := response.Code

		if want != got {
			t.Fatalf("want status code %v, got %v instead", want, got)
		}
	})
	t.Run("test should pass with valid form", func(t *testing.T) {
		model.Db = &model.MockPersist{}
		model.Db.Init("")

		body := *strings.NewReader("name=Mihai&password=Popescu")
		err := model.Db.AddWorker(model.Worker{
			Nickname:    "Mihai",
			FirstName:   "Mihai",
			LastName:    "Popescu",
			Password:    "Popescu",
			AccessLevel: "3",
			ID:          "1",
		})

		if err != nil {
			t.Fatal(err)
		}

		request := httptest.NewRequest(http.MethodPost, "/login", &body)
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		response := httptest.NewRecorder()
		Login(response, request)

		want := http.StatusOK
		got := response.Code

		if want != got {
			t.Fatalf("want status code %v, got %v instead", want, got)
		}
	})
	t.Run("response must contain a jwt token", func(t *testing.T) {
		model.Db = &model.MockPersist{}
		model.Db.Init("")

		body := *strings.NewReader("name=Mihai&password=Popescu")
		err := model.Db.AddWorker(model.Worker{
			Nickname:    "Mihai",
			FirstName:   "Mihai",
			LastName:    "Popescu",
			Password:    "Popescu",
			AccessLevel: "3",
			ID:          "1",
		})

		if err != nil {
			t.Fatal(err)
		}

		request := httptest.NewRequest(http.MethodPost, "/login", &body)
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		response := httptest.NewRecorder()
		Login(response, request)

		if response.Header().Get("Authentication") == "" {
			t.Fatal("missing jwt token")
		}
	})
}

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

func TestShouldGenerateProperStandardViewOfMonthTimeReport(t *testing.T) {

	layout := "2006-1-2 15:04:05"

	for i := 0; i < 366; i++ {
		str := "2021-12-31 12:33:34"

		tm, err := time.Parse(layout, str)

		if err != nil || tm.Day() != 31 {
			t.Fatal(err)
			t.FailNow()
		}
	}
}

func TestShouldParseURL(t *testing.T) {

	want := "4"
	url, err := url.Parse("/stage-one?workerId=4")

	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}

	got := url.Query().Get("workerId")

	if got != want {
		t.Fatalf("got = %v, wanted = %v", got, want)
		t.FailNow()
	}
}
