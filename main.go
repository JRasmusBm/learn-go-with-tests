package main

import (
	"fmt"
	"io"
	"os"
)

const start = 3
const finalWord = "Go!"

func Countdown(out io.Writer) {
	for i := start; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
