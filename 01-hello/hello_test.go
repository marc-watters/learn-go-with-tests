package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("\ngot: \t%q\nwant:\t%q", got, want)
	}
}
