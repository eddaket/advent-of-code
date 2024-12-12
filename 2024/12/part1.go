package main

import (
	"strings"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	price := 0
	for i := range lines {
		for j := range lines[i] {
			r := lines[i][j]
			area, perim := p1dfs(lines, visited, r, i, j, n, m)
			price += area * perim
		}
	}
	return price
}

func p1dfs(lines []string, visited [][]bool, r byte, i, j, n, m int) (int, int) {
	if i < 0 || j < 0 || i >= n || j >= m {
		return 0, 1
	}

	if lines[i][j] != r {
		return 0, 1
	}

	if visited[i][j] {
		return 0, 0
	}

	visited[i][j] = true
	area := 1
	perim := 0
	for _, dir := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		a, p := p1dfs(lines, visited, r, i+dir[0], j+dir[1], n, m)
		area += a
		perim += p
	}
	return area, perim
}
