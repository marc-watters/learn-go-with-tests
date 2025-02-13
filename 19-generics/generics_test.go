package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Marc")
	})
}

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot: \t%v\nwant:\t%v", got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want any) {
	t.Helper()

	if got == want {
		t.Errorf("\ngot: \t%v\nwant:\t%v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("\ngot: \t%v\nwant: true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("\ngot: \t%v\nwant: false", got)
	}
}

func TestStack(t *testing.T) {
	t.Run("generic stack", func(t *testing.T) {
		myStackOfInts := NewStack[int]()

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a ting, then check if it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())

		// can get the numbers we put in as numbesr, not untyped interface{}
		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
}
