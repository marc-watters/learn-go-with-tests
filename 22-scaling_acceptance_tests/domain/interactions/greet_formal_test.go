package go_specs_greet_test

import (
	"go-specs-greet/v2/specifications"
	"testing"

	interactions "go-specs-greet/v2/domain/interactions"
)

func TestIntroduce(t *testing.T) {
	specifications.IntroduceSpecification(
		t,
		specifications.IntroduceAdapter(interactions.Introduce),
	)
}
