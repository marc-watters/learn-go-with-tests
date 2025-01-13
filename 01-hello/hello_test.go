package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello by name", func(t *testing.T) {
		got := Hello("Marc")
		want := "Hello, Marc!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("use default name if passed empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot: \t%q\nwant:\t%q", got, want)
	}
}
