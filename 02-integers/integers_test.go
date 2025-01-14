package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("\ngot: \t%d\nwant:\t%d", sum, expected)
	}
}
