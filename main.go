package main

import (
	"log"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/web/controller"
)

const serverPort string = ":8080"

func main() {

	model.Db = &model.MysqlDB{}
	err := model.Db.Init("")
	if err != nil {
		log.Fatalln(err)
	}

	if err := model.Db.Connect(); err != nil {
		log.Fatalln(err)
	}

	fs := http.FileServer(http.Dir("./web/view"))
	http.Handle("/", fs)

	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.LogOutHandler)
	http.HandleFunc("/save-time", controller.SaveTimeHandler)

	http.Handle("/index", controller.JwtMiddleware(http.HandlerFunc(controller.HomePageHandler)))
	log.Println(http.ListenAndServe(serverPort, nil))
}

/*
func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := filepath.Join("templates", "layout.html")
  fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

  tmpl, _ := template.ParseFiles(lp, fp)
  tmpl.ExecuteTemplate(w, "layout", nil)
}*/
