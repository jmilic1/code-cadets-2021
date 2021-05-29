package main

import (
	"flag"
	"github.com/pkg/errors"
	"log"

	"code-cadets-2021/lecture_1/07_testing/filter"
)

func main() {
	var countFrom, countTo, divisor int

	flag.IntVar(&countFrom, "count-from", 1, "Value (inclusive) from which to start counting")
	flag.IntVar(&countTo, "count-to", 10, "Value (inclusive) to count to")
	flag.IntVar(&divisor, "divisor", 1, "The divisor used for filtering")

	flag.Parse()

	numbers, err := filter.GetDivisibleFromRange(countFrom, countTo, divisor)
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error while calculating tax"))
	}

	log.Printf("%v", numbers)
}
