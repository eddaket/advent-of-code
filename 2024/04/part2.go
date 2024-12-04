package main

import "strings"

func part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] != 'A' {
				continue
			}

			tl := lines[i-1][j-1]
			br := lines[i+1][j+1]
			bl := lines[i+1][j-1]
			tr := lines[i-1][j+1]

			if !(tl == 'M' || tl == 'S') {
				continue
			}
			if !(br == 'M' || br == 'S') {
				continue
			}
			if !(bl == 'M' || bl == 'S') {
				continue
			}
			if !(tr == 'M' || tr == 'S') {
				continue
			}

			if tl == br || bl == tr {
				continue
			}

			sum++
		}
	}
	return sum
}
