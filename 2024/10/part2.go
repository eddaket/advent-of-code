package main

import "strings"

func part2(input string) int {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	sum := 0
	for i := range lines {
		for j := range lines[i] {
			sum += p2dfs(lines, j, i, n, m, '0')
		}
	}
	return sum
}

func p2dfs(lines []string, x, y, n, m int, target byte) int {
	if x < 0 || x >= m || y < 0 || y >= n {
		return 0
	}

	if lines[y][x] != target {
		return 0
	}

	if target == '9' {
		return 1
	}

	count := 0
	for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		count += p2dfs(lines, x+dir[0], y+dir[1], n, m, target+1)
	}
	return count
}
