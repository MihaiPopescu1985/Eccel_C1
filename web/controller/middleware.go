package controller

import (
	"net/http"
	"strconv"
	"strings"

	"example.com/c1/model"
	"example.com/c1/util"
)

// AuthMiddleware ...
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.RequestURI, "/css/") || strings.HasPrefix(r.RequestURI, "/js/") {
			next.ServeHTTP(w, r)
		} else {

			nameCookie, errName := r.Cookie("name")
			passCookie, errPass := r.Cookie("pass")

			var worker model.Worker
			var urlRedirect string = "/"

			if errName != nil {
				util.Log.Println(errName)

			} else if errPass != nil {
				util.Log.Println(errPass)

			} else {
				worker = model.Db.GetUserByNameAndPassword(nameCookie.Value, passCookie.Value)

				if &worker != nil {
					switch worker.AccessLevel {
					case 1:
						urlRedirect = "/stage-one?workerId=" + strconv.Itoa(int(worker.ID))

					case 2:
						urlRedirect = "/stage-two"

					case 3:
						urlRedirect = "/stage-three"
					}
				}
			}

			if err := r.ParseForm(); err != nil {
				util.Log.Println(err)
			}

			nameForm := r.FormValue("name")
			passForm := r.FormValue("password")

			if nameForm != "" || passForm != "" {

				worker := model.Db.GetUserByNameAndPassword(nameForm, passForm)
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
						urlRedirect = "/stage-one?workerId=" + strconv.Itoa(int(worker.ID))

					case 2:
						urlRedirect = "/stage-two"

					case 3:
						urlRedirect = "/stage-three"
					}
				}
			}

			if r.RequestURI == "/log-out" || urlRedirect == r.RequestURI {
				next.ServeHTTP(w, r)
			} else {
				http.Redirect(w, r, urlRedirect, 300)
			}
		}

	})
}
