package model

import (
	"testing"
)

func TestShouldConvertStringTimeToHoursAndMinutes(t *testing.T) {

	minutes := "61"
	convertedTime := ToHoursAndMinutes(minutes)
	expectedTime := "1h1m"

	if convertedTime != expectedTime {
		t.Fatalf("Expected: %v; received: %v", expectedTime, convertedTime)
		t.FailNow()
	}
}
