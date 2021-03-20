package controller

import (
	"net/url"
	"testing"
	"time"
)

func TestShouldGenerateProperStandardViewOfMonthTimeReport(t *testing.T) {

	layout := "2006-1-2 15:04:05"

	for i := 0; i < 366; i++ {
		str := "2021-12-31 12:33:34"

		tm, err := time.Parse(layout, str)

		if err != nil || tm.Day() != 31 {
			t.Fatal(err)
			t.FailNow()
		}
	}
}

func TestShouldParseURL(t *testing.T) {

	want := "4"
	url, err := url.Parse("/stage-one?workerId=4")

	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}

	got := url.Query().Get("workerId")

	if got != want {
		t.Fatalf("got = %v, wanted = %v", got, want)
		t.FailNow()
	}
}
