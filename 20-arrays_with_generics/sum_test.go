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
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("\ngot: \t%d\nwant:\t%d", got, want)
		}
	}

	t.Run("sum tails of varying slices", func(t *testing.T) {
		got := SumAllTails(
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		)
		want := []int{5, 11, 17}

		checkSums(t, got, want)
	})

	t.Run("with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}

		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})
}

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot: \t%v\nwant:\t%v", got, want)
	}
}
