package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"example.com/c1/service"
)

const stageOnePage string = "./web/view/stageOneAccess.html"

type workerStatus struct {
	WorkerName string
	WorkerID   int
	TimeRaport map[int][]string
	Status     string
	WorkedTime string
}

// StageOneHandler TODO: edit function & function description.
// Based on worker's id got from url query,
// retrieve current month time report by calling
// database stored procedure: SELECT_TIME_RAPORT(WORKER_ID, CURRENT_MONTH).
// This page must also display worker's status (inactive/active/break)
// and current working time.
func StageOneHandler(writer http.ResponseWriter, request *http.Request) {

	workerID, _ := strconv.Atoi(request.URL.Query().Get("workerId"))
	pageContent := workerStatus{
		"",
		workerID,
		make(map[int][]string, 0),
		"INACTIVE",
		"0",
	}

	pageContent.WorkerName = service.Dao.RetrieveWorkerName(workerID)

	pageContent.Status, pageContent.WorkedTime = service.Dao.RetrieveWorkerStatus(workerID)
	currentMonth := int(time.Now().Month())

	pageContent.TimeRaport = service.Dao.RetrieveCurrentMonthTimeRaport(workerID, currentMonth)

	templ, err := template.New("stageOne").ParseFiles(stageOnePage)
	if err != nil {
		fmt.Println(err)
	}

	err = templ.ExecuteTemplate(writer, "stageOneAccess.html", pageContent)
	if err != nil {
		fmt.Println(err)
	}
}
