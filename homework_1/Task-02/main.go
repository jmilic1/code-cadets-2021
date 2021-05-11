package main

import (
	"log"

	. "code-cadets-2021/homework_1/Task-02/progressive_tax_calculator"
)

func main() {
	AddTaxRange(1000, 0)
	AddTaxRange(5000, 0.1)
	AddTaxRange(10000, 0.2)
	err := Finalize(0.3)

	if err != nil {
		log.Fatal(err, "Error while calculating tax!")
	}

	num := CalculateProgressiveTax(7000)
	log.Println(num)
}
