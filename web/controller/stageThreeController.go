package controller

import (
	"html/template"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
)

const stageThreePage string = "./web/view/stageThreeAccess.html"

// StageThreeHandler - TODO: write about
func StageThreeHandler(writer http.ResponseWriter, request *http.Request) {

	activeWorkdays := model.Db.RetrieveActiveWorkdays()
	templ, err := template.New("stageThree").ParseFiles(stageThreePage)

	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(writer, "stageThreeAccess.html", activeWorkdays)
	if err != nil {
		util.Log.Println(err)
	}
}
