package main

import (
	"math"
	"strings"
)

func part1(input string) int {
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
		validMap := map[int]bool{0: true}
		for i := 0; i < len(line); i++ {
			if !validMap[i] {
				continue
			}

			for j := min; j <= max && i+j <= len(line); j++ {
				if towelMap[line[i:i+j]] {
					validMap[i+j] = true
				}
			}
		}

		if validMap[len(line)] {
			count++
		}
	}

	return count
}
