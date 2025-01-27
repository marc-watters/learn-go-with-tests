package racer

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Racer(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for: %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		if _, err := http.Get(url); err != nil {
			fmt.Fprintf(os.Stderr, "error getting URL: %v", err)
		}
		close(ch)
	}()
	return ch
}
