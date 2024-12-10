package main

import "strings"

func part1(input string) int {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	sum := 0
	for i := range lines {
		for j := range lines[i] {
			visited := make([][]bool, n)
			for i := range visited {
				visited[i] = make([]bool, m)
			}
			sum += p1dfs(lines, visited, j, i, n, m, '0')
		}
	}
	return sum
}

func p1dfs(lines []string, visited [][]bool, x, y, n, m int, target byte) int {
	if x < 0 || x >= m || y < 0 || y >= n {
		return 0
	}

	if lines[y][x] != target || visited[y][x] {
		return 0
	}

	visited[y][x] = true
	if target == '9' {
		return 1
	}

	count := 0
	for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		count += p1dfs(lines, visited, x+dir[0], y+dir[1], n, m, target+1)
	}
	return count
}
