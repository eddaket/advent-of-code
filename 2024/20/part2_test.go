package main

import "testing"

var part2Sample = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`
var part2Expected = 41

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
