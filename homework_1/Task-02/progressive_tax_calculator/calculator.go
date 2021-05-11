// Package progressive_tax_calculator contains progressive tax calculator
package progressive_tax_calculator

import (
	"math"
	"sort"

	"github.com/pkg/errors"
)

// Private map property containing upper bound of brackets and their corresponding tax rate.
var bracketTaxes = make(map[int]float64)

// Private array property containing upper bounds of brackets.
// Used as a utility collection to sort brackets by their upper bounds in increasing order.
var bracketRanges = make([]int, 0, 1)

// validateInput validates if bracketTaxes holds valid progressive tax data.
// returns error if data are not valid
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

// ResetCalc deletes cached tax brackets
func ResetCalc() {
	bracketTaxes = make(map[int]float64)
	bracketRanges = make([]int, 0, 1)
}

// AddTaxRange adds a new bracket or overwrites an old bracket in bracketTaxes
func AddTaxRange(upperBound int, percentage float64) {
	bracketTaxes[upperBound] = percentage
}

// Finalize adds the given percentage as the final tax bracket and validates if tax brackets were properly defined.
// throws error if tax brackets are not valid
func Finalize(percentage float64) error {
	bracketTaxes[math.MaxInt32] = percentage

	bracketRanges = make([]int, 0, len(bracketTaxes))
	for k := range bracketTaxes {
		bracketRanges = append(bracketRanges, k)
	}

	sort.Ints(bracketRanges)
	return validateInput()
}

// CalculateProgressiveTax computes progressive tax for given income
func CalculateProgressiveTax(income int) float64 {
	var tax float64
	var offset int

	for _, k := range bracketRanges {
		percent := bracketTaxes[k]

		factor := math.Min(float64(income-offset), float64(k-offset))
		tax += factor * percent

		offset = k
		if income <= k || k == math.MaxInt32 {
			break
		}
	}

	return tax
}
