package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"example.com/c1/model"
	"example.com/c1/util"
)

const (
	activeProjectsPage   string = "./web/view/stage-two/active-projects.html"
	sentProjectsPage     string = "./web/view/stage-two/sent-projects.html"
	stageTwofreeDaysPage string = "./web/view/stage-two/free-days.html"

	/* Not used. The edit project and edit worker are modals */
	//editProjectPage    string = "./web/view/editProject.html"
	//editWorkerPage     string = "./web/view/editWorker.html"
)

// StageTwoHandler TODO: write about
func StageTwoHandler(w http.ResponseWriter, r *http.Request) {

	switch parseURI(r, "view") {
	case "add-project":
		addProject(&w, r)

	case "edit-project":
		editProject(&w, r)

	case "sent-projects":
		sentProjectsView(&w, r)

	case "free-days":
		pageContent := model.Db.RetrieveFreeDays()
		serveFreeDaysPage(&w, r, pageContent)

	case "delete-free-day":
		freeDayID := parseURI(r, "free-day")
		model.Db.DeleteFreeDay(freeDayID)
		http.Redirect(w, r, "/?view=free-days", http.StatusFound)

	case "add-free-day":
		freeDay := parseURI(r, "free-day")
		model.Db.AddFreeDay(freeDay)
		http.Redirect(w, r, "/?view=free-days", http.StatusFound)
	default:
		showActiveProjects(w, r)
	}
}

func serveFreeDaysPage(w *http.ResponseWriter, r *http.Request, pageContent []string) {

	templ, err := template.New("freeDays").ParseFiles(stageTwofreeDaysPage)
	if err != nil {
		util.Log.Println(err)
	}
	if err = templ.ExecuteTemplate(*w, "free-days.html", pageContent); err != nil {
		util.Log.Println(err)
	}
}

func sentProjectsView(w *http.ResponseWriter, r *http.Request) {

	pageContent := model.Db.RetrieveSentProjects()
	for k, v := range pageContent {
		pageContent[k] = toHoursAndMinutes(v)
	}

	templ, err := template.New("sentProjects").ParseFiles(sentProjectsPage)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(*w, "sent-projects.html", pageContent)
	if err != nil {
		util.Log.Println(err)
	}
}

func editProject(w *http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		util.Log.Panicln(err)
	}

	project := model.Project{
		ID:          r.FormValue("id"),
		IPAddress:   r.FormValue("ip"),
		GeNumber:    r.FormValue("ge-no"),
		RoNumber:    r.FormValue("ro-no"),
		Description: r.FormValue("descr"),
		DeviceID:    r.FormValue("dev-id"),
		Begin:       r.FormValue("begin"),
		End:         r.FormValue("end"),
		IsActive: func() bool {
			isActive, err := strconv.ParseBool(r.FormValue("active"))
			if err != nil {
				util.Log.Panicln(err)
			}
			return isActive
		}(),
	}
	model.Db.UpdateProject(project)
	http.Redirect(*w, r, "/", http.StatusFound)
}

func addProject(w *http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		util.Log.Panicln(err)
	}
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

	err = templ.ExecuteTemplate(w, "active-projects.html", activeProjects)
	if err != nil {
		util.Log.Println(err)
	}
}
