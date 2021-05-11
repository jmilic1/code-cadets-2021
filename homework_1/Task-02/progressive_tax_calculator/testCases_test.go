package progressive_tax_calculator_test

type taxClassInput struct {
	taxClassUpperBound int
	percentage         float64
}

type testCase struct {
	classInput      []taxClassInput
	finalPercentage float64
	valueInput      int

	expectedOutput float64
	expectingError bool
}

// getTestCases returns an array of test cases.
func getTestCases() []testCase {
	return []testCase{
		{
			classInput: []taxClassInput{
				{
					taxClassUpperBound: 1000,
					percentage:         0,
				},
				{
					taxClassUpperBound: 5000,
					percentage:         0.1,
				},
				{
					taxClassUpperBound: 10000,
					percentage:         0.2,
				},
			},
			finalPercentage: 0.3,
			valueInput:      7000,

			expectedOutput: 800,
			expectingError: false,
		},
		{
			classInput: []taxClassInput{
				{
					taxClassUpperBound: 10000,
					percentage:         0.1,
				},
				{
					taxClassUpperBound: 20000,
					percentage:         0.2,
				},
			},
			finalPercentage: 0.3,
			valueInput:      20000,

			expectedOutput: 3000,
			expectingError: false,
		},
		{
			classInput: []taxClassInput{
				{
					taxClassUpperBound: 10000,
					percentage:         0.1,
				},
				{
					taxClassUpperBound: 20000,
					percentage:         0.2,
				},
			},
			finalPercentage: 0.3,
			valueInput:      25000,

			expectedOutput: 4500,
			expectingError: false,
		},
		{
			classInput: []taxClassInput{
				{
					taxClassUpperBound: 10000,
					percentage:         0.1,
				},
				{
					taxClassUpperBound: 20000,
					percentage:         0.15,
				},
				{
					taxClassUpperBound: 20000,
					percentage:         0.15,
				},
				{
					taxClassUpperBound: 30000,
					percentage:         0.2,
				},
				{
					taxClassUpperBound: 40000,
					percentage:         0.25,
				},
				{
					taxClassUpperBound: 50000,
					percentage:         0.3,
				},
			},
			finalPercentage: 0.35,
			valueInput:      70000,

			expectedOutput: 17000,
			expectingError: false,
		},
		{
			finalPercentage: 0,
			valueInput:      70000,

			expectedOutput: 0,
			expectingError: false,
		},
		{
			classInput: []taxClassInput{
				{
					taxClassUpperBound: 1000,
					percentage:         0.1,
				},
			},
			finalPercentage: 0.2,
			valueInput:      7000,

			expectedOutput: 1300,
			expectingError: false,
		},
		{
			classInput: []taxClassInput{
				{
					taxClassUpperBound: 1000,
					percentage:         0.1,
				},
			},
			finalPercentage: 0.2,
			valueInput:      0,

			expectingError: false,
		},
		{
			classInput: []taxClassInput{
				{
					taxClassUpperBound: 1000,
					percentage:         0.1,
				},
			},
			finalPercentage: 0,
			valueInput:      0,

			expectingError: true,
		},
		{
			classInput: []taxClassInput{
				{
					taxClassUpperBound: 1000,
					percentage:         0.3,
				},
				{
					taxClassUpperBound: 2000,
					percentage:         0.2,
				},
			},
			finalPercentage: 0.4,
			valueInput:      0,

			expectingError: true,
		},

		{
			classInput: []taxClassInput{
				{
					taxClassUpperBound: 1000,
					percentage:         0.3,
				},
				{
					taxClassUpperBound: 1000,
					percentage:         0.4,
				},
			},
			finalPercentage: 0.5,
			valueInput:      2000,

			expectedOutput: 900,
			expectingError: false,
		},

		{
			classInput: []taxClassInput{
				{
					taxClassUpperBound: 1000,
					percentage:         0.3,
				},
				{
					taxClassUpperBound: 1000,
					percentage:         0.2,
				},
			},
			finalPercentage: 0.1,
			valueInput:      0,

			expectingError: true,
		},
	}
}
