package main

import (
	"log"
	"net/http"

	"example.com/c1/web/controller"
)

func main() {

	handler := http.HandlerFunc(controller.HomePageHandler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
