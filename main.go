package main

import (
	"log"
	"net/http"

	"example.com/c1/model"
	"example.com/c1/util"
	"example.com/c1/web/controller"
	"github.com/gorilla/mux"
)

const serverPort string = ":8181"

func main() {

	util.InitLogger()
	model.Db.Connect()
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
	router.HandleFunc("/stage-one", controller.StageOneHandler)
	router.HandleFunc("/stage-two", controller.StageTwoHandler)
	router.HandleFunc("/stage-three", controller.StageThreeHandler)
	router.HandleFunc("/log-out", controller.LogOutHandler)

	router.NotFoundHandler = controller.PageNotFoundHandler{}

	router.Use(controller.AuthMiddleware)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/view/"))))

	log.Fatal(http.ListenAndServe(serverPort, router))
}
