package main

import (
	"strconv"
	"strings"
)

func part2(input string) string {
	n := 70
	m := 70

	grid := make([][]int, n+1)
	for i := range grid {
		grid[i] = make([]int, m+1)
	}
	grid[n][m] = 1

	for _, line := range strings.Split(input, "\n") {
		pair := strings.Split(line, ",")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		grid[y][x] = -1

		if p2bfs(grid, n, m) < 0 {
			return line
		}

		p2cleanup(grid)
	}

	return "not found"
}

func p2bfs(grid [][]int, n, m int) int {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	q := []p2{{0, 0, 0}}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		if c.x < 0 || c.y < 0 || c.x > m || c.y > n {
			continue
		}

		if grid[c.y][c.x] == 1 {
			return c.s
		}

		if grid[c.y][c.x] != 0 {
			continue
		}

		grid[c.y][c.x] = -2
		for _, dir := range dirs {
			q = append(q, p2{c.x + dir[0], c.y + dir[1], c.s + 1})
		}
	}
	return -1
}

func p2cleanup(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == -2 {
				grid[i][j] = 0
			}
		}
	}
}

type p2 struct {
	x int
	y int
	s int
}
