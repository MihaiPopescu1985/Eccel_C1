package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const stageTwoPage string = "/home/mihai/Documents/C1/project/web/view/stageTwoAccess.html"

// StageTwoHandler TODO: write about
func StageTwoHandler(writer http.ResponseWriter, request *http.Request) {
	fileContent, err := ioutil.ReadFile(stageTwoPage)

	// TODO: proper handle error
	if err != nil {
		fmt.Println(err)
	}
	// TODO: handle error when writing file content
	writer.Write(fileContent)
}
