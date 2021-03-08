package controller

import (
	"io/ioutil"
	"net/http"

	"example.com/c1/util"
)

const notFoundPage string = "./web/view/notFound.html"

// PageNotFoundHandler ...
type PageNotFoundHandler struct{}

// ServeHTTP handles not found pages
func (notFound PageNotFoundHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	page, err := ioutil.ReadFile(notFoundPage)

	if err != nil {
		util.Log.Println(err)
	}
	if _, err := writer.Write(page); err != nil {
		util.Log.Println(err)
	}
}
