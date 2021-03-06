package controller

import (
	"io/ioutil"
	"net/http"

	"example.com/c1/util"
)

const homePage string = "./web/view/home.html"

// HomePageHandler - home page of web interface
// Checks for name and password, redirect user to specific pages.
func HomePageHandler(response http.ResponseWriter, request *http.Request) {

	fileContent, err := ioutil.ReadFile(homePage)
	if err != nil {
		util.Log.Panicln(err)
	}

	response.Write(fileContent)
}
