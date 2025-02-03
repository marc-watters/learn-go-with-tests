package collections

import (
	"slices"
	"strings"
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

func TestBadBank(t *testing.T) {
	var (
		marc  = Account{Name: "Marc", Balance: 100}
		mark  = Account{Name: "Mark", Balance: 75}
		marco = Account{Name: "Marco", Balance: 200}

		transactions = []Transaction{
			NewTransaction(marc, mark, 75),
			NewTransaction(marco, marc, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(marc), float64(50))
	AssertEqual(t, newBalanceFor(mark), float64(150))
	AssertEqual(t, newBalanceFor(marco), float64(175))
}

type Person struct{ Name string }

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			{Name: "Marc Watters"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Marc")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Marc Watters"})
	})
}

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot: \t%v\nwant:\t%v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("\ngot: \t%v\nwant: true", got)
	}
}
