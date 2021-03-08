package util

import (
	"io/ioutil"
	"strconv"
	"testing"
	"time"
)

func TestLoggerShouldCreateANewFileNamedAsCurrentDate(t *testing.T) {
	year, month, day := time.Now().Date()

	filename := "logs/"
	filename += strconv.Itoa(year) + "-" + month.String() + "-" + strconv.Itoa(day) + ".txt"

	//	Log.Init()

	_, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error("error reading file " + filename)
		t.FailNow()
	}
}
