package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("\ngot: \t%.2f\nwant:\t%.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.0, 5.0}
		got := rectangle.Area()
		want := 25.0

		if got != want {
			t.Errorf("\ngot: \t%g\nwant:\t%g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("\ngot: \t%g\nwant:\t%g", got, want)
		}
	})
}
