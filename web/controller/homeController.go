package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"example.com/c1/service"
)

const homePage string = "./web/view/home.html"

// HomePageHandler - home page of web interface
// Checks for name and password, redirect user to specific pages.
func HomePageHandler(response http.ResponseWriter, request *http.Request) {

	var dao service.DAO
	dao.Connect()
	defer dao.CloseConnection()

	request.ParseForm()

	name := request.FormValue("name")
	password := request.FormValue("password")

	if name != "" || password != "" {

		worker := dao.GetUserByNameAndPassword(name, password)
		dao.SaveWebToken(worker.ID)
		token = dao.GetActiveToken(worker.ID)

		request.AddCookie(&http.Cookie{
			Name:   "date",
			Value:  token.Date,
			MaxAge: 0,
			Secure: true,
		})
		request.AddCookie(&http.Cookie{
			Name:   "value",
			Value:  token.Token,
			MaxAge: 0,
			Secure: true,
		})
		request.AddCookie(&http.Cookie{
			Name:   "user",
			Value:  strconv.Itoa(token.WorkerID),
			MaxAge: 0,
			Secure: true,
		})

		switch worker.AccessLevel {
		case 1:
			urlRedirect = "/stage-one?workerId=" + strconv.Itoa(worker.ID)

		case 2:
			urlRedirect = "/stage-two"

		case 3:
			urlRedirect = "/stage-three"
		}
	}

	fileContent, err := ioutil.ReadFile(homePage)
	if err != nil {
		fmt.Println("Error opening homepage file.")
		fmt.Println(err)
	}

	response.Write(fileContent)
}
