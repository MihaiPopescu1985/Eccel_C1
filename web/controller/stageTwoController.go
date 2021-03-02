package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"example.com/c1/model"
	"example.com/c1/service"
)

const stageTwoPage string = "./web/view/stageTwoAccess.html"

type hrPage struct {
	Workers        []model.Worker
	ActiveProjects []model.Project
}

// StageTwoHandler TODO: write about
func StageTwoHandler(writer http.ResponseWriter, request *http.Request) {

	var pageContent hrPage

	pageContent.ActiveProjects = service.Dao.RetrieveActiveProjects()
	pageContent.Workers = service.Dao.RetrieveAllWorkers()

	templ, err := template.New("stageTwo").ParseFiles(stageTwoPage)
	if err != nil {
		fmt.Println(err)
	}

	err = templ.ExecuteTemplate(writer, "stageTwoAccess.html", pageContent)
	if err != nil {
		fmt.Println(err)
	}
}
