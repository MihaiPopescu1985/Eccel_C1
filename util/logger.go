package util

import (
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// Log is a global variable used to log events
var Log Logger

// Logger ...
type Logger struct {
}

// Init is used for initializing logger, for example
// if a database connection is needed.
func (log *Logger) Init() {

	createLoggerDir()
	createDayLogFile()
}

func createDayLogFile() {
	year, month, day := time.Now().Date()

	filename := "logs/" + strconv.Itoa(year) + "-" + month.String() + "-" + strconv.Itoa(day) + ".txt"
	err := ioutil.WriteFile(filename, []byte(filename+"\n\n"), 0644)

	if err != nil {
		panic("Error initialising logger. Check if a logger file can be created.")
	}
}

func createLoggerDir() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0777)

		if err != nil {
			panic("Error creating logs directory.")
		}
	}
}

// LogInfo is used to log information about an event
func (log *Logger) LogInfo(message string) {

}
