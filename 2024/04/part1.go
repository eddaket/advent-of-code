package main

import (
	"strings"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")

	dirs := [][]int{
		{1, -1},
		{1, 0},
		{1, 1},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}

	runes := []byte{'X', 'M', 'A', 'S'}

	sum := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != runes[0] {
				continue
			}

			for _, dir := range dirs {
				valid := true
				for k := 1; k < 4; k++ {
					newI := i + dir[0]*k
					newJ := j + dir[1]*k
					if newI < 0 || newI >= len(lines) || newJ < 0 || newJ >= len(lines[newI]) {
						valid = false
						break
					}

					if lines[newI][newJ] != runes[k] {
						valid = false
						break
					}
				}
				if valid {
					sum++
				}
			}
		}
	}
	return sum
}
