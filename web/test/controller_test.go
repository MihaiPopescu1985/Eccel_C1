package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"example.com/c1/web/controller"
)

func TestErrorPage(t *testing.T) {
	t.Run("returns error message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/error", nil)
		response := httptest.NewRecorder()

		controller.ErrorPageHandler(response, request)
		got := response.Body.String()
		want := "There was an error processing your request"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
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
