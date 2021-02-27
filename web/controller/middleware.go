package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/c1/model"
	"example.com/c1/service"
)

// AuthMiddleware ...
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")

		nameCookie, errName := r.Cookie("name")
		passCookie, errPass := r.Cookie("pass")

		var worker model.Worker
		var urlRedirect string = "/"

		if errName == nil && errPass == nil {
			worker = service.Dao.GetUserByNameAndPassword(nameCookie.Value, passCookie.Value)

			if &worker != nil {
				switch worker.AccessLevel {
				case 1:
					urlRedirect = "/stage-one?workerId=" + strconv.Itoa(worker.ID)

				case 2:
					urlRedirect = "/stage-two"

				case 3:
					urlRedirect = "/stage-three"
				}
			}
		}

		r.ParseForm()

		nameForm := r.FormValue("name")
		passForm := r.FormValue("password")

		if nameForm != "" || passForm != "" {

			fmt.Println("form")

			worker := service.Dao.GetUserByNameAndPassword(nameForm, passForm)
			if &worker != nil {
				http.SetCookie(w, &http.Cookie{
					Name:   "name",
					Value:  worker.Nickname,
					Secure: true,
				})
				http.SetCookie(w, &http.Cookie{
					Name:   "pass",
					Value:  worker.Password,
					Secure: true,
				})

				switch worker.AccessLevel {
				case 1:
					urlRedirect = "/stage-one?workerId=" + strconv.Itoa(worker.ID)

				case 2:
					urlRedirect = "/stage-two"

				case 3:
					urlRedirect = "/stage-three"
				}
			}
		}

		fmt.Println("redirect = " + urlRedirect)
		fmt.Println("request = " + r.RequestURI)

		if r.RequestURI == "/log-out" || urlRedirect == r.RequestURI {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, urlRedirect, 300)
		}
	})
}
