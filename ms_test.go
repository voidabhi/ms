package ms

import (
	"testing"
)

func TestParseWithDays(t *testing.T) {
	want := float64(172800000)
	got, err := Parse("2 days")
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	if want != got {
		t.Errorf("Expected %s but got %s", want, got)
	}
}
