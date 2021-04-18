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
	got := readSettingsFromFile()

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("error reading settings")
	}
}

func TestShouldSetDriverAndCredentials(t *testing.T) {
	const wantDriver string = "mysql"
	const wantConnSet string = "root:R00tpassword@/EccelC1"

	Db.getDBSettings()

	if wantDriver != Db.settings.driver {
		t.Errorf("driver is wrong: " + Db.settings.driver)
	}

	if wantConnSet != Db.settings.settings {
		t.Errorf("connection settings are wrong: " + Db.settings.settings)
	}
}
