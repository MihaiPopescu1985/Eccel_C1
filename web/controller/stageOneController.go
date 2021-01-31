package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const stageOnePage string = "/home/mihai/Documents/C1/project/web/view/stageOneAccess.html"

// StageOneHandler TODO: write about
func StageOneHandler(writer http.ResponseWriter, request *http.Request) {
	fileContent, err := ioutil.ReadFile(stageOnePage)

	// TODO: proper handle error
	if err != nil {
		fmt.Println(err)
	}
	writer.Write(fileContent)
}
