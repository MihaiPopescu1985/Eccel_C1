package controller

import (
	"errors"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
)

var ignoreURL []string = []string{
	"/log-out",
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
		for _, v := range ignoreURL {
			if v == r.RequestURI {
				next.ServeHTTP(w, r)
				return
			}
		}
		var worker model.Worker

		if name, pass, err := verifyCookie(r); err == nil {
			worker = model.Db.GetUserByNameAndPassword(name, pass)

		} else if name, pass, err := parseForm(r); err == nil {
			worker = model.Db.GetUserByNameAndPassword(name, pass)
			setCookies(&w, name, pass)
		}

		switch worker.AccessLevel {
		case "1":
			StageOneHandler(&worker, w, r)
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

func setCookies(w *http.ResponseWriter, name string, pass string) {
	http.SetCookie(*w, &http.Cookie{
		Name:   "name",
		Value:  name,
		Secure: true,
	})
	http.SetCookie(*w, &http.Cookie{
		Name:   "pass",
		Value:  pass,
		Secure: true,
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
