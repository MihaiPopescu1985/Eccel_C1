package controller

import (
	"io/ioutil"
	"log"
	"net/http"

	"example.com/c1/util"
)

const logOutPage string = "./web/view/log-out.html"

// LogOutHandler ...
func LogOutHandler(rw http.ResponseWriter, r *http.Request) {

	token := getTokenFromCookie(r)
	util.RemoveActiveToken([]byte(token))

	deleteTokenCookie(&rw, r)

	fileContent, err := ioutil.ReadFile(logOutPage)
	if err != nil {
		log.Panicln(err)
	}

	if _, err := rw.Write(fileContent); err != nil {
		log.Panicln(err)
	}
}
