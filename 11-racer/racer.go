package racer

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	if _, err := http.Get(url); err != nil {
		fmt.Fprintf(os.Stderr, "error getting url: %v", err)
		os.Exit(1)
	}
	return time.Since(start)
}
