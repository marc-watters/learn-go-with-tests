package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countDownStart = 3
	finalWord      = "Go!"
)

func Countdown(out io.Writer) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
