package progressive_tax_calculator

import (
	"math"
	"sort"

	"github.com/pkg/errors"
)

var bracketTaxes = make(map[int]float64)
var bracketRanges = make([]int, 0, 1)

func validateInput() error {
	var lastPercent float64
	for _, k := range bracketRanges {
		percent := bracketTaxes[k]

		if percent < lastPercent {
			return errors.New("Tax percentages are not monotonic through class brackets")
		}

		lastPercent = percent
	}

	return nil
}

func Initialize() {
	bracketTaxes = make(map[int]float64)
	bracketRanges = make([]int, 0, 1)
}

func AddTaxRange(upperBound int, percentage float64) {
	if bracketTaxes[upperBound] != 0 {

	}
	bracketTaxes[upperBound] = percentage
}

func Finalize(percentage float64) error {
	bracketTaxes[math.MaxInt32] = percentage

	bracketRanges = make([]int, 0, len(bracketTaxes))
	for k := range bracketTaxes {
		bracketRanges = append(bracketRanges, k)
	}

	sort.Ints(bracketRanges)
	return validateInput()
}

func CalculateProgressiveTax(value int) float64 {
	var tax float64
	var offset int

	for _, k := range bracketRanges {
		percent := bracketTaxes[k]

		factor := math.Min(float64(value-offset), float64(k-offset))
		tax += factor * percent

		offset = k
		if value <= k || k == math.MaxInt32 {
			break
		}
	}

	return tax
}
