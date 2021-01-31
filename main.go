package main

import (
	"log"
	"net/http"

	"example.com/c1/web/controller"
	"github.com/gorilla/mux"
)

const homePage string = "/home/mihai/Documents/C1/project/web/view/directorPage.html"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomePageHandler)
	router.NotFoundHandler = controller.PageNotFoundHandler{}
	log.Fatal(http.ListenAndServe(":8080", router))
}
