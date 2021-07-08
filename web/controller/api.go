package controller

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func deleteTokenCookie(rw *http.ResponseWriter, r *http.Request) {
	cookie := getCookieFromRequest(r, "token")
	if cookie != nil {
		cookie.MaxAge = -1
		http.SetCookie(*rw, cookie)
	}
}

func getTokenFromCookie(r *http.Request) string {
	for _, c := range r.Cookies() {
		if c.Name == "token" {
			return c.Value
		}
	}
	return ""
}

func getCookieFromRequest(r *http.Request, cookieName string) *http.Cookie {
	for _, c := range r.Cookies() {
		if c.Name == cookieName {
			return c
		}
	}
	return nil
}

func parseURI(r *http.Request, URI string) string {
	uri, err := url.Parse(r.RequestURI)

	if err != nil {
		log.Panic(err)
	}
	return uri.Query().Get(URI)
}

// toHoursAndMinutes converts minutes to hours and minutes.
// For example: toHoursAndMinutes("61") returns "1:01".
func toHoursAndMinutes(minutes string) string {

	workedMinutes, err := strconv.Atoi(minutes)
	if err != nil {
		log.Println(err)
	}
	sign := ""
	if workedMinutes < 0 {
		workedMinutes *= -1
		sign = "-"
	}
	workedHours := int(workedMinutes / 60)
	workedMinutes = int(workedMinutes % 60)

	workedTime := sign + strconv.Itoa(workedHours) + ":"
	if workedMinutes < 10 {
		workedTime = workedTime + "0"
	}
	workedTime = workedTime + strconv.Itoa(workedMinutes)

	return workedTime
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
