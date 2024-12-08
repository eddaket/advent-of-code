package main

import (
	"math"
	"strings"
)

func part2(input string) int {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	nodeMap := make(map[rune][][]int)
	for i, line := range lines {
		for j, r := range line {
			if r == '.' {
				continue
			}

			list, ok := nodeMap[r]
			if !ok {
				list = make([][]int, 0)
			}
			list = append(list, []int{i, j})
			nodeMap[r] = list
		}
	}

	nodes := make([][]bool, n)
	for i := range nodes {
		nodes[i] = make([]bool, m)
	}

	count := 0
	for _, list := range nodeMap {
		for x := range list {
			for y := range list {
				if y <= x {
					continue
				}

				i1 := list[x][0]
				j1 := list[x][1]
				i2 := list[y][0]
				j2 := list[y][1]
				iDiff := int(math.Abs(float64(i1 - i2)))
				jDiff := int(math.Abs(float64(j1 - j2)))

				i1valid := true
				j1valid := true
				i2valid := true
				j2valid := true

				for (i1valid && j1valid) || (i2valid && j2valid) {
					if i1valid && j1valid && !nodes[i1][j1] {
						nodes[i1][j1] = true
						count++
					}

					if i2valid && j2valid && !nodes[i2][j2] {
						nodes[i2][j2] = true
						count++
					}

					if i1 <= i2 {
						i1 -= iDiff
						i2 += iDiff
					} else {
						i1 += iDiff
						i2 -= iDiff
					}

					if j1 <= j2 {
						j1 -= jDiff
						j2 += jDiff
					} else {
						j1 += jDiff
						j2 -= jDiff
					}

					i1valid = i1 >= 0 && i1 < n
					j1valid = j1 >= 0 && j1 < m
					i2valid = i2 >= 0 && i2 < n
					j2valid = j2 >= 0 && j2 < m
				}
			}
		}
	}
	return count
}
