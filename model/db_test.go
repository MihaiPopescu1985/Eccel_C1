package model

import (
	"reflect"
	"testing"
)

func TestShouldGetSettingsFromFile(t *testing.T) {
	want := `{
	"driver":"mysql",
	"user":"root",
	"password":"R00tpassword",
	"URL":"/",
	"name":"EccelC1"
}`
	const file string = "settings.json"
	Db.Init(file)

	got := Db.readSettingsFromFile()

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("error reading settings")
	}
}

func TestShouldSetDriverAndCredentials(t *testing.T) {
	const wantDriver string = "mysql"
	const wantConnSet string = "root:R00tpassword@/EccelC1"
	const file string = "settings.json"

	Db.Init(file)
	Db.getDBSettings()

	if wantDriver != Db.settings.driver {
		t.Errorf("driver is wrong: " + Db.settings.driver)
	}

	if wantConnSet != Db.settings.settings {
		t.Errorf("connection settings are wrong: " + Db.settings.settings)
	}
}

func TestShouldInitDbWithStandardFile(t *testing.T) {
	const file string = "settings.json"

	Db.Init(file)

	if Db.settingsFile != file {
		t.Fatalf("standard file not loaded")
	}
}

func TestShouldConnectToDatabase(t *testing.T) {
	const (
		settingsFile string = "testSettings.json"
	)
	Db.Init(settingsFile)
	Db.Connect()

	if !Db.IsConnected() {
		t.Fatalf("no database connection could be established")
	}
}
