package controller

import (
	"io/ioutil"
	"net/http"
)

//const homePage string = "/home/mihai/Documents/C1/project/web/view/home.html"
const homePage string = "/home/mihai/Documents/C1/project/web/view/directorPage.html"

// HomePageHandler - home page of web interface
func HomePageHandler(writer http.ResponseWriter, request *http.Request) {
	page, _ := ioutil.ReadFile(homePage)

	writer.Write(page)
}
