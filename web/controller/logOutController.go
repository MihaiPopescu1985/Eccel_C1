package controller

import (
	"io/ioutil"
	"log"
	"net/http"
)

const logOutPage string = "./web/view/log-out.html"

// LogOutHandler ...
func LogOutHandler(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:   "name",
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "pass",
		MaxAge: -1,
	})

	fileContent, err := ioutil.ReadFile(logOutPage)
	if err != nil {
		log.Panicln(err)
	}

	if _, err := w.Write(fileContent); err != nil {
		log.Panicln(err)
	}
}
