package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("\ngot: \t%.2f\nwant:\t%.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{5.0, 5.0}
	got := Area(rectangle)
	want := 25.0

	if got != want {
		t.Errorf("\ngot: \t%.2f\nwant:\t%.2f", got, want)
	}
}
