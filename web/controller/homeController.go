package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"example.com/c1/service"
)

const homePage string = "/home/mihai/Documents/C1/project/web/view/home.html"

// HomePageHandler - home page of web interface
// Checks for name and password, redirect user to specific pages.
func HomePageHandler(response http.ResponseWriter, request *http.Request) {
	fileContent, err := ioutil.ReadFile(homePage)
	if err != nil {
		fmt.Println("Error opening homepage file.")
		fmt.Println(err)
	}

	// TODO: check for errors
	request.ParseForm()
	name := request.FormValue("name")

	password := request.FormValue("password")
	var urlRedirect string

	if name != "" || password != "" {
		var dao service.DAO
		dao.Connect()
		defer dao.CloseConnection()
		worker := dao.GetUserByNameAndPassword(name, password)

		switch worker.AccessLevel {
		case 1:
			urlRedirect = "/stage-one?workerId=" + strconv.Itoa(worker.ID)

		case 2:
			urlRedirect = "/stage-two"

		case 3:
			urlRedirect = "/stage-three"
		}

		fmt.Println(worker.ID)
	}
	if urlRedirect != "" {
		http.Redirect(response, request, urlRedirect, 300)
	} else {
		// TODO: check for errors when write file content
		response.Write(fileContent)
	}
}
