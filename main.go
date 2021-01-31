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
	router.NotFoundHandler = controller.PageNotFoundHandler{}
	log.Fatal(http.ListenAndServe(":8080", router))
}
