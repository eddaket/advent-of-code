package main

import "testing"

var part2Sample = `3   4
4   3
2   5
1   3
3   9
3   3`
var part2Expected = 31

func TestPart2(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	testCases := []testCase{
		{
			name:     "Part 2 Sample",
			input:    part2Sample,
			expected: part2Expected,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := part2(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
