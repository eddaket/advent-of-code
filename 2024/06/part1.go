package main

import (
	"strings"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")
	m := len(lines)
	n := len(lines[0])
	x := -1
	y := -1
	for i, line := range lines {
		for j, r := range line {
			if r == '^' {
				x = j
				y = i
				break
			}
		}
		if x != -1 {
			break
		}
	}

	visited := map[int]bool{}
	dist, _ := p1dfs(lines, visited, x, y, m, n, 0)
	return dist
}

var p1Dirs = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func p1dfs(lines []string, visited map[int]bool, x, y, m, n, dir int) (int, bool) {
	if x < 0 || y < 0 || x >= n || y >= m {
		return 0, true
	}

	if lines[y][x] == '#' {
		return 0, false
	}

	// 130x130. y00x
	count := 0
	key := 1000*y + x
	if !visited[key] {
		visited[key] = true
		count++
	}

	dist, valid := p1dfs(lines, visited, x+p1Dirs[dir][0], y+p1Dirs[dir][1], m, n, dir)
	for !valid {
		dir = (dir + 1) % 4
		dist, valid = p1dfs(lines, visited, x+p1Dirs[dir][0], y+p1Dirs[dir][1], m, n, dir)
	}

	return dist + count, true
}
