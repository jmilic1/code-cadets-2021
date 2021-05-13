package calculator_test

import "code-cadets-2021/homework_1/Task-02/calculator"

// test case model for unit testing
type testCase struct {
	taxBrackets  []calculator.TaxBracket
	finalTaxRate float64
	income       float64

	expectedOutput float64
	expectingError bool
}

// getTestCases returns an array of test cases.
func getTestCases() []testCase {
	return []testCase{
		{
			taxBrackets: []calculator.TaxBracket{
				{TaxRate: 0, Threshold: 1000},
				{TaxRate: 0.1, Threshold: 5000},
				{TaxRate: 0.2, Threshold: 10000},
			},
			finalTaxRate: 0.3,
			income:       7000,

			expectedOutput: 800,
			expectingError: false,
		},
		{
			taxBrackets: []calculator.TaxBracket{
				{TaxRate: 0.1, Threshold: 10000},
				{TaxRate: 0.2, Threshold: 20000},
			},
			finalTaxRate: 0.3,
			income:       20000,

			expectedOutput: 3000,
			expectingError: false,
		},
		{
			taxBrackets: []calculator.TaxBracket{
				{TaxRate: 0.1, Threshold: 10000},
				{TaxRate: 0.2, Threshold: 20000},
			},
			finalTaxRate: 0.3,
			income:       25000,

			expectedOutput: 4500,
			expectingError: false,
		},
		{
			taxBrackets: []calculator.TaxBracket{
				{TaxRate: 0.1, Threshold: 10000},
				{TaxRate: 0.15, Threshold: 20000},
				{TaxRate: 0.2, Threshold: 30000},
				{TaxRate: 0.25, Threshold: 40000},
				{TaxRate: 0.3, Threshold: 50000},
			},
			finalTaxRate: 0.35,
			income:       70000,

			expectedOutput: 17000,
			expectingError: false,
		},
		{
			taxBrackets: []calculator.TaxBracket{
				{},
			},
			finalTaxRate: 0.1,
			income:       70000,

			expectedOutput: 7000,
			expectingError: false,
		},
		{
			taxBrackets: []calculator.TaxBracket{
				{TaxRate: 0.1, Threshold: 1000},
			},
			finalTaxRate: 0.2,
			income:       7000,

			expectedOutput: 1300,
			expectingError: false,
		},
		{
			taxBrackets: []calculator.TaxBracket{
				{TaxRate: 0.1, Threshold: 1000},
			},
			finalTaxRate: 0.2,
			income:       0,

			expectingError: false,
		},
		{
			taxBrackets: []calculator.TaxBracket{
				{TaxRate: 0.1, Threshold: 1000},
			},

			expectingError: true,
		},
	}
}
