package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const logOutPage string = "./web/view/logOut.html"

// LogOutHandler ...
func LogOutHandler(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:   "name",
		Value:  "",
		Secure: true,
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "pass",
		Value:  "",
		Secure: true,
		MaxAge: -1,
	})

	fileContent, err := ioutil.ReadFile(logOutPage)
	if err != nil {
		fmt.Println("Error opening logOutPage file.")
		fmt.Println(err)
	}

	w.Write(fileContent)
}
