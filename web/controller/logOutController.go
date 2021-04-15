package controller

import (
	"io/ioutil"
	"net/http"

	"example.com/c1/util"
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
		util.Log.Panicln(err)
	}

	if _, err := w.Write(fileContent); err != nil {
		util.Log.Panicln(err)
	}
}
