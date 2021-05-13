package calculator_test

// test case model for unit testing
type testCase struct {
	thresholds []float32
	taxRates   []float32
	income     float32

	expectedOutput float32
	expectingError bool
}

// getTestCases returns an array of test cases.
func getTestCases() []testCase {
	return []testCase{
		{
			thresholds: []float32{1000, 5000, 10000},
			taxRates:   []float32{0, 0.1, 0.2, 0.3},
			income:     7000,

			expectedOutput: 800,
			expectingError: false,
		},
		{
			thresholds: []float32{10000, 20000},
			taxRates:   []float32{0.1, 0.2, 0.3},
			income:     20000,

			expectedOutput: 3000,
			expectingError: false,
		},
		{
			thresholds: []float32{10000, 20000},
			taxRates:   []float32{0.1, 0.2, 0.3},
			income:     25000,

			expectedOutput: 4500,
			expectingError: false,
		},
		{
			thresholds: []float32{10000, 20000, 30000, 40000, 50000},
			taxRates:   []float32{0.1, 0.15, 0.2, 0.25, 0.3, 0.35},
			income:     70000,

			expectedOutput: 17000,
			expectingError: false,
		},
		{
			thresholds: []float32{},
			taxRates:   []float32{0.1},
			income:     70000,

			expectedOutput: 7000,
			expectingError: false,
		},
		{
			thresholds: []float32{1000},
			taxRates:   []float32{0.1, 0.2},
			income:     7000,

			expectedOutput: 1300,
			expectingError: false,
		},
		{
			thresholds: []float32{1000},
			taxRates:   []float32{0.1, 0.2},
			income:     0,

			expectingError: false,
		},
		{
			thresholds: []float32{1000},
			taxRates:   []float32{0.1},

			expectingError: true,
		},
		{
			thresholds: []float32{1000},
			taxRates:   []float32{0.1, 0},

			expectingError: true,
		},
		{
			thresholds: []float32{1000, 1000},
			taxRates:   []float32{0.3, 0.4, 0.5},

			expectingError: true,
		},
		{
			thresholds: []float32{10000, 20000, 20000},
			taxRates:   []float32{0.1, 0.15, 0.15, 0.2},

			expectingError: true,
		},
	}
}
