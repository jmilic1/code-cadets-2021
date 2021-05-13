package calculator

import (
	"github.com/pkg/errors"
)

// TaxBracket models containing the TaxRate and the Threshold for a certain bracket
type TaxBracket struct {
	TaxRate   float64
	Threshold float64
}

func validateBracketsIncreasing(taxBrackets []TaxBracket) error {
	lastThreshold := 0.0
	lastTaxRate := 0.0

	for _, taxBracket := range taxBrackets {
		threshold := taxBracket.Threshold
		taxRate := taxBracket.TaxRate

		if lastThreshold != 0 && lastThreshold >= threshold {
			return errors.New("values are not monotonically increasing")
		}
		if lastTaxRate != 0 && lastTaxRate >= taxRate {
			return errors.New("values are not monotonically increasing")
		}

		lastThreshold = threshold
		lastTaxRate = taxRate
	}
	return nil
}

func validateBracketInput(taxBrackets []TaxBracket, finalTaxRate float64) error {
	err := validateBracketsIncreasing(taxBrackets)
	if err != nil {
		return err
	}

	taxRatesLength := len(taxBrackets)
	if taxRatesLength != 0 && finalTaxRate <= taxBrackets[taxRatesLength-1].TaxRate {
		return errors.New("rate of final tax bracket is lesser than the tax rate from another bracket")
	}

	return nil
}

func computeTax(brackets []TaxBracket, income float64, finalTaxRate float64) float64 {
	var tax float64
	var lastThreshold float64

	for _, bracket := range brackets {
		threshold := bracket.Threshold
		rate := bracket.TaxRate

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
// given taxBrackets need to be sorted in increasing order based on their taxRates and thresholds.
// finalTaxRate needs to be greater than any taxRate defined in taxBrackets
func CalculateProgressiveTax(taxBrackets []TaxBracket, finalTaxRate float64, income float64) (float64, error) {
	err := validateBracketInput(taxBrackets, finalTaxRate)
	if err != nil {
		return 0, err
	}

	return computeTax(taxBrackets, income, finalTaxRate), nil
}
