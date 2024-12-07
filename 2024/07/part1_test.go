package main

import "testing"

var part1Sample = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`
var part1Expected = 3749

func TestPart1(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected int
	}

	testCases := []testCase{
		{
			name:     "Part 1 Sample",
			input:    part1Sample,
			expected: part1Expected,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := part1(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
