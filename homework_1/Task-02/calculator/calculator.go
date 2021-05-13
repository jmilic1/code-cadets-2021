package calculator

import (
	"math"

	"github.com/pkg/errors"
)

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

func validateBracketInput(thresholds []float32, taxRates []float32) error {
	if len(thresholds)+1 != len(taxRates) {
		return errors.New("incorrect number of thresholds given the number of tax rates")
	}

	err := validateSliceIncreasing(thresholds)
	if err != nil {
		return err
	}

	return validateSliceIncreasing(taxRates)
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

	bracket := taxBracket{
		threshold: math.MaxFloat32,
		taxRate:   taxRates[len(taxRates)-1],
	}
	brackets = append(brackets, bracket)
	return brackets
}

func CalculateProgressiveTax(thresholds []float32, taxRates []float32, income float32) (float32, error) {
	err := validateBracketInput(thresholds, taxRates)
	if err != nil {
		return 0, err
	}

	brackets := prepareBracketData(thresholds, taxRates)

	var tax float32
	var lastThreshold float32

	for _, bracket := range brackets {
		threshold := bracket.threshold
		rate := bracket.taxRate

		incomeWithinBracket := math.Min(float64(income-lastThreshold), float64(threshold-lastThreshold))
		tax += float32(incomeWithinBracket) * rate

		if income <= threshold {
			break
		}

		lastThreshold = threshold
	}

	return tax, nil
}
