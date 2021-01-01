package service

import (
	"reflect"
	"testing"

	"example.com/c1/service"
)

func TestCommandIfWorkingTimeIsNotStarted(t *testing.T) {

	const deviceName string = "Pepper_C1-1A6318"
	const cardUID = "045D91B22C5E80"

	var dao service.DAO

	want := "INSERT INTO WORKDAY " +
		"(WORKERID, PROJECTID, STARTTIME) " +
		"VALUES ()"
	got := dao.InsertWorkday(deviceName, cardUID)

	if !reflect.DeepEqual(want, got) {
		t.FailNow()
	}
}

func TestCommandWorkingTimeWasStarted(t *testing.T) {
	t.FailNow()
}
