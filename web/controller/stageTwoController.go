package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"example.com/c1/model"
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
	case "delete-free-day":
		freeDayID := parseURI(r, "free-day")
		model.Db.DeleteFreeDay(freeDayID)
		http.Redirect(w, r, "/index?view=free-days", http.StatusFound)
	case "add-free-day":
		freeDay := parseURI(r, "free-day")
		model.Db.AddFreeDay(freeDay)
		http.Redirect(w, r, "/index?view=free-days", http.StatusFound)
	case "edit-project":
		editProject(&w, r)
	case "add-project":
		addProject(&w, r)
	case "sent-projects":
		sentProjectsView(&w, r)
	case "workers":
		showActiveProjects(w, r)
	case "free-days":
		pageContent, err := model.Db.RetrieveFreeDays()
		if err != nil {
			ErrorPageHandler(w, r)
		} else {
			serveFreeDaysPage(&w, r, pageContent)
		}
	default:
		showActiveProjects(w, r)
	}
}

func serveFreeDaysPage(w *http.ResponseWriter, r *http.Request, pageContent []string) {

	templ, err := template.New("freeDays").ParseFiles(stageTwofreeDaysPage)
	if err != nil {
		log.Println(err)
	}
	if err = templ.ExecuteTemplate(*w, "free-days.html", pageContent); err != nil {
		log.Println(err)
	}
}

func sentProjectsView(w *http.ResponseWriter, r *http.Request) {

	pageContent, err := model.Db.RetrieveSentProjects()
	if err != nil {
		log.Println(err)
		ErrorPageHandler(*w, r)
	} else {
		for k, v := range pageContent {
			pageContent[k] = toHoursAndMinutes(v)
		}

		templ, err := template.New("sentProjects").ParseFiles(sentProjectsPage)
		if err != nil {
			log.Println(err)
		}

		err = templ.ExecuteTemplate(*w, "sent-projects.html", pageContent)
		if err != nil {
			log.Println(err)
		}
	}
}

func editProject(w *http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	project := model.Project{
		ID:          r.FormValue("proj-id"),
		IPAddress:   r.FormValue("ip"),
		GeNumber:    r.FormValue("ge-no"),
		RoNumber:    r.FormValue("ro-no"),
		Description: r.FormValue("description"),
		DeviceID:    r.FormValue("dev-id"),
		Begin:       r.FormValue("start-date"),
		End:         r.FormValue("end-date"),
		IsActive: func() bool {
			isActive, err := strconv.ParseBool(r.FormValue("active"))
			if err != nil {
				log.Println(err)
			}
			return isActive
		}(),
	}
	model.Db.UpdateProject(project)
	http.Redirect(*w, r, "/index", http.StatusFound)
}

func addProject(w *http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
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
	http.Redirect(*w, r, "/index", http.StatusFound)
}

func showActiveProjects(w http.ResponseWriter, r *http.Request) {

	activeProjects, err := model.Db.RetrieveActiveProjects()
	if err != nil {
		log.Println(err)
		ErrorPageHandler(w, r)
		return
	} else {
		templ, err := template.New("stageTwo").ParseFiles(activeProjectsPage)
		if err != nil {
			log.Println(err)
			ErrorPageHandler(w, r)
			return
		}

		err = templ.ExecuteTemplate(w, "active-projects.html", activeProjects)
		if err != nil {
			log.Println(err)
			ErrorPageHandler(w, r)
			return
		}
	}
}
