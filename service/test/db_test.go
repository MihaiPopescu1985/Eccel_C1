package service

import (
	"reflect"
	"testing"

	"example.com/c1/service"
)

func TestDbConnection(t *testing.T) {
	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		t.FailNow()
	}
}

func TestStartAndStopTimeOnWorkdayCommand(t *testing.T) {

	const deviceName string = "Pepper_C1-1A6318"
	const cardUID = "045D91B22C5E80"

	var dao service.DAO

	want := "CALL INSERT_INTO_WORKDAY(\"Pepper_C1-1A6318\", \"045D91B22C5E80\");"
	got := dao.InsertIntoWorkday(deviceName, cardUID)

	if !reflect.DeepEqual(want, got) {
		t.FailNow()
	}
}

func TestExecuteInsertIntoWorkday(t *testing.T) {
	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		t.FailNow()
	}

	const deviceName string = "Pepper_C1-1A6318"
	const cardUID = "045D91B22C5E80"

	want := "CALL INSERT_INTO_WORKDAY(\"Pepper_C1-1A6318\", \"045D91B22C5E80\");"
	got := dao.InsertIntoWorkday(deviceName, cardUID)

	if !reflect.DeepEqual(want, got) {
		t.FailNow()
	}

	dao.Execute(got)
}
