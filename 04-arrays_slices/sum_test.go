package collections

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("\ngot: \t%d\nwant:\t%d", got, want)
	}
}
