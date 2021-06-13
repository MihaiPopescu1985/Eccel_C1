package controller

import (
	"errors"
	"log"
	"net/http"

	"example.com/c1/model"
)

// permitedURL is storing urls that don't need user authentication
var permitedURL []string = []string{
	"/",
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

// accessStages is storing the access map.key that an user must have in order to access the map.value url.
var accessStages map[string]string = map[string]string{
	"1": "/stage-one",
	"2": "/stage-two",
	"3": "/stage-three",
}

// AuthMiddleware ...
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Cache-Control", "no-cache")
		if isURLPermited(r) {
			next.ServeHTTP(w, r)
			return
		} else {
			name, pass, err := formFromRequest(r)
			if err != nil {
				log.Println(err)
				name, pass, err = verifyCookie(r)
				if err != nil {
					log.Println(err)
				} else {
					worker, err := model.Db.GetUserByNameAndPassword(name, pass)
					if err != nil {
						log.Println(err)
					} else {
						if v, ok := accessStages[worker.AccessLevel]; ok {
							if v == r.RequestURI {
								next.ServeHTTP(w, r)
								return
							}
						}
					}
				}
			} else {
				worker, err := model.Db.GetUserByNameAndPassword(name, pass)
				if err != nil {
					log.Println(err)
				} else {
					if v, ok := accessStages[worker.AccessLevel]; ok {
						if v == r.RequestURI {
							setCookies(&w, name, pass)
							next.ServeHTTP(w, r)
							return
						}
					}
				}
			}
			w.WriteHeader(http.StatusForbidden)
		}
	})
}

func isURLPermited(r *http.Request) bool {
	for _, v := range permitedURL {
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

func formFromRequest(r *http.Request) (string, string, error) {

	if err := r.ParseForm(); err != nil {
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
	return "", "", err
}

func verifyCookie(r *http.Request) (string, string, error) {

	nameCookie, err := r.Cookie("name")
	if err != nil {
		log.Println(err)
		return "", "", err
	}
	passCookie, err := r.Cookie("pass")
	if err != nil {
		log.Println(err)
		return "", "", err
	}
	return nameCookie.Value, passCookie.Value, nil
}
