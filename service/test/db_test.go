package service

import (
	"fmt"
	"reflect"
	"testing"

	"example.com/c1/service"
)

func TestDbConnection(t *testing.T) {

	// TODO: mock a database connection
	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}
}

func TestDeviceTableColumns(t *testing.T) {

	// TODO: mock a database connection
	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	const command string = "SELECT * FROM DEVICE WHERE ID=0;"
	columnsNames := []string{"ID", "NAME", "IP", "ISENDPOINT"}

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

	// TODO: mock a database connection
	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	const command string = "SELECT * FROM POSITION WHERE ID=0;"
	columnsNames := []string{"ID", "POSITION"}

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

	// TODO: mock a database connection
	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	const command string = "SELECT * FROM PROJECT WHERE ID=0;"
	columnsNames := []string{"ID", "GENUMBER", "RONUMBER", "DESCRIPTION", "DEVICEID", "ACTIVE"}

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

	// TODO: mock a database connection
	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	const command string = "SELECT * FROM WORKDAY WHERE ID=0;"
	columnsNames := []string{"ID", "WORKERID", "PROJECTID", "STARTTIME", "STOPTIME"}

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

	// TODO: mock a database connection
	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	const command string = "SELECT * FROM WORKER WHERE ID=0;"
	columnsNames := []string{"ID", "FIRSTNAME", "LASTNAME", "CARDNUMBER",
		"POSITIONID", "ISACTIVE", "NICKNAME",
		"PASSWORD", "ACCESSLEVEL"}

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

func TestRetrieveActiveWorkday(t *testing.T) {

	// TODO: mock a database connection
	var dao service.DAO
	dao.Connect()

	if !dao.IsConnected() {
		fmt.Println("Error connecting to database")
		t.FailNow()
	}

	const command string = "SELECT * FROM ACTIVEWORKDAY WHERE ID=0;"
	columnsNames := []string{"ID", "WORKER", "RO_NUMBER",
		"GE_NUMBER", "PROJ_DESCRIPTION"}

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
