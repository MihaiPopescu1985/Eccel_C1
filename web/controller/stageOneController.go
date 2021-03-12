package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"example.com/c1/model"
	"example.com/c1/util"
)

const stageOnePage string = "./web/view/stageOneAccess.html"

type workerStatus struct {
	WorkerName string
	WorkerID   int
	Overtime   string
	TimeRaport map[int][]string
	Status     string
	WorkedTime string
}

// StageOneHandler TODO: edit function & function description.
// Based on worker's id got from url query,
// retrieve current month time report by calling
// database stored procedure: SELECT_TIME_RAPORT(WORKER_ID, CURRENT_MONTH).
func StageOneHandler(writer http.ResponseWriter, request *http.Request) {

	var pageContent workerStatus

	pageContent.setWorkerID(request)
	pageContent.setWorkerName()
	pageContent.setOvertime()
	pageContent.setStatusAndWorkedTime()
	pageContent.setDetailedTimeReport()

	servePage(&pageContent, &writer)
}

func servePage(pageContent *workerStatus, writer *http.ResponseWriter) {
	templ, err := template.New("stageOne").ParseFiles(stageOnePage, "./web/view/css/stage-one-style.css")
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(*writer, "stageOneAccess.html", pageContent)
	if err != nil {
		util.Log.Println(err)
	}
}

func (pageContent *workerStatus) setDetailedTimeReport() {
	currentYear := int(time.Now().Year())
	currentMonth := int(time.Now().Month())

	pageContent.TimeRaport = model.Db.RetrieveCurrentMonthTimeRaport(pageContent.WorkerID, currentMonth, currentYear)
	for k := range pageContent.TimeRaport {
		pageContent.TimeRaport[k][5] = model.ToHoursAndMinutes(pageContent.TimeRaport[k][5])
	}
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
	var err error
	pageContent.WorkerID, err = strconv.Atoi(r.URL.Query().Get("workerId"))

	if err != nil {
		util.Log.Println(err)
	}
}
