package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("\ngot: \t%s\nwant:\t%s\ngiven:\t%s", got, want, "test")
	}
}
