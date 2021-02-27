package main

import (
	"log"
	"net/http"

	"example.com/c1/service"
	"example.com/c1/web/controller"
	"github.com/gorilla/mux"
)

func main() {

	service.Dao.Connect()

	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomePageHandler)
	router.HandleFunc("/stage-one", controller.StageOneHandler)
	router.HandleFunc("/stage-two", controller.StageTwoHandler)
	router.HandleFunc("/stage-three", controller.StageThreeHandler)
	router.HandleFunc("/log-out", controller.LogOutHandler)

	router.NotFoundHandler = controller.PageNotFoundHandler{}

	router.Use(controller.AuthMiddleware)

	log.Fatal(http.ListenAndServe(":8080", router))
}
