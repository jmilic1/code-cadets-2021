package progressive_tax_calculator

import (
	"math"
	"sort"

	"github.com/pkg/errors"
)

var classTaxes = make(map[int]float64)
var classRanges = make([]int, 0, 1)

func validatePercentages() error {
	var lastPercent float64
	for _, k := range classRanges {
		percent := classTaxes[k]

		if lastPercent == 0 {
			lastPercent = percent
		} else {
			if percent < lastPercent {
				return errors.New("Percentages are not in order of class ranges")
			}
		}
	}

	return nil
}

func Initialize() {
	classTaxes = make(map[int]float64)
	classRanges = make([]int, 0, 1)
}

func AddTaxRange(upperBound int, percentage float64) {
	classTaxes[upperBound] = percentage
}

func Finalize(percentage float64) error {
	classTaxes[math.MaxInt32] = percentage

	classRanges = make([]int, 0, len(classTaxes))
	for k := range classTaxes {
		classRanges = append(classRanges, k)
	}

	sort.Ints(classRanges)
	return validatePercentages()
}

func CalculateProgressiveTax(value int) float64 {
	var tax float64
	var offset int

	for _, k := range classRanges {
		percent := classTaxes[k]

		temp := k - offset

		if k == math.MaxInt32 {
			tax += float64(value-offset) * percent
			break
		}

		if value <= k {
			tax += float64(value-temp) * percent
			break
		}

		tax += float64(temp) * percent
		offset = k
	}

	return tax
}
