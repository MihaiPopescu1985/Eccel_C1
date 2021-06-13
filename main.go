package main

import (
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
	"example.com/c1/web/controller"
	"github.com/gorilla/mux"
)

const serverPort string = ":8080"

func main() {

	util.InitLogger()
	model.Db = &model.MysqlDB{}
	err := model.Db.Init("")
	if err != nil {
		util.Log.Println(err)
	}

	if err := model.Db.Connect(); err != nil {
		util.Log.Println(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomePageHandler)
	router.HandleFunc("/log-out", controller.LogOutHandler)
	router.HandleFunc("/error", controller.ErrorPageHandler)
	router.HandleFunc("/save-time", controller.SaveTimeHandler)

	router.NotFoundHandler = controller.PageNotFoundHandler{}
	router.Use(controller.AuthMiddleware)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/view/"))))
	util.Log.Println(http.ListenAndServe(serverPort, router))
}
