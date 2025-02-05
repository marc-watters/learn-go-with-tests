package go_specs_greet_test

import (
	go_specs_greet "go-specs-greet/v2"
	"go-specs-greet/v2/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(go_specs_greet.Greet),
	)
}
