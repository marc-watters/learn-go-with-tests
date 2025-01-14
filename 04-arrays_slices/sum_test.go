package collections

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("\ngot: \t%d\nwant:\t%d", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{6, 15}

	if !slices.Equal(got, want) {
		t.Errorf("\ngot: \t%d\nwant:\t%d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum tails of varying slices", func(t *testing.T) {
		got := SumAllTails(
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		)
		want := []int{5, 11, 17}

		if !slices.Equal(got, want) {
			t.Errorf("\ngot: \t%d\nwant:\t%d", got, want)
		}
	})

	t.Run("with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}

		if !slices.Equal(got, want) {
			t.Errorf("\ngot: \t%d\nwant:\t%d", got, want)
		}
	})
}
