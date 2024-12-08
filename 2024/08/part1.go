package main

import (
	"math"
	"strings"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	nodes := make([][]bool, n)
	for i := range nodes {
		nodes[i] = make([]bool, m)
	}

	nodeMap := make(map[rune][][]int)
	for i, line := range lines {
		for j, r := range line {
			if r == '.' {
				continue
			}

			// nodes[i][j] = true
			list, ok := nodeMap[r]
			if !ok {
				list = make([][]int, 0)
			}
			list = append(list, []int{i, j})
			nodeMap[r] = list
		}
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

				var ip1, ip2, jp1, jp2 int
				if i1 <= i2 {
					ip1 = i1 - iDiff
					ip2 = i2 + iDiff
				} else {
					ip1 = i1 + iDiff
					ip2 = i2 - iDiff
				}

				if j1 <= j2 {
					jp1 = j1 - jDiff
					jp2 = j2 + jDiff
				} else {
					jp1 = j1 + jDiff
					jp2 = j2 - jDiff
				}

				if ip1 >= 0 && ip1 < n && jp1 >= 0 && jp1 < m && !nodes[ip1][jp1] {
					nodes[ip1][jp1] = true
					count++
				}

				if ip2 >= 0 && ip2 < n && jp2 >= 0 && jp2 < m && !nodes[ip2][jp2] {
					nodes[ip2][jp2] = true
					count++
				}
			}
		}
	}
	return count
}
