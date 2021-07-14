package controller

import (
	"html/template"
	"log"
	"net/http"

	"example.com/c1/model"
)

const stageThreePage string = "./web/view/stage-three/stage-three-access.html"

// StageThreeHandler - TODO: write about
func StageThreeHandler(writer http.ResponseWriter, request *http.Request) {

	activeWorkers, err := model.Db.RetrieveActiveWorkers()
	if err != nil {
		log.Println(err)
		ErrorPageHandler(writer, request)
	} else {
		for k := range activeWorkers {
			activeWorkers[k][2] = toHoursAndMinutes(activeWorkers[k][2])
		}
		templ, err := template.New("stageThree").ParseFiles(stageThreePage)

		if err != nil {
			log.Println(err)
			ErrorPageHandler(writer, request)
			return
		}

		err = templ.ExecuteTemplate(writer, "stage-three-access.html", activeWorkers)
		if err != nil {
			log.Println(err)
			ErrorPageHandler(writer, request)
			return
		}
	}
}
