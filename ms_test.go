package ms

import (
	"testing"
)

func TestParseWithDays(t *testing.T) {
	want := float64(172800000)
	got, err := Parse("2 days")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if want != got {
		t.Errorf("Expected %f but got %f", want, got)
	}
}

func TestParseWithDaysShort(t *testing.T) {
	want := float64(172800000)
	got, err := Parse("2 d")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if want != got {
		t.Errorf("Expected %f but got %f", want, got)
	}
}

func TestParseWithHours(t *testing.T) {
	want := float64(9000000)
	got, err := Parse("2.5 hrs")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if want != got {
		t.Errorf("Expected %f but got %f", want, got)
	}
}

func TestParseWithHoursShort(t *testing.T) {
	want := float64(36000000)
	got, err := Parse("10 h")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if want != got {
		t.Errorf("Expected %f but got %f", want, got)
	}
}

func TestParseWithYears(t *testing.T) {
	want := float64(31557600000)
	got, err := Parse("1y")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if want != got {
		t.Errorf("Expected %f but got %f", want, got)
	}
}

func TestParseWithMinutes(t *testing.T) {
	want := float64(60000)
	got, err := Parse("1m")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if want != got {
		t.Errorf("Expected %f but got %f", want, got)
	}
}

func TestParseWithSeconds(t *testing.T) {
	want := float64(5000)
	got, err := Parse("5s")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if want != got {
		t.Errorf("Expected %f but got %f", want, got)
	}
}
