package controller

import (
	"net/http"
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
		/*
			var dao service.DAO
			dao.Connect()
			defer dao.CloseConnection()

			var urlRedirect string = "/"

			dateCookie, errDate := r.Cookie("date")
			valueCookie, errValue := r.Cookie("value")
			userCookie, errUser := r.Cookie("user")

			if errDate != nil || errValue != nil || errUser != nil {
				if token.Date != dateCookie.Value || token.Token != valueCookie.Value || &token != nil {

				}

				workerID, _ := strconv.Atoi(userCookie.Value)
				token := dao.GetActiveToken(workerID)

				r.ParseForm()

				name := r.FormValue("name")
				password := r.FormValue("password")

				if name != "" || password != "" {

					worker := dao.GetUserByNameAndPassword(name, password)
					dao.SaveWebToken(worker.ID)
					token = dao.GetActiveToken(worker.ID)

					r.AddCookie(&http.Cookie{
						Name:   "date",
						Value:  token.Date,
						MaxAge: 0,
						Secure: true,
					})
					r.AddCookie(&http.Cookie{
						Name:   "value",
						Value:  token.Token,
						MaxAge: 0,
						Secure: true,
					})
					r.AddCookie(&http.Cookie{
						Name:   "user",
						Value:  strconv.Itoa(token.WorkerID),
						MaxAge: 0,
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

			if r.RequestURI == "/" {

				http.Redirect(w, r, urlRedirect, 300)
			} else {
				next.ServeHTTP(w, r)
			}
		*/
	})
}
