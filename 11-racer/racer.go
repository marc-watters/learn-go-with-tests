package racer

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	if _, err := http.Get(a); err != nil {
		fmt.Fprintf(os.Stderr, "error getting url: %v", err)
		os.Exit(1)
	}
	aDuration := time.Since(startA)

	startB := time.Now()
	if _, err := http.Get(b); err != nil {
		fmt.Fprintf(os.Stderr, "error getting url: %v", err)
		os.Exit(1)
	}
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}
	return b
}
