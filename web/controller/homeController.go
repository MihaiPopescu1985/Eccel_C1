package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const homePage string = "./web/view/home.html"

// HomePageHandler - home page of web interface
// Checks for name and password, redirect user to specific pages.
func HomePageHandler(response http.ResponseWriter, request *http.Request) {

	fileContent, err := ioutil.ReadFile(homePage)
	if err != nil {
		fmt.Println("Error opening homepage file.")
		fmt.Println(err)
	}

	response.Write(fileContent)
}
