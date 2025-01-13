package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Marc")
	want := "Hello, Marc!"

	if got != want {
		t.Errorf("\ngot: \t%q\nwant:\t%q", got, want)
	}
}
