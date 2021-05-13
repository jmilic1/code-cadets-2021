package calculator

import (
	"github.com/pkg/errors"
)

// taxBracket models containing the taxRate and the threshold for a certain bracket
type taxBracket struct {
	taxRate   float32
	threshold float32
}

func validateSliceIncreasing(values []float32) error {
	lastThreshold := float32(0)
	for _, threshold := range values {
		if lastThreshold != 0 && lastThreshold >= threshold {
			return errors.New("values are not monotonically increasing")
		}

		lastThreshold = threshold
	}
	return nil
}

func validateBracketInput(thresholds []float32, taxRates []float32, finalTaxRate float32) error {
	if len(thresholds) != len(taxRates) {
		return errors.New("incorrect number of thresholds given the number of tax rates")
	}

	err := validateSliceIncreasing(thresholds)
	if err != nil {
		return err
	}

	err = validateSliceIncreasing(taxRates)
	if err != nil {
		return err
	}

	taxRatesLength := len(taxRates)
	if taxRatesLength == 0 {
		return nil
	}

	if finalTaxRate <= taxRates[taxRatesLength-1] {
		return errors.New("rate of final tax bracket is lesser than the tax rate from another bracket")
	}

	return nil
}

func prepareBracketData(thresholds []float32, taxRates []float32) []taxBracket {
	var brackets []taxBracket
	for index, bracketThreshold := range thresholds {
		bracketTaxRate := taxRates[index]
		bracket := taxBracket{
			threshold: bracketThreshold,
			taxRate:   bracketTaxRate,
		}
		brackets = append(brackets, bracket)
	}

	return brackets
}

func computeTax(brackets []taxBracket, income float32, finalTaxRate float32) float32 {
	var tax float32
	var lastThreshold float32

	for _, bracket := range brackets {
		threshold := bracket.threshold
		rate := bracket.taxRate

		// income does not exceed final bracket
		if income <= threshold {
			incomeWithinBracket := income - lastThreshold
			tax += incomeWithinBracket * rate

			return tax
		}

		incomeWithinBracket := threshold - lastThreshold
		tax += incomeWithinBracket * rate

		lastThreshold = threshold
	}

	incomeWithinBracket := income - lastThreshold
	tax += incomeWithinBracket * finalTaxRate

	return tax
}

// CalculateProgressiveTax calculate tax based on given input.
// given threshold and taxRate slices need to be of the same length and sorted in increasing order.
// finalTaxRate needs to be greater than any other element in taxRates
func CalculateProgressiveTax(thresholds []float32, taxRates []float32, finalTaxRate float32, income float32) (float32, error) {
	err := validateBracketInput(thresholds, taxRates, finalTaxRate)
	if err != nil {
		return 0, err
	}

	brackets := prepareBracketData(thresholds, taxRates)

	return computeTax(brackets, income, finalTaxRate), nil
}
