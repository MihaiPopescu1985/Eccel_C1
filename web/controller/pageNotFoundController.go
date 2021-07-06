package controller

import (
	"io/ioutil"
	"log"
	"net/http"
)

const notFoundPage string = "./web/view/not-found.html"

// PageNotFoundHandler ...
type PageNotFoundHandler struct{}

// ServeHTTP handles not found pages
func (notFound PageNotFoundHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	page, err := ioutil.ReadFile(notFoundPage)

	if err != nil {
		log.Println(err)
	}
	if _, err := writer.Write(page); err != nil {
		log.Println(err)
	}
}
