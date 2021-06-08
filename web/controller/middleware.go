package controller

import (
	"errors"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
)

var ignoreURL []string = []string{
	"/save-time",
	"/log-out",
	"/error",
	"/favicon.ico",
	"/css/common.css",
	"/css/stage-one-style.css",
	"/css/stage-two-style.css",
	"/css/stage-three-style.css",
	"/css/stage-two-edit-project.css",
	"/css/stage-two-edit-worker.css",
	"/js/stage-one.js",
	"/js/stage-two.js",
	"/view/js/stage-two.js",
}

// AuthMiddleware ...
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Cache-Control", "no-cache")
		util.Log.Println(r.RequestURI)

		if isLinkPermited(r) {
			next.ServeHTTP(w, r)
			return
		}

		var worker *model.Worker
		var name string
		var pass string
		var err error

		if name, pass, err = verifyCookie(r); err == nil {
			worker, err = model.Db.GetUserByNameAndPassword(name, pass)
		}
		if err != nil {
			name, pass, err = parseForm(r)
			if err != nil {
				if r.RequestURI == "/" {
					next.ServeHTTP(w, r)
					return
				}
				util.Log.Println(err)
				ErrorPageHandler(w, r)
				return
			}

			worker, err = model.Db.GetUserByNameAndPassword(name, pass)
			if err != nil {
				util.Log.Println(err)
				ErrorPageHandler(w, r)
				return
			}
			setCookies(&w, name, pass)
		}

		switch worker.AccessLevel {
		case "1":
			StageOneHandler(worker, w, r)
		case "2":
			StageTwoHandler(w, r)
		case "3":
			StageThreeHandler(w, r)
		default:
			http.SetCookie(w, &http.Cookie{
				Name:   "name",
				MaxAge: -1,
			})
			http.SetCookie(w, &http.Cookie{
				Name:   "pass",
				MaxAge: -1,
			})
			HomePageHandler(w, r)
		}
	})
}

func isLinkPermited(r *http.Request) bool {
	for _, v := range ignoreURL {
		if v == r.RequestURI {
			return true
		}
	}
	return false
}

func setCookies(w *http.ResponseWriter, name string, pass string) {
	http.SetCookie(*w, &http.Cookie{
		Name:  "name",
		Value: name,
	})
	http.SetCookie(*w, &http.Cookie{
		Name:  "pass",
		Value: pass,
	})
}

func parseForm(r *http.Request) (string, string, error) {

	if err := r.ParseForm(); err != nil {
		util.Log.Println(err)
		return "", "", err
	}

	name := r.FormValue("name")
	pass := r.FormValue("password")

	if name != "" && pass != "" {
		r.Form.Del("name")
		r.Form.Del("password")

		return name, pass, nil
	}

	err := errors.New("invalid form entered")
	util.Log.Println(err)

	return "", "", err
}

func verifyCookie(r *http.Request) (string, string, error) {

	nameCookie, err := r.Cookie("name")
	if err != nil {
		util.Log.Println(err)
		return "", "", err
	}
	passCookie, err := r.Cookie("pass")
	if err != nil {
		util.Log.Println(err)
		return "", "", err
	}
	return nameCookie.Value, passCookie.Value, nil
}
