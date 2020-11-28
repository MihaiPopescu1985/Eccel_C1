package test/c1device

import (
	"reflect"
	"testing"
)

func TestGetCommand(t *testing.T) {

	var input []byte = []byte{0x01}
	var want []byte = []byte{0xF5, 0x03, 0x00, 0xFC, 0xFF, 0x01, 0xD1, 0xF1}
	var output []byte = GetCommand(input)

	if !reflect.DeepEqual(want, output) {
		t.Error("Desired output differs from expected.")
		t.FailNow()
	}
}
