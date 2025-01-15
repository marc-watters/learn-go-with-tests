package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("\ngot: \t%.2f\nwant:\t%.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(5.0, 5.0)
	want := 25.0

	if got != want {
		t.Errorf("\ngot: \t%.2f\nwant:\t%.2f", got, want)
	}
}
