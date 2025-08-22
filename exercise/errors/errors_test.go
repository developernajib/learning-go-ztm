package timeparse

import "testing"

func TestParseTimeSuccess(t *testing.T) {
	result, err := ParseTime("14:07:33")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Hour != 14 || result.Minute != 7 || result.Second != 33 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestParseTimeInvalidFormat(t *testing.T) {
	_, err := ParseTime("14:07")

	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestParseTimeInvalidHour(t *testing.T) {
	_, err := ParseTime("aa:07:33")

	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
