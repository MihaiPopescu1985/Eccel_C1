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
	workerStatusPage   string = "./web/view/stage-one/WorkerStatus.html"
	detailedReportPage string = "./web/view/stage-one/DetailedView.html"
	standardReportPage string = "./web/view/stage-one/StandardView.html"
	addWorkdayPage     string = "./web/view/stage-one/AddWorkday.html"
	freeDaysPage       string = "./web/view/stage-one/FreeDays.html"
	dateLayout         string = "2006-1-2 15:04:05"
)

type workerStatus struct {
	Worker     *model.Worker
	Status     string
	WorkedTime string
	Overtime   string
}

type timeReport struct {
	Worker     *model.Worker
	TimeReport map[string][]string
}

type addWorkday struct {
	Worker         *model.Worker
	ActiveProjects []model.Project
}

type freeDays struct {
	Worker   *model.Worker
	FreeDays map[int]string
}

// StageOneHandler TODO: edit function & function description.
// Based on worker's id retrieve current month time report by calling
// database stored procedure: SELECT_TIME_RAPORT(WORKER_ID, CURRENT_MONTH).

func StageOneHandler(worker *model.Worker, writer http.ResponseWriter, request *http.Request) {

	switch parseURI(request, "view") {
	case "detailed-view":

		var pageContent timeReport
		pageContent.Worker = worker

		pageContent.TimeReport = getDetailedReport(worker.ID)
		prepareDetailedReport(pageContent.TimeReport)
		serveDetailedPage(&pageContent, &writer)

	case "standard-view":

		var pageContent timeReport
		pageContent.Worker = worker

		pageContent.TimeReport = getStandardReport(worker.ID)
		prepareStandardReport(pageContent.TimeReport)

		serveStandardPage(&pageContent, &writer)

	case "add-workday":

		pageContent := addWorkday{
			Worker:         worker,
			ActiveProjects: model.Db.RetrieveActiveProjects(),
		}
		serveAddWorkdayPage(&pageContent, &writer)

	case "save-workday":
		saveForm(&writer, request, *worker)

	case "free-days":

		pageContent := freeDays{
			Worker:   worker,
			FreeDays: model.Db.RetrieveFreeDays(),
		}
		serveFreeDaysPage(&writer, request, &pageContent)

	default:
		var pageContent workerStatus

		pageContent.Worker = worker
		pageContent.setStatusAndWorkedTime()
		pageContent.setOvertime()

		serveStatusPage(&pageContent, &writer)
	}
}

func serveFreeDaysPage(w *http.ResponseWriter, r *http.Request, pageContent *freeDays) {

	templ, err := template.New("freeDays").ParseFiles(freeDaysPage)
	if err != nil {
		util.Log.Println(err)
	}
	if err = templ.ExecuteTemplate(*w, "FreeDays.html", *pageContent); err != nil {
		util.Log.Println(err)
	}
}

func saveForm(w *http.ResponseWriter, r *http.Request, worker model.Worker) {
	if err := r.ParseForm(); err != nil {
		util.Log.Println(err)
	}

	formProject := r.FormValue("projects")
	if formProject == "" {
		util.Log.Println("Error parsing project ID.")
		return
	}

	formDay := r.FormValue("day")
	formStartHour := r.FormValue("startHour")
	formStartMinute := r.FormValue("startMinute")
	formStopHour := r.FormValue("stopHour")
	formStopMinute := r.FormValue("stopMinute")

	if formDay != "" && formStartHour != "" && formStartMinute != "" && formStopHour != "" && formStopMinute != "" {
		model.Db.AddWorkday(worker.ID, formProject,
			formatTime(formDay, formStartHour, formStartMinute),
			formatTime(formDay, formStopHour, formStopMinute))

		http.Redirect(*w, r, "/", http.StatusFound)
	}
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

func serveAddWorkdayPage(workday *addWorkday, writer *http.ResponseWriter) {

	templ, err := template.New("addWorkday").ParseFiles(addWorkdayPage)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(*writer, "AddWorkday.html", *workday)
	if err != nil {
		util.Log.Println(err)
	}
}

func serveStandardPage(report *timeReport, writer *http.ResponseWriter) {

	templ, err := template.New("standardReport").ParseFiles(standardReportPage)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(*writer, "StandardView.html", *report)
	if err != nil {
		util.Log.Println(err)
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

func getStandardReport(wID string) map[string][]string {
	report := getDetailedReport(wID)
	var standardReport = make(map[string][]string)

	for _, v := range report {
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

		if _, exists := standardReport[key]; !exists {
			standardReport[key] = make([]string, 31)
		}

		currentMinutes := 0
		if standardReport[key][day] != "" {
			if currentMinutes, err = strconv.Atoi(standardReport[key][day]); err != nil {
				util.Log.Println(err)
			}
		}
		standardReport[key][day] = strconv.Itoa(currentMinutes + workedMinutes)
	}
	return standardReport
}

func serveDetailedPage(report *timeReport, writer *http.ResponseWriter) {
	templ, err := template.New("detailedReport").ParseFiles(detailedReportPage)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(*writer, "DetailedView.html", *report)
	if err != nil {
		util.Log.Println(err)
	}
}

func prepareDetailedReport(rawReport map[string][]string) {

	for k, v := range rawReport {
		v[5] = toHoursAndMinutes(v[5])

		delete(rawReport, k)
		rawReport[k] = v
	}
}

func getDetailedReport(wID string) map[string][]string {
	currentYear := strconv.Itoa(time.Now().Year())
	currentMonth := strconv.Itoa((int(time.Now().Month())))

	report := model.Db.RetrieveCurrentMonthTimeRaport(wID, currentMonth, currentYear)
	return report
}

func parseURI(r *http.Request, URI string) string {
	uri, err := url.Parse(r.RequestURI)

	if err != nil {
		util.Log.Panic(err)
	}
	return uri.Query().Get(URI)
}

func serveStatusPage(pageContent *workerStatus, writer *http.ResponseWriter) {
	templ, err := template.New("workerStatus").ParseFiles(workerStatusPage)
	if err != nil {
		util.Log.Println(err)
	}

	err = templ.ExecuteTemplate(*writer, "WorkerStatus.html", *pageContent)
	if err != nil {
		util.Log.Println(err)
	}
}

func (pageContent *workerStatus) setOvertime() {
	pageContent.Overtime = model.Db.RetrieveOvertime(pageContent.Worker.ID)
	pageContent.Overtime = toHoursAndMinutes(pageContent.Overtime)
}

// toHoursAndMinutes converts minutes to hours and minutes.
// For example: toHoursAndMinutes("61") returns "1:01m".
func toHoursAndMinutes(minutes string) string {

	workedMinutes, err := strconv.Atoi(minutes)
	if err != nil {
		util.Log.Panicln(err)
	}

	workedHours := workedMinutes / 60
	if workedMinutes < 0 {
		workedMinutes *= -1
	}
	workedMinutes = workedMinutes % 60

	var workedTime strings.Builder

	workedTime.WriteString(strconv.Itoa(workedHours))
	workedTime.WriteString(":")

	if workedMinutes < 10 {
		workedTime.WriteString("0")
	}

	workedTime.WriteString(strconv.Itoa(workedMinutes))

	return workedTime.String()
}

func (pageContent *workerStatus) setStatusAndWorkedTime() {
	pageContent.Status, pageContent.WorkedTime = model.Db.RetrieveWorkerStatus(pageContent.Worker.ID)
	pageContent.WorkedTime = toHoursAndMinutes(pageContent.WorkedTime)
}
