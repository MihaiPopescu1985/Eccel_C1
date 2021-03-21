package controller

import (
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"example.com/c1/model"
	"example.com/c1/util"
)

const (
	stageOnePage string = "./web/view/stageOneAccess.html"
	dateLayout   string = "2006-1-2 15:04:05"
	css          string = "./web/view/css/stage-one-style.css"
)

type workerStatus struct {
	WorkerName         string
	WorkerID           int
	Overtime           string
	TimeReport         map[int][]string
	StandardTimeReport map[string][]string
	Status             string
	WorkedTime         string
	ActiveProjects     []model.Project
}

// StageOneHandler TODO: edit function & function description.
// Based on worker's id got from url query,
// retrieve current month time report by calling
// database stored procedure: SELECT_TIME_RAPORT(WORKER_ID, CURRENT_MONTH).

func StageOneHandler(writer http.ResponseWriter, request *http.Request) {

	var pageContent = newWorkerStatus(request)

	saveForm(request, &pageContent)
	servePage(&pageContent, &writer)
}

func servePage(pageContent *workerStatus, writer *http.ResponseWriter) {
	templ, err := template.New("stageOne").ParseFiles(stageOnePage, css)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(*writer, "stageOneAccess.html", pageContent)
	if err != nil {
		util.Log.Println(err)
	}
}

func saveForm(r *http.Request, status *workerStatus) {
	if err := r.ParseForm(); err != nil {
		util.Log.Println(err)
	}

	formProject, err := strconv.Atoi(r.FormValue("projects"))
	if err != nil {
		util.Log.Println(err)
		return
	}

	formDay := r.FormValue("day")
	formStartHour := r.FormValue("startHour")
	formStartMinute := r.FormValue("startMinute")
	formStopHour := r.FormValue("stopHour")
	formStopMinute := r.FormValue("stopMinute")

	model.Db.AddWorkday(status.WorkerID, formProject,
		formatTime(formDay, formStartHour, formStartMinute),
		formatTime(formDay, formStopHour, formStopMinute))

	r.Form.Del("projects")
	r.Form.Del("day")
	r.Form.Del("startHour")
	r.Form.Del("startMinute")
	r.Form.Del("stopHour")
	r.Form.Del("stopMinute")
}

func formatTime(day, hour, minute string) string {

	var time strings.Builder
	time.WriteString(day)
	time.WriteString(" ")
	time.WriteString(hour)
	time.WriteString(":")
	time.WriteString(minute)
	time.WriteString(":00")

	return time.String()
}

func newWorkerStatus(request *http.Request) workerStatus {

	var pageContent workerStatus

	pageContent.setWorkerID(request)
	pageContent.setWorkerName()
	pageContent.setOvertime()
	pageContent.setStatusAndWorkedTime()
	pageContent.setTimeReport()
	pageContent.setActiveProjects()

	return pageContent
}

func (pageContent *workerStatus) setActiveProjects() {
	pageContent.ActiveProjects = model.Db.RetrieveActiveProjects()
}

func (pageContent *workerStatus) setTimeReport() {

	currentYear := int(time.Now().Year())
	currentMonth := int(time.Now().Month())

	pageContent.TimeReport = model.Db.RetrieveCurrentMonthTimeRaport(pageContent.WorkerID, currentMonth, currentYear)
	pageContent.StandardTimeReport = make(map[string][]string)

	for _, v := range pageContent.TimeReport {
		key := v[0] + " (" + v[1] + ")" // the key is formed by appending german project no. and romanian project no.

		date, err := time.Parse(dateLayout, v[3]) // parsing starttime
		if err != nil {
			util.Log.Println(err)
		}

		day := date.Day() - 1 // day must correspond with slice index wich starts at 0
		workedMinutes, err := strconv.Atoi(v[5])
		if err != nil {
			util.Log.Println(err)
		}

		if _, exists := pageContent.StandardTimeReport[key]; exists == false {
			pageContent.StandardTimeReport[key] = make([]string, 31)
		}

		currentMinutes := 0
		if pageContent.StandardTimeReport[key][day] != "" {
			if currentMinutes, err = strconv.Atoi(pageContent.StandardTimeReport[key][day]); err != nil {
				util.Log.Println(err)
			}
		}

		pageContent.StandardTimeReport[key][day] = strconv.Itoa(currentMinutes + workedMinutes)
	}

	for _, v := range pageContent.TimeReport {
		v[5] = toHoursAndMinutes(v[5])
	}
	for _, value := range pageContent.StandardTimeReport {
		for i, v := range value {
			if v != "" {
				value[i] = toHoursAndMinutes(v)
			}
		}
	}
}

// toHoursAndMinutes converts minutes to hours and minutes.
// For example: toHoursAndMinutes("61") returns "1h1m".
func toHoursAndMinutes(minutes string) string {

	workedMinutes, err := strconv.Atoi(minutes)
	if err != nil {
		util.Log.Println(err)
	}

	workedHours := workedMinutes / 60
	workedMinutes = workedMinutes - (workedHours * 60)

	workedTime := strconv.Itoa(workedHours) + ":" + strconv.Itoa(workedMinutes)
	return workedTime
}

func (pageContent *workerStatus) setStatusAndWorkedTime() {
	pageContent.Status, pageContent.WorkedTime = model.Db.RetrieveWorkerStatus(pageContent.WorkerID)
}

func (pageContent *workerStatus) setOvertime() {
	pageContent.Overtime = model.Db.RetrieveOvertime(pageContent.WorkerID)
}

func (pageContent *workerStatus) setWorkerName() {
	pageContent.WorkerName = model.Db.RetrieveWorkerName(pageContent.WorkerID)
}

func (pageContent *workerStatus) setWorkerID(r *http.Request) {

	url, err := url.Parse(r.RequestURI)

	if err != nil {
		util.Log.Println(err)
		http.Redirect(nil, nil, "/", http.StatusBadRequest)
	}

	pageContent.WorkerID, err = strconv.Atoi(url.Query().Get("workerId"))

	if err != nil {
		util.Log.Println(err)
		http.Redirect(nil, nil, "/", http.StatusBadRequest)
	}
}
