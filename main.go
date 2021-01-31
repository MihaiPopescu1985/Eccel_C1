package main

import (
	"log"
	"net/http"

	"example.com/c1/web/controller"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomePageHandler)
	router.HandleFunc("/stage-one", controller.StageOneHandler)
	router.HandleFunc("/stage-two", controller.StageTwoHandler)
	router.HandleFunc("/stage-three", controller.StageThreeHandler)

	router.NotFoundHandler = controller.PageNotFoundHandler{}
	log.Fatal(http.ListenAndServe(":8080", router))
}
