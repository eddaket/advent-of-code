package main

import (
	"strings"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	toSave := 100

	x := -1
	y := -1
	visited := make([][]bool, n)
	grid := make([][]int, n)
	for i := range lines {
		visited[i] = make([]bool, m)
		grid[i] = make([]int, m)
		for j := range lines[i] {
			if lines[i][j] == '#' {
				grid[i][j] = -1
			}

			if lines[i][j] == 'S' {
				x = j
				y = i
			}
		}
	}

	count := p1dfs(lines, grid, visited, x, y, n, m, 1, toSave)
	return count
}

func p1dfs(lines []string, grid [][]int, visited [][]bool, x, y, n, m, s, toSave int) int {
	if lines[y][x] == '#' {
		return 0
	}

	if visited[y][x] && grid[y][x] <= s {
		return 0
	}

	visited[y][x] = true
	grid[y][x] = s

	count := 0
	for _, dir := range [][]int{{0, 2}, {2, 0}, {0, -2}, {-2, 0}} {
		nx := x + dir[0]
		ny := y + dir[1]
		if nx < 0 || ny < 0 || nx >= m || ny >= n {
			continue
		}

		if grid[ny][nx] <= 0 {
			continue
		}

		if s-grid[ny][nx]-2 >= toSave {
			count++
		}
	}

	for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		count += p1dfs(lines, grid, visited, x+dir[0], y+dir[1], n, m, s+1, toSave)
	}

	return count
}
