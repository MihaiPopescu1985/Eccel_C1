package controller

import "testing"

func TestToHoursAndMinutes(t *testing.T) {
	got := toHoursAndMinutes("1")
	want := "0:01"
	if want != got {
		t.Error(got + "!=" + want)
	}
	got = toHoursAndMinutes("-1")
	want = "-0:01"
	if want != got {
		t.Error(got + "!=" + want)
	}
	got = toHoursAndMinutes("60")
	want = "1:00"
	if want != got {
		t.Error(got + "!=" + want)
	}
	got = toHoursAndMinutes("-60")
	want = "-1:00"
	if want != got {
		t.Error(got + "!=" + want)
	}
	got = toHoursAndMinutes("61")
	want = "1:01"
	if want != got {
		t.Error(got + "!=" + want)
	}
	got = toHoursAndMinutes("-75")
	want = "-1:15"
	if want != got {
		t.Error(got + "!=" + want)
	}
}
