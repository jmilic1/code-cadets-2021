// main package contains program executable which uses tax calculator
package main

import (
	"log"

	"github.com/pkg/errors"

	"code-cadets-2021/homework_1/Task-02/calculator"
)

// main entrypoint for demonstrating progressive tax calculator
func main() {

	thresholds := []float32{1000, 5000, 10000}
	taxRates := []float32{0, 0.1, 0.2, 0.3}
	tax, err := calculator.CalculateProgressiveTax(thresholds, taxRates, 7000)

	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error while calculating tax"),
		)
	}

	log.Println(tax)
}
