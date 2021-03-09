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
	TimeRaport map[int][]string
	Status     string
	WorkedTime string

	FreeDays map[int]string
}

// StageOneHandler TODO: edit function & function description.
// Based on worker's id got from url query,
// retrieve current month time report by calling
// database stored procedure: SELECT_TIME_RAPORT(WORKER_ID, CURRENT_MONTH).
// This page must also display worker's status (inactive/active/break)
// and current working time.
func StageOneHandler(writer http.ResponseWriter, request *http.Request) {

	workerID, err := strconv.Atoi(request.URL.Query().Get("workerId"))
	if err != nil {
		util.Log.Println(err)
	}

	pageContent := workerStatus{
		WorkerName: "",
		WorkerID:   workerID,
		TimeRaport: make(map[int][]string, 0),
		Status:     "INACTIVE",
		WorkedTime: "0",

		FreeDays: make(map[int]string, 0),
	}

	pageContent.WorkerName = model.Db.RetrieveWorkerName(workerID)

	pageContent.Status, pageContent.WorkedTime = model.Db.RetrieveWorkerStatus(workerID)
	currentMonth := int(time.Now().Month())

	pageContent.TimeRaport = model.Db.RetrieveCurrentMonthTimeRaport(workerID, currentMonth)
	pageContent.FreeDays = model.Db.RetrieveFreeDays()

	templ, err := template.New("stageOne").ParseFiles(stageOnePage, "./web/view/css/stage-one-style.css")
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(writer, "stageOneAccess.html", pageContent)
	if err != nil {
		util.Log.Println(err)
	}
}
