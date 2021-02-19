package controller

import (
	"net/http"
	"strconv"

	"example.com/c1/service"
)

/*
verifica cookie
daca sunt cookie
	daca se obtine id-ul lucratorului
		se cauta token-ul salvat in baza de date (daca este activ)
		daca token-ul este activ si corespunde cu cookie
			se trece mai departe
		daca token-ul nu este activ sau nu corespunde cu cookie

			se redirectioneaza catre pagina principala
	daca nu se obtine id-ul lucratorului
		se redirectioneaza catre pagina principala
daca nu sunt cookie
	se redirectioneaza catre pagina principala
*/

// AuthMiddleware ...
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		urlRedirect := "/"

		var dao service.DAO
		dao.Connect()
		defer dao.CloseConnection()

		userID, err := verifyUserCookies(r)

		// if user id is found inside cookie
		if err == nil {
			// verify the rest of the cookies
			if verifyDateAndValueCookies(userID, r, dao) {
				// if everithing is ok, move forward
				next.ServeHTTP(w, r)
			}
		} else {
			http.Redirect(w, r, urlRedirect, 300)
		}
		/*
			 {
				if token.Date != dateCookie.Value || token.Token != valueCookie.Value || &token != nil {

				}

				workerID, _ := strconv.Atoi(userCookie.Value)
				token := dao.GetActiveToken(workerID)



			if r.RequestURI == "/" {

				http.Redirect(w, r, urlRedirect, 300)
			} else {
				next.ServeHTTP(w, r)
			}
		*/
	})
}

// Returns the user id from parsing cookie and error if the cookie cannot be parsed
func verifyUserCookies(r *http.Request) (int, error) {
	userCookie, errUser := r.Cookie("user")

	if errUser == nil {
		return strconv.Atoi(userCookie.Value)
	}
	return -1, errUser
}

func verifyDateAndValueCookies(userID int, r *http.Request, dao service.DAO) bool {
	token := dao.GetActiveToken(userID)
	if &token != nil {
		dateCookie, errDate := r.Cookie("date")
		valueCookie, errValue := r.Cookie("value")

		if errDate != nil && errValue != nil {
			if dateCookie.Value == token.Date && valueCookie.Value == token.Token {
				return true
			}
		}
	}
	return false
}
