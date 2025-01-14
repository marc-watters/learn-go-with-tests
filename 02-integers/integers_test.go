package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("\ngot: \t%d\nwant:\t%d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(5, 5)
	fmt.Println(sum)
	// Output: 10
}
