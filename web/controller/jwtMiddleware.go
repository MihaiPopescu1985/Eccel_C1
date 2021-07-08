package controller

import (
	"log"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token := getTokenFromCookie(r)

		// if Authentication header contains a valid token
		if util.IsGoodToken([]byte(token)) && util.IsTokenActive([]byte(token)) {
			// try to retrieve user id from token
			userID, err := util.GetUserIDFromToken([]byte(token))
			if err != nil {
				log.Println(err)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				// if we have the user id, try to get user from database
				user, err := model.Db.GetWorker(userID)
				if err != nil {
					log.Println(err)
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}

				newToken := util.RefreshToken([]byte(token))
				http.SetCookie(rw, &http.Cookie{
					Name:     "token",
					Value:    string(newToken),
					HttpOnly: true,
				})

				// redirect user
				switch user.AccessLevel {
				case "1":
					StageOneHandler(user, rw, r)
					return
				case "2":
					StageTwoHandler(rw, r)
					return
				case "3":
					StageThreeHandler(rw, r)
					return
				default:
					rw.WriteHeader(http.StatusUnauthorized)
					return
				}
			}
		}
		http.Redirect(rw, r, "/login", http.StatusFound)
	})
}
