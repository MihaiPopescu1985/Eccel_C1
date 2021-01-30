package controller

import (
	"io/ioutil"
	"net/http"
)

const notFoundPage string = "/home/mihai/Documents/C1/project/web/view/notFound.html"

// PageNotFoundHandler ...
type PageNotFoundHandler struct{}

// ServeHTTP handles not found pages
func (notFound PageNotFoundHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	page, _ := ioutil.ReadFile(notFoundPage)
	writer.Write(page)
}
