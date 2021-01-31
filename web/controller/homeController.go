package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const homePage string = "/home/mihai/Documents/C1/project/web/view/home.html"

// HomePageHandler - home page of web interface
func HomePageHandler(writer http.ResponseWriter, request *http.Request) {
	fileContent, err := ioutil.ReadFile(homePage)
	if err != nil {
		fmt.Println("Error opening homepage file.")
		fmt.Println(err)
	}
	// TODO: check for errors when write file content
	writer.Write(fileContent)
}
