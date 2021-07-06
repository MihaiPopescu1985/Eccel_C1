package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"example.com/c1/model"
)

const (
	workerStatusPage     string = "./web/view/stage-one/worker-status.html"
	detailedReportPage   string = "./web/view/stage-one/detailed-view.html"
	standardReportPage   string = "./web/view/stage-one/standard-view.html"
	addWorkdayPage       string = "./web/view/stage-one/add-workday.html"
	stageOneFreeDaysPage string = "./web/view/stage-one/free-days.html"
	dateLayout           string = "2006-1-2 15:04:05"
)

type workerStatus struct {
	Worker     *model.Worker
	Status     string
	WorkedTime string
	Overtime   string
}

type timeReport struct {
	Worker     *model.Worker
	TimeReport [][]string
}

type standardTimeReport struct {
	Worker     *model.Worker
	TimeReport map[string][]string
}

type addWorkday struct {
	Worker         *model.Worker
	ActiveProjects []model.Project
}

type freeDays struct {
	Worker   *model.Worker
	FreeDays []string
}

// StageOneHandler TODO: edit function & function description.
// Based on worker's id retrieve current month time report by calling
// database stored procedure: SELECT_TIME_RAPORT(WORKER_ID, CURRENT_MONTH).

func StageOneHandler(worker *model.Worker, writer http.ResponseWriter, request *http.Request) {

	switch parseURI(request, "view") {
	case "detailed-view":

		var pageContent timeReport
		var err error

		pageContent.Worker = worker
		pageContent.TimeReport, err = getDetailedReport(worker.ID)
		if err != nil {
			log.Println(err)
			ErrorPageHandler(writer, request)
		} else {
			prepareDetailedReport(pageContent.TimeReport)
			serveDetailedPage(&pageContent, &writer)
		}

	case "standard-view":

		var pageContent standardTimeReport
		var err error

		pageContent.Worker = worker
		pageContent.TimeReport, err = getStandardReport(worker.ID)
		if err != nil {
			log.Println(err)
			ErrorPageHandler(writer, request)
		} else {
			prepareStandardReport(pageContent.TimeReport)
			serveStandardPage(&pageContent, &writer)
		}

	case "add-workday":

		var pageContent addWorkday
		var err error

		pageContent.Worker = worker
		pageContent.ActiveProjects, err = model.Db.RetrieveActiveProjects()

		if err != nil {
			log.Println(err)
			ErrorPageHandler(writer, request)
		} else {
			serveAddWorkdayPage(&pageContent, &writer)
		}

	case "save-workday":
		saveWorkdayForm(&writer, request, *worker)

	case "free-days":

		var pageContent freeDays
		var err error

		pageContent.Worker = worker
		pageContent.FreeDays, err = model.Db.RetrieveFreeDays()

		if err != nil {
			log.Println(err)
			ErrorPageHandler(writer, request)
		} else {
			pageContent.serveFreeDaysPage(&writer, request)
		}

	default:
		var pageContent workerStatus
		var err error

		pageContent.Worker = worker
		err = pageContent.setStatusAndWorkedTime()
		if err != nil {
			log.Println(err)
			ErrorPageHandler(writer, request)
			break
		}
		err = pageContent.setOvertime()
		if err != nil {
			log.Println(err)
			ErrorPageHandler(writer, request)
			break
		}
		serveStatusPage(&pageContent, &writer)
	}
}

func (pageContent *freeDays) serveFreeDaysPage(w *http.ResponseWriter, r *http.Request) {

	templ, err := template.New("freeDays").ParseFiles(stageOneFreeDaysPage)
	if err != nil {
		log.Println(err)
	}
	if err = templ.ExecuteTemplate(*w, "free-days.html", *pageContent); err != nil {
		log.Println(err)
	}
}

func saveWorkdayForm(w *http.ResponseWriter, r *http.Request, worker model.Worker) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	formProject := r.FormValue("projects")
	if formProject == "" {
		log.Println("Error parsing project ID.")
		return
	}

	formDay := r.FormValue("day")
	formStartHour := r.FormValue("start-hour")
	formStartMinute := r.FormValue("start-minute")
	formStopHour := r.FormValue("stop-hour")
	formStopMinute := r.FormValue("stop-minute")

	if formDay != "" && formStartHour != "" && formStartMinute != "" && formStopHour != "" && formStopMinute != "" {
		model.Db.AddWorkday(worker.ID, formProject,
			formatTime(formDay, formStartHour, formStartMinute),
			formatTime(formDay, formStopHour, formStopMinute))

		http.Redirect(*w, r, "/", http.StatusFound)
	}
}

func serveAddWorkdayPage(workday *addWorkday, writer *http.ResponseWriter) {

	templ, err := template.New("addWorkday").ParseFiles(addWorkdayPage)
	if err != nil {
		log.Println(err)
	}

	err = templ.ExecuteTemplate(*writer, "add-workday.html", *workday)
	if err != nil {
		log.Println(err)
	}
}

func serveStandardPage(report *standardTimeReport, writer *http.ResponseWriter) {

	templ, err := template.New("standardReport").ParseFiles(standardReportPage)
	if err != nil {
		log.Println(err)
	}

	err = templ.ExecuteTemplate(*writer, "standard-view.html", *report)
	if err != nil {
		log.Println(err)
	}
}

func prepareStandardReport(report map[string][]string) {

	for _, min := range report {
		for k := range min {
			if min[k] != "" {
				min[k] = toHoursAndMinutes(min[k])
			}
		}
	}
}

func getStandardReport(wID string) (map[string][]string, error) {

	report, err := getDetailedReport(wID)
	if err != nil {
		return nil, err
	}
	var standardReport = make(map[string][]string)

	for _, v := range report {
		key := v[0] + " (" + v[1] + ")" // the key is formed by appending german project no. and romanian project no.

		date, err := time.Parse(dateLayout, v[3]) // parsing starttime
		if err != nil {
			log.Println(err)
		}

		day := date.Day() - 1 // day must correspond with slice index wich starts at 0
		workedMinutes, err := strconv.Atoi(v[5])
		if err != nil {
			log.Println(err)
		}

		if _, exists := standardReport[key]; !exists {
			standardReport[key] = make([]string, 31)
		}

		currentMinutes := 0
		if standardReport[key][day] != "" {
			if currentMinutes, err = strconv.Atoi(standardReport[key][day]); err != nil {
				log.Println(err)
			}
		}
		standardReport[key][day] = strconv.Itoa(currentMinutes + workedMinutes)
	}
	return standardReport, err
}

func serveDetailedPage(report *timeReport, writer *http.ResponseWriter) {
	templ, err := template.New("detailedReport").ParseFiles(detailedReportPage)
	if err != nil {
		log.Println(err)
	}

	err = templ.ExecuteTemplate(*writer, "detailed-view.html", *report)
	if err != nil {
		log.Println(err)
	}
}

func prepareDetailedReport(rawReport [][]string) {

	for k, v := range rawReport {
		v[5] = toHoursAndMinutes(v[5])

		//delete(rawReport, k)
		rawReport[k] = v
	}
}

func getDetailedReport(wID string) ([][]string, error) {
	currentYear := strconv.Itoa(time.Now().Year())
	currentMonth := strconv.Itoa((int(time.Now().Month())))

	report, err := model.Db.RetrieveCurrentMonthTimeRaport(wID, currentMonth, currentYear)
	return report, err
}

func serveStatusPage(pageContent *workerStatus, writer *http.ResponseWriter) {
	templ, err := template.New("workerStatus").ParseFiles(workerStatusPage)
	if err != nil {
		log.Println(err)
	}

	err = templ.ExecuteTemplate(*writer, "worker-status.html", *pageContent)
	if err != nil {
		log.Println(err)
	}
}

func (pageContent *workerStatus) setOvertime() error {
	var err error
	pageContent.Overtime, err = model.Db.RetrieveMinutesOvertime(pageContent.Worker.ID)
	if err == nil {
		pageContent.Overtime = toHoursAndMinutes(pageContent.Overtime)
	}
	return err
}

func (pageContent *workerStatus) setStatusAndWorkedTime() error {
	var err error
	pageContent.Status, pageContent.WorkedTime, err = model.Db.RetrieveWorkerStatus(pageContent.Worker.ID)
	if err == nil {
		pageContent.WorkedTime = toHoursAndMinutes(pageContent.WorkedTime)
	}
	return err
}
