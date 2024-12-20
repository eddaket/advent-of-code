package main

import (
	"math"
	"strings"
)

func part2(input string) int {
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

	count := p2dfs(lines, grid, visited, x, y, n, m, 1, toSave)
	return count
}

var z = make(map[int]int)

func p2dfs(lines []string, grid [][]int, visited [][]bool, x, y, n, m, s, toSave int) int {
	if lines[y][x] == '#' {
		return 0
	}

	if visited[y][x] && grid[y][x] <= s {
		return 0
	}

	visited[y][x] = true
	grid[y][x] = s

	count := 0

	for cx := x - 20; cx <= x+20; cx++ {
		if cx < 0 || cx >= m {
			continue
		}

		for cy := y - 20; cy <= y+20; cy++ {
			if cy < 0 || cy >= n {
				continue
			}

			dist := int(math.Abs(float64(x-cx)) + math.Abs(float64(y-cy)))

			if dist > 20 {
				continue
			}

			if grid[cy][cx] <= 0 {
				continue
			}

			saved := s - (grid[cy][cx] + dist)
			if saved >= toSave {
				z[saved]++
				count++
			}
		}
	}

	for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		count += p2dfs(lines, grid, visited, x+dir[0], y+dir[1], n, m, s+1, toSave)
	}

	return count
}
