// B"H

/*
average2 calculates the average of several numbers.

Note, this is a command line tool. So `go install` it first.

➜  06 git:(master) ✗ cd average2

➜  average2 git:(master) ✗ go install

➜  average2 git:(master) ✗ average2 12 23 34
Average: 23.00

➜  average2 git:(master) ✗ average2 12 23 34 45 99 9999999
Average: 1666702.00

*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// -- ------------------------------------------
// Note this is a variadic function.
func average(numbers ...float64) float64 {
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}

// -- ------------------------------------------

func main() {

	// Get the command-line args:
	arguments := os.Args[1:]

	var numbers []float64

	// -- --------------------------------------
	for _, argument := range arguments {

		number, err := strconv.ParseFloat(argument, 64)

		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}
	// -- --------------------------------------

	// Note how we add `...` to convert the slice into "filling" the variadic parameter:
	fmt.Printf("Average: %0.2f\n", average(numbers...))
}
