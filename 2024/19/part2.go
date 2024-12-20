package main

import (
	"math"
	"strings"
)

func part2(input string) int {
	lines := strings.Split(input, "\n")
	towels := strings.Split(lines[0], ", ")

	min := math.MaxInt
	max := 0
	towelMap := map[string]bool{}
	for _, towel := range towels {
		if len(towel) < min {
			min = len(towel)
		}

		if len(towel) > max {
			max = len(towel)
		}

		towelMap[towel] = true
	}

	count := 0
	for _, line := range lines[2:] {
		validMap := map[int]int{0: 1}
		for i := 0; i < len(line); i++ {
			if validMap[i] == 0 {
				continue
			}

			for j := min; j <= max && i+j <= len(line); j++ {
				if towelMap[line[i:i+j]] {
					validMap[i+j] += validMap[i]
				}
			}
		}
		count += validMap[len(line)]
	}

	return count
}
