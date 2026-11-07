// Command stats prints the sum and the mean of the numbers given as arguments.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andrews-agent/gc-playground/stats"
)

func main() {
	var numbers []float64
	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "not a number: %s\n", arg)
			os.Exit(1)
		}
		numbers = append(numbers, n)
	}
	fmt.Printf("sum %g\n", stats.Sum(numbers))
	fmt.Printf("mean %g\n", stats.Mean(numbers))
}
