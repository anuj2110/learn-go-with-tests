package mocking

import (
	"fmt"
	"io"
	"os"
)

const (
	finalWord = "Go!"
	countdown = 3
)

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdown; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
