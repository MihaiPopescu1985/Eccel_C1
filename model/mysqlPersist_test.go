package model

import (
	"reflect"
	"testing"
)

func TestShouldGetSettingsFromFile(t *testing.T) {

	db := MysqlDB{}

	want := `{
	"driver":"mysql",
	"user":"root",
	"password":"R00tpassword",
	"URL":"/",
	"name":"EccelC1"
}`
	const file string = "settings.json"
	db.Init(file)

	got, _ := db.readSettingsFromFile()

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("error reading settings")
	}
}

func TestShouldSetDriverAndCredentials(t *testing.T) {

	db := MysqlDB{}

	const wantDriver string = "mysql"
	const wantConnSet string = "root:R00tpassword@/EccelC1"
	const file string = "settings.json"

	db.Init(file)
	db.getDBSettings()

	if wantDriver != db.settings.driver {
		t.Errorf("driver is wrong: " + db.settings.driver)
	}

	if wantConnSet != db.settings.settings {
		t.Errorf("connection settings are wrong: " + db.settings.settings)
	}
}

func TestShouldInitDbWithStandardFile(t *testing.T) {

	db := MysqlDB{}
	const file string = "settings.json"

	db.Init(file)

	if db.settingsFile != file {
		t.Fatalf("standard file not loaded")
	}
}

func TestShouldConnectToDatabase(t *testing.T) {

	db := MysqlDB{}
	const (
		settingsFile string = "testSettings.json"
	)
	db.Init(settingsFile)
	db.Connect()

	if err := db.IsConnected(); err != nil {
		t.Fatalf("no database connection could be established")
	}
}
