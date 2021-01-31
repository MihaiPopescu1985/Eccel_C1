package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"example.com/c1/service"
)

const stageThreePage string = "/home/mihai/Documents/C1/project/web/view/stageThreeAccess.html"

// StageThreeHandler - TODO: write about
func StageThreeHandler(writer http.ResponseWriter, request *http.Request) {

	var dao service.DAO
	dao.Connect()
	defer dao.CloseConnection()

	activeWorkdays := dao.RetrieveActiveWorkdays(dao.ExecuteQuery(dao.SelectActiveWorkday()))

	templ, err := template.New("stageThree").ParseFiles(stageThreePage)
	if err != nil {
		panic(err)
	}

	err = templ.ExecuteTemplate(writer, "stageThreeAccess.html", activeWorkdays)
	if err != nil {
		fmt.Println(err)
	}
}
