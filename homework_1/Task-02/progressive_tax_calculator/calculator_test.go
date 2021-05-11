package progressive_tax_calculator_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	. "code-cadets-2021/homework_1/Task-02/progressive_tax_calculator"
)

func TestProgressiveTax(t *testing.T) {
	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			Initialize()
			for _, val := range tc.classInput {
				AddTaxRange(val.taxClassUpperBound, val.percentage)
			}
			actualErr := Finalize(tc.finalPercentage)
			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				actualOutput := CalculateProgressiveTax(tc.valueInput)

				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}
		})
	}
}
