package go_specs_greet_test

import (
	interactions "go-specs-greet/v2/domain/interactions"
	"go-specs-greet/v2/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)
}
