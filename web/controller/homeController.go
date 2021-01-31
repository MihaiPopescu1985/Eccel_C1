package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"example.com/c1/service"
)

//const homePage string = "/home/mihai/Documents/C1/project/web/view/home.html"
const homePage string = "/home/mihai/Documents/C1/project/web/view/directorPage.html"

// HomePageHandler - home page of web interface
func HomePageHandler(writer http.ResponseWriter, request *http.Request) {

	var dao service.DAO
	dao.Connect()
	defer dao.CloseConnection()

	activeWorkdays := dao.RetrieveActiveWorkdays(dao.ExecuteQuery(dao.SelectActiveWorkday()))

	templ, err := template.New("test").ParseFiles(homePage)
	if err != nil {
		panic(err)
	}

	err = templ.ExecuteTemplate(writer, "directorPage.html", activeWorkdays)
	if err != nil {
		fmt.Println(err)
	}
}
