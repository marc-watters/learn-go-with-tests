package racer

import (
	"fmt"
	"net/http"
	"os"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
