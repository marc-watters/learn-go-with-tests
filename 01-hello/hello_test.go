package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello by name", func(t *testing.T) {
		got := Hello("Marc")
		want := "Hello, Marc!"

		if got != want {
			t.Errorf("\ngot: \t%q\nwant:\t%q", got, want)
		}
	})

	t.Run("use default name if passed empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("\ngot: \t%q\nwant:\t%q", got, want)
		}
	})
}
