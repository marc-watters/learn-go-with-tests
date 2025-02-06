package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type FormalGreeter interface {
	Introduce(name string) (string, error)
}

func IntroduceSpecification(t *testing.T, executive FormalGreeter) {
	got, err := executive.Introduce("Marc")
	assert.NoError(t, err)
	assert.Equal(t, got, "An honor to meet you, Marc.")
}
