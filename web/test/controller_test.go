package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/c1/web/controller"
)

func TestHomePageController(t *testing.T) {

	const homePage string = "active projects, devices, active workers"

	t.Run("returns active project, devices and active workers", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		controller.HomePageHandler(response, request)

		got := response.Body.String()
		want := homePage

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
