package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"example.com/c1/service"
)

const stageOnePage string = "/home/mihai/Documents/C1/project/web/view/stageOneAccess.html"

type workerStatus struct {
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
		workerID,
		make(map[int][]string, 0),
		"INACTIVE",
		"0",
	}

	var dao service.DAO
	dao.Connect()
	defer dao.CloseConnection()

	var command string = "CALL SELECT_WORKER_STATUS(" + strconv.Itoa(pageContent.WorkerID) + ");"
	pageContent.Status, pageContent.WorkedTime = dao.RetrieveWorkerStatus(dao.ExecuteQuery(command))

	currentMonth := int(time.Now().Month())
	command = "CALL SELECT_MONTH_TIME_RAPORT(" + strconv.Itoa(pageContent.WorkerID) + ", " + strconv.Itoa(currentMonth) + ");"

	pageContent.TimeRaport = dao.RetrieveCurrentMonthTimeRaport(dao.ExecuteQuery(command))

	templ, err := template.New("stageOne").ParseFiles(stageOnePage)
	if err != nil {
		fmt.Println(err)
	}

	err = templ.ExecuteTemplate(writer, "stageOneAccess.html", pageContent)
	if err != nil {
		fmt.Println(err)
	}
}
