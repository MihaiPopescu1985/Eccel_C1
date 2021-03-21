package controller

import (
	"net/http"
	"text/template"

	"example.com/c1/model"
	"example.com/c1/util"
)

const stageTwoPage string = "./web/view/stageTwoAccess.html"

type hrPage struct {
	Workers        []model.Worker
	ActiveProjects []model.Project
}

// StageTwoHandler TODO: write about
func StageTwoHandler(writer http.ResponseWriter, request *http.Request) {

	var pageContent hrPage

	if err := request.ParseForm(); err != nil {
		util.Log.Println(err)
	}

	geNoForm := request.FormValue("ge-no")
	roNoForm := request.FormValue("ro-no")
	descrForm := request.FormValue("description")
	startDateForm := request.FormValue("start-date")

	if geNoForm != "" {
		model.Db.AddProject(geNoForm, roNoForm, descrForm, startDateForm)
	}

	pageContent.ActiveProjects = model.Db.RetrieveActiveProjects()
	pageContent.Workers = model.Db.RetrieveAllWorkers()

	templ, err := template.New("stageTwo").ParseFiles(stageTwoPage)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(writer, "stageTwoAccess.html", pageContent)
	if err != nil {
		util.Log.Println(err)
	}
}
