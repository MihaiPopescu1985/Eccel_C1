package controller

import (
	"fmt"
	"net/http"
)

// HomePageHandler - home page of web interface
func HomePageHandler(httpWriter http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprint(httpWriter, "active projects, devices, active workers")
}
