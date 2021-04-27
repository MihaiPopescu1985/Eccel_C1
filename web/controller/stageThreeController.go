package controller

import (
	"html/template"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
)

const stageThreePage string = "./web/view/stage-three-access.html"

// StageThreeHandler - TODO: write about
func StageThreeHandler(writer http.ResponseWriter, request *http.Request) {

	activeWorkdays, err := model.Db.RetrieveActiveWorkdays()
	if err != nil {
		util.Log.Println(err)
		ErrorPageHandler(writer, request)
	} else {
		templ, err := template.New("stageThree").ParseFiles(stageThreePage)

		if err != nil {
			util.Log.Println(err)
			ErrorPageHandler(writer, request)
			return
		}

		err = templ.ExecuteTemplate(writer, "stage-three-access.html", activeWorkdays)
		if err != nil {
			util.Log.Println(err)
			ErrorPageHandler(writer, request)
			return
		}
	}
}
