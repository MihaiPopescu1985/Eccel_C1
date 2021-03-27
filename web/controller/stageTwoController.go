package controller

import (
	"html/template"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
)

const (
	stageTwoPage    string = "./web/view/stageTwoAccess.html"
	editProjectPage string = "./web/view/editProject.html"
)

var actions []string = []string{"editProject", "editWorker", "addProject", "addWorker"}

type hrPage struct {
	Workers        []model.Worker
	ActiveProjects []model.Project
	Positions      map[int]string
}

// StageTwoHandler TODO: write about
func StageTwoHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		util.Log.Println(err)
	}

	switch r.FormValue("action") {

	case "addProject":
		addProject(&w, r)

	case "editProject":
		editProject(&w, r)

	case "saveProject":
		saveProject(&w, r)

	case "addWorker":
		addWorker(&w, r)

	case "editWorker":
		editWorker(&w, r)

	default:
		showMainWindow(&w, r)
	}
}

func editWorker(w *http.ResponseWriter, r *http.Request) {
}

func saveProject(w *http.ResponseWriter, r *http.Request) {
	model.Db.UpdateProject(model.Project{
		ID:          r.FormValue("id"),
		GeNumber:    r.FormValue("ge-no"),
		RoNumber:    r.FormValue("ro-no"),
		Description: r.FormValue("descr"),
		IPAddress:   r.FormValue("ip"),
		DeviceID:    r.FormValue("dev-id"),
		IsActive: func() bool {
			if r.FormValue("active") == "true" {
				return true
			}
			return false
		}(),
		Begin: r.FormValue("begin"),
		End:   r.FormValue("end"),
	})

	http.Redirect(*w, r, "/", 302)
}

func editProject(w *http.ResponseWriter, r *http.Request) {
	project := model.Db.GetProject(r.FormValue("project"))

	templ, err := template.New("editProject").ParseFiles(editProjectPage)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(*w, "editProject.html", project)
	if err != nil {
		util.Log.Println(err)
	}
}

func addWorker(w *http.ResponseWriter, r *http.Request) {
	worker := model.Worker{
		FirstName:   r.FormValue("first-name"),
		LastName:    r.FormValue("last-name"),
		CardNumber:  r.FormValue("card-number"),
		Position:    r.FormValue("positions"),
		Nickname:    r.FormValue("nickname"),
		Password:    r.FormValue("password"),
		AccessLevel: 1,
	}

	if worker.FirstName != "" && worker.LastName != "" &&
		worker.CardNumber != "" && worker.Position != "" &&
		worker.Nickname != "" && worker.Password != "" {

		model.Db.AddWorker(worker)

		http.Redirect(*w, r, "/", 302)
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
	http.Redirect(*w, r, "/", 302)
}

func showMainWindow(w *http.ResponseWriter, r *http.Request) {
	var pageContent hrPage

	pageContent.ActiveProjects = model.Db.RetrieveActiveProjects()
	pageContent.Workers = model.Db.RetrieveAllWorkers()
	pageContent.Positions = model.Db.RetrieveAllPositions()

	templ, err := template.New("stageTwo").ParseFiles(stageTwoPage)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(*w, "stageTwoAccess.html", pageContent)
	if err != nil {
		util.Log.Println(err)
	}
}
