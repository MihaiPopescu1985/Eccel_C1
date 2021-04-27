package main

import (
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
	"example.com/c1/web/controller"
	"github.com/gorilla/mux"
)

const serverPort string = ":8181"

func main() {

	util.InitLogger()

	util.Log.Println(model.Db.Init(""))
	util.Log.Println(model.Db.Connect())
	/*
		endPoint := c1device.C1Device{
			IP:        "192.168.0.91",
			WsChannel: make(chan []byte),
		}
		endPoint.UseDevice()

		operational := c1device.C1Device{
			IP:        "192.168.0.92",
			WsChannel: make(chan []byte),
		}
		operational.UseDevice()
	*/

	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomePageHandler)
	router.HandleFunc("/log-out", controller.LogOutHandler)
	router.HandleFunc("/error", controller.ErrorPageHandler)

	router.NotFoundHandler = controller.PageNotFoundHandler{}
	router.Use(controller.AuthMiddleware)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/view/"))))
	util.Log.Println(http.ListenAndServe(serverPort, router))
}
