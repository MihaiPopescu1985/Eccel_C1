package controller

import (
	"io/ioutil"
	"log"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
	"golang.org/x/crypto/bcrypt"
)

var loginPage string = `./web/view/index.html`

func Login(rw http.ResponseWriter, r *http.Request) {

	deleteTokenCookie(&rw, r)

	switch r.Method {
	case http.MethodGet:
		fileContent, err := ioutil.ReadFile(loginPage)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Write(fileContent)

	case http.MethodPost:
		// parsing form
		err := r.ParseForm()
		if err != nil {
			log.Println("TODO: handle error parsing form")
			rw.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		userName := r.FormValue("name")
		password := r.FormValue("password")

		if userName == "" || password == "" {
			log.Println("TODO: handle invalid user or password")
			rw.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		// verify existing user in database
		user, err := model.Db.GetUserByNickname(userName)
		if err != nil || user == nil {
			log.Println(err)
			log.Println("TODO: handle invalid credentials")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			log.Println(err)
			log.Println("TODO: handle invalid password")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		// create jwt token
		token, err := util.GenJWTToken(user.ID, user.AccessLevel)
		util.AddActiveToken(token)

		if err != nil {
			log.Println("TODO: handle error generating jwt token")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// add token to response
		http.SetCookie(rw, &http.Cookie{
			Name:     "token",
			Value:    string(token),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})

		// redirect
		http.Redirect(rw, r, "/index", http.StatusSeeOther)

	default:
		rw.WriteHeader(http.StatusUnauthorized)
	}
}
