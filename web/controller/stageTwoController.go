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
	Positions      map[int]string
}

// StageTwoHandler TODO: write about
func StageTwoHandler(writer http.ResponseWriter, request *http.Request) {

	var pageContent hrPage

	parsingForms(&writer, request)

	pageContent.ActiveProjects = model.Db.RetrieveActiveProjects()
	pageContent.Workers = model.Db.RetrieveAllWorkers()
	pageContent.Positions = model.Db.RetrieveAllPositions()

	templ, err := template.New("stageTwo").ParseFiles(stageTwoPage)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(writer, "stageTwoAccess.html", pageContent)
	if err != nil {
		util.Log.Println(err)
	}
}

func parsingForms(w *http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		util.Log.Println(err)
	}

	geNoForm := r.FormValue("ge-no")
	roNoForm := r.FormValue("ro-no")
	descrForm := r.FormValue("description")
	startDateForm := r.FormValue("start-date")

	if geNoForm != "" && roNoForm != "" && descrForm != "" && startDateForm != "" {
		model.Db.AddProject(geNoForm, roNoForm, descrForm, startDateForm)

		r.Form.Del("ge-no")
		r.Form.Del("ro-no")
		r.Form.Del("description")
		r.Form.Del("start-date")

		http.Redirect(*w, r, "/", 302)
	}

	firstName := r.FormValue("first-name")
	lastName := r.FormValue("last-name")
	cardNumber := r.FormValue("card-number")
	position := r.FormValue("positions")
	nickName := r.FormValue("nickname")
	password := r.FormValue("password")

	if firstName != "" && lastName != "" && cardNumber != "" && position != "" && nickName != "" && password != "" {
		model.Db.AddWorker(firstName, lastName, cardNumber, position, nickName, password)

		r.Form.Del("first-name")
		r.Form.Del("last-name")
		r.Form.Del("card-number")
		r.Form.Del("positions")
		r.Form.Del("nickname")
		r.Form.Del("password")

		http.Redirect(*w, r, "/", 302)
	}
}
