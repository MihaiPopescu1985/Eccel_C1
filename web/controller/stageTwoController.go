package controller

import (
	"html/template"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
)

const (
	activeProjectsPage string = "./web/view/stage-two/activeProjects.html"
	editProjectPage    string = "./web/view/editProject.html"
	editWorkerPage     string = "./web/view/editWorker.html"
)

// StageTwoHandler TODO: write about
func StageTwoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.FormValue("action") {
	case "addProject":
		addProject(&w, r)

	default:
		showActiveProjects(w, r)
	}
}

func addProject(w *http.ResponseWriter, r *http.Request) {
	project := model.Project{
		GeNumber:    r.FormValue("ge-no"),
		RoNumber:    r.FormValue("ro-no"),
		Description: r.FormValue("description"),
		Begin:       r.FormValue("start-date"),
	}

	if project.GeNumber != "" && project.RoNumber != "" &&
		project.Description != "" && project.Begin != "" {

		model.Db.AddProject(project)
	}
	http.Redirect(*w, r, "/", http.StatusFound)
}

func showActiveProjects(w http.ResponseWriter, r *http.Request) {

	activeProjects := model.Db.RetrieveActiveProjects()

	templ, err := template.New("stageTwo").ParseFiles(activeProjectsPage)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(w, "activeProjects.html", activeProjects)
	if err != nil {
		util.Log.Println(err)
	}
}
