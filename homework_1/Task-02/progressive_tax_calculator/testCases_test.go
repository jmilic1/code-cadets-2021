// Package contains unit tests and their test cases
package progressive_tax_calculator_test

// input model for defining a tax bracket
type taxBracketInput struct {
	taxBracketUpperBound int
	percentage           float64
}

// test case model for unit testing
type testCase struct {
	bracketInputs   []taxBracketInput
	finalPercentage float64
	income          int

	expectedOutput float64
	expectingError bool
}

// getTestCases returns an array of test cases.
func getTestCases() []testCase {
	return []testCase{
		{
			bracketInputs: []taxBracketInput{
				{
					taxBracketUpperBound: 1000,
					percentage:           0,
				},
				{
					taxBracketUpperBound: 5000,
					percentage:           0.1,
				},
				{
					taxBracketUpperBound: 10000,
					percentage:           0.2,
				},
			},
			finalPercentage: 0.3,
			income:          7000,

			expectedOutput: 800,
			expectingError: false,
		},
		{
			bracketInputs: []taxBracketInput{
				{
					taxBracketUpperBound: 10000,
					percentage:           0.1,
				},
				{
					taxBracketUpperBound: 20000,
					percentage:           0.2,
				},
			},
			finalPercentage: 0.3,
			income:          20000,

			expectedOutput: 3000,
			expectingError: false,
		},
		{
			bracketInputs: []taxBracketInput{
				{
					taxBracketUpperBound: 10000,
					percentage:           0.1,
				},
				{
					taxBracketUpperBound: 20000,
					percentage:           0.2,
				},
			},
			finalPercentage: 0.3,
			income:          25000,

			expectedOutput: 4500,
			expectingError: false,
		},
		{
			bracketInputs: []taxBracketInput{
				{
					taxBracketUpperBound: 10000,
					percentage:           0.1,
				},
				{
					taxBracketUpperBound: 20000,
					percentage:           0.15,
				},
				{
					taxBracketUpperBound: 20000,
					percentage:           0.15,
				},
				{
					taxBracketUpperBound: 30000,
					percentage:           0.2,
				},
				{
					taxBracketUpperBound: 40000,
					percentage:           0.25,
				},
				{
					taxBracketUpperBound: 50000,
					percentage:           0.3,
				},
			},
			finalPercentage: 0.35,
			income:          70000,

			expectedOutput: 17000,
			expectingError: false,
		},
		{
			finalPercentage: 0,
			income:          70000,

			expectedOutput: 0,
			expectingError: false,
		},
		{
			bracketInputs: []taxBracketInput{
				{
					taxBracketUpperBound: 1000,
					percentage:           0.1,
				},
			},
			finalPercentage: 0.2,
			income:          7000,

			expectedOutput: 1300,
			expectingError: false,
		},
		{
			bracketInputs: []taxBracketInput{
				{
					taxBracketUpperBound: 1000,
					percentage:           0.1,
				},
			},
			finalPercentage: 0.2,
			income:          0,

			expectingError: false,
		},
		{
			bracketInputs: []taxBracketInput{
				{
					taxBracketUpperBound: 1000,
					percentage:           0.1,
				},
			},
			finalPercentage: 0,
			income:          0,

			expectingError: true,
		},
		{
			bracketInputs: []taxBracketInput{
				{
					taxBracketUpperBound: 1000,
					percentage:           0.3,
				},
				{
					taxBracketUpperBound: 2000,
					percentage:           0.2,
				},
			},
			finalPercentage: 0.4,
			income:          0,

			expectingError: true,
		},

		{
			bracketInputs: []taxBracketInput{
				{
					taxBracketUpperBound: 1000,
					percentage:           0.3,
				},
				{
					taxBracketUpperBound: 1000,
					percentage:           0.4,
				},
			},
			finalPercentage: 0.5,
			income:          2000,

			expectedOutput: 900,
			expectingError: false,
		},

		{
			bracketInputs: []taxBracketInput{
				{
					taxBracketUpperBound: 1000,
					percentage:           0.3,
				},
				{
					taxBracketUpperBound: 1000,
					percentage:           0.2,
				},
			},
			finalPercentage: 0.1,
			income:          0,

			expectingError: true,
		},
	}
}
