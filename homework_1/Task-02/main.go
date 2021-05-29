// main package contains program executable which uses tax calculator
package main

import (
	"log"

	"github.com/pkg/errors"

	"code-cadets-2021/homework_1/Task-02/calculator"
)

// main entrypoint for demonstrating progressive tax calculator
func main() {

	taxBrackets := []calculator.TaxBracket{
		{
			TaxRate:        0,
			UpperThreshold: 1000,
		},
		{
			TaxRate:        0.1,
			UpperThreshold: 5000,
		},
		{
			TaxRate:        0.2,
			UpperThreshold: 10000,
		},
	}
	finalTaxRate := 0.3
	income := 7000.0

	tax, err := calculator.CalculateProgressiveTax(taxBrackets, finalTaxRate, income)

	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error while calculating tax"),
		)
	}

	log.Println(tax)
}
