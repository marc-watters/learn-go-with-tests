package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("\nrepeated:\t%q\nexpected:\t%q", repeated, expected)
	}
}
