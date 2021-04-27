package controller

import "net/http"

func ErrorPageHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("There was an error processing your request"))
}
