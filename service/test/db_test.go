package service

import (
	"fmt"
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

func TestDeviceTableColumns(t *testing.T) {

	const command string = "SELECT * FROM DEVICE WHERE ID=0;"
	columnsNames := []string{"ID", "NAME", "IP", "ISENDPOINT"}

	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	deviceColumns, err := dao.ExecuteQuery(command).Columns()

	if err != nil {
		fmt.Println("Error executing query." + err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(columnsNames, deviceColumns) {
		fmt.Println("Columns are not named as expected.")
		fmt.Println(columnsNames)
		fmt.Println(deviceColumns)
		t.FailNow()
	}
}

func TestPositionTableColumns(t *testing.T) {

	const command string = "SELECT * FROM POSITION WHERE ID=0;"
	columnsNames := []string{"ID", "POSITION"}

	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	deviceColumns, err := dao.ExecuteQuery(command).Columns()

	if err != nil {
		fmt.Println("Error executing query." + err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(columnsNames, deviceColumns) {
		fmt.Println("Columns are not named as expected.")
		fmt.Println(columnsNames)
		fmt.Println(deviceColumns)
		t.FailNow()
	}
}

func TestProjectTableColumns(t *testing.T) {

	const command string = "SELECT * FROM PROJECT WHERE ID=0;"
	columnsNames := []string{"ID", "GENUMBER", "RONUMBER", "DESCRIPTION", "DEVICEID", "ACTIVE"}

	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	deviceColumns, err := dao.ExecuteQuery(command).Columns()

	if err != nil {
		fmt.Println("Error executing query." + err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(columnsNames, deviceColumns) {
		fmt.Println("Columns are not named as expected.")
		fmt.Println(columnsNames)
		fmt.Println(deviceColumns)
		t.FailNow()
	}
}

func TestWorkdayTableColumns(t *testing.T) {

	const command string = "SELECT * FROM WORKDAY WHERE ID=0;"
	columnsNames := []string{"ID", "WORKERID", "PROJECTID", "STARTTIME", "STOPTIME"}

	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	deviceColumns, err := dao.ExecuteQuery(command).Columns()

	if err != nil {
		fmt.Println("Error executing query." + err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(columnsNames, deviceColumns) {
		fmt.Println("Columns are not named as expected.")
		fmt.Println(columnsNames)
		fmt.Println(deviceColumns)
		t.FailNow()
	}
}

func TestWorkerTableColumns(t *testing.T) {

	const command string = "SELECT * FROM WORKER WHERE ID=0;"
	columnsNames := []string{"ID", "FIRSTNAME", "LASTNAME", "CARDNUMBER",
		"POSITIONID", "ISACTIVE", "NICKNAME",
		"PASSWORD", "ACCESSLEVEL"}

	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	deviceColumns, err := dao.ExecuteQuery(command).Columns()

	if err != nil {
		fmt.Println("Error executing query." + err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(columnsNames, deviceColumns) {
		fmt.Println("Columns are not named as expected.")
		fmt.Println(columnsNames)
		fmt.Println(deviceColumns)
		t.FailNow()
	}
}
