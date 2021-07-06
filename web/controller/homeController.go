package controller

import (
	"io/ioutil"
	"log"
	"net/http"
)

const homePage string = "./web/view/index.html"

// HomePageHandler - home page of web interface
// Checks for name and password, redirect user to specific pages.
func HomePageHandler(response http.ResponseWriter, request *http.Request) {

	fileContent, err := ioutil.ReadFile(homePage)
	if err != nil {
		log.Panicln(err)
	}

	response.Write(fileContent)
}
