package calculator_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	. "code-cadets-2021/homework_1/Task-02/calculator"
)

// tests progressive tax computation
func TestProgressiveTax(t *testing.T) {
	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			actualOutput, actualErr := CalculateProgressiveTax(tc.thresholds, tc.taxRates, tc.finalTaxRate, tc.income)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}
		})
	}
}
