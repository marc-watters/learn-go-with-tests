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
