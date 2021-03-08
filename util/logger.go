package util

import (
	"log"
	"os"
	"time"
)

const (
	logDir  string = "logs"
	logFile string = "log.txt"
)

// Log is a global variable used to log events
var Log log.Logger

// InitLogger initialize the global Log variable.
func InitLogger() {

	createLoggerDir()

	file, fileErr := os.OpenFile(logDir+"/"+logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if fileErr != nil {
		panic(fileErr)
	}

	Log = *log.New(file, time.Now().String(), log.Lshortfile)
}

func createLoggerDir() {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0666)

		if err != nil {
			panic(err)
		}
	}
}
