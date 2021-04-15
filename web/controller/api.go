package controller

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"example.com/c1/util"
)

func parseURI(r *http.Request, URI string) string {
	uri, err := url.Parse(r.RequestURI)

	if err != nil {
		util.Log.Panic(err)
	}
	return uri.Query().Get(URI)
}

// toHoursAndMinutes converts minutes to hours and minutes.
// For example: toHoursAndMinutes("61") returns "1:01m".
func toHoursAndMinutes(minutes string) string {

	workedMinutes, err := strconv.Atoi(minutes)
	if err != nil {
		util.Log.Println(err)
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
