package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/c1/model"
)

// func TestParseUri(t *testing.T) {
// 	// rawUrl := "http://192.168.0.109:8080/?action=edit-project&proj-id=16&ge-no=301104010&ro-no=451078001&description=Cabinet+S10&dev-id=4&start-date=2021-04-28&active=true&end=2021-06-09"

// }

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

func TestMiddleware(t *testing.T) {
	// populate database with workers
	model.Db = &model.MockPersist{}
	model.Db.Init("")
	model.Db.AddWorker(model.Worker{
		Nickname:    "validStageOneName",
		Password:    "validStageOnePassword",
		AccessLevel: "1",
	})
	model.Db.AddWorker(model.Worker{
		Nickname:    "validStageTwoName",
		Password:    "validStageTwoPassword",
		AccessLevel: "2",
	})
	model.Db.AddWorker(model.Worker{
		Nickname:    "validStageThreeName",
		Password:    "validStageThreePassword",
		AccessLevel: "3",
	})

	t.Run("permited URLs must pass middleware", func(t *testing.T) {

		for _, url := range permitedURL {
			request := httptest.NewRequest(http.MethodGet, url, nil)
			response := httptest.NewRecorder()

			mid := AuthMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {}))
			mid.ServeHTTP(response, request)

			if response.Code != http.StatusOK {
				t.Fatalf("wrong http code. got %v", response.Code)
			}
		}
	})

	t.Run("unlogged user must not gain access to forbidden urls", func(t *testing.T) {
		requests := []*http.Request{
			httptest.NewRequest(http.MethodGet, "/stage-one", nil),
			httptest.NewRequest(http.MethodGet, "/stage-two", nil),
			httptest.NewRequest(http.MethodGet, "/stage-three", nil),
			httptest.NewRequest(http.MethodGet, "/test-url", nil),
		}

		for _, r := range requests {
			response := httptest.NewRecorder()

			mid := AuthMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {}))
			mid.ServeHTTP(response, r)

			if response.Code != http.StatusForbidden {
				t.Fatalf("wrong code, wanted %v, got %v on url %v", http.StatusForbidden, response.Code, r.RequestURI)
			}
		}
	})

	t.Run("test access levels by submitting a form with credentials",
		func(t *testing.T) {
			// mocking a submit form containing worker's name and password and access level that we expect to be granted
			testForms := []struct {
				givenForm strings.Reader
				wantMatch string
			}{
				{
					*strings.NewReader("name=validStageOneName&password=validStageOnePassword"),
					"1",
				},
				{
					*strings.NewReader("name=validStageTwoName&password=validStageTwoPassword"),
					"2",
				},
				{
					*strings.NewReader("name=validStageThreeName&password=validStageThreePassword"),
					"3",
				},
			}

			// test
			for key, value := range accessStages {
				for _, v := range testForms {
					testStage(t, key, value, v.givenForm, v.wantMatch)
				}
			}
		})

	t.Run("test access levels by submitting cookies", func(t *testing.T) {
		tests := []struct {
			request      *http.Request
			cookies      []http.Cookie
			expectedCode int
		}{
			{
				httptest.NewRequest(http.MethodGet, accessStages["1"], nil),
				[]http.Cookie{{Name: "name", Value: "validStageOneName"}, {Name: "pass", Value: "validStageOnePassword"}},
				http.StatusOK,
			},
			{
				httptest.NewRequest(http.MethodGet, accessStages["2"], nil),
				[]http.Cookie{{Name: "name", Value: "validStageOneName"}, {Name: "pass", Value: "validStageOnePassword"}},
				http.StatusForbidden,
			},
			{
				httptest.NewRequest(http.MethodGet, accessStages["3"], nil),
				[]http.Cookie{{Name: "name", Value: "validStageOneName"}, {Name: "pass", Value: "validStageOnePassword"}},
				http.StatusForbidden,
			},
			{
				httptest.NewRequest(http.MethodGet, accessStages["1"], nil),
				[]http.Cookie{{Name: "name", Value: "validStageTwoName"}, {Name: "pass", Value: "validStageTwoPassword"}},
				http.StatusForbidden,
			},
			{
				httptest.NewRequest(http.MethodGet, accessStages["2"], nil),
				[]http.Cookie{{Name: "name", Value: "validStageTwoName"}, {Name: "pass", Value: "validStageTwoPassword"}},
				http.StatusOK,
			},
			{
				httptest.NewRequest(http.MethodGet, accessStages["3"], nil),
				[]http.Cookie{{Name: "name", Value: "validStageTwoName"}, {Name: "pass", Value: "validStageTwoPassword"}},
				http.StatusForbidden,
			},
			{
				httptest.NewRequest(http.MethodGet, accessStages["1"], nil),
				[]http.Cookie{{Name: "name", Value: "validStageThreeName"}, {Name: "pass", Value: "validStageThreePassword"}},
				http.StatusForbidden,
			},
			{
				httptest.NewRequest(http.MethodGet, accessStages["2"], nil),
				[]http.Cookie{{Name: "name", Value: "validStageThreeName"}, {Name: "pass", Value: "validStageThreePassword"}},
				http.StatusForbidden,
			},
			{
				httptest.NewRequest(http.MethodGet, accessStages["3"], nil),
				[]http.Cookie{{Name: "name", Value: "validStageThreeName"}, {Name: "pass", Value: "validStageThreePassword"}},
				http.StatusOK,
			},
		}

		for _, v := range tests {
			for _, cookie := range v.cookies {
				v.request.AddCookie(&cookie)
			}
			response := httptest.NewRecorder()

			mid := AuthMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {}))
			mid.ServeHTTP(response, v.request)

			if response.Code != v.expectedCode {
				t.Fatalf("got code %v, expected %v", response.Code, v.expectedCode)
			}
		}
	})
}

func testStage(t *testing.T, key, path string, body strings.Reader, match string) {
	t.Helper()

	request := httptest.NewRequest(http.MethodPost, path, &body)
	request.RequestURI = path
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()

	mid := AuthMiddleware(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {}))
	mid.ServeHTTP(response, request)

	switch key {
	case match:
		if response.Code != http.StatusOK {
			t.Fatalf("want response code %v but got %v on url: %v, given body: %v", http.StatusOK, response.Code, path, body)
		}
	default:
		if response.Code != http.StatusForbidden {
			t.Fatalf("want response code %v but got %v on url: %v, given body: %v", http.StatusForbidden, response.Code, path, body)
		}
	}
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
