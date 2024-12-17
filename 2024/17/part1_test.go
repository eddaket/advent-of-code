package main

import "testing"

var part1Sample = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
var part1Expected = "4,6,3,5,6,3,5,2,1,0"

func TestPart1(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
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
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
