package controller

import (
	"errors"
	"net/http"
	"strconv"

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
			http.Redirect(w, r, "/", 302)
		}

		if _, err := strconv.Atoi(worker.ID); err == nil {
			switch worker.AccessLevel {
			case "1":
				r.RequestURI = "/stage-one?workerId=" + worker.ID
				StageOneHandler(w, r)
			case "2":
				r.RequestURI = "/stage-two"
				StageTwoHandler(w, r)
			case "3":
				r.RequestURI = "/stage-three"
				StageThreeHandler(w, r)
			}
		} else {
			util.Log.Println(err)
			http.SetCookie(w, &http.Cookie{
				Name:   "name",
				MaxAge: -1,
			})
			http.SetCookie(w, &http.Cookie{
				Name:   "pass",
				MaxAge: -1,
			})
			next.ServeHTTP(w, r)
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

	err := errors.New("Invalid form entered.")
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
