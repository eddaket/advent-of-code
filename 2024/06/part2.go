package main

import (
	"strings"
)

func part2(input string) int {
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

	ans := map[int]bool{}
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '.' {
				visited := map[int]bool{}
				p2dfs(lines, visited, ans, i*1000+j, x, y, m, n, 0)
			}
		}
	}
	return len(ans)
}

var p2Dirs = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func p2dfs(lines []string, visited, ans map[int]bool, blocked, x, y, m, n, dir int) bool {
	if x < 0 || y < 0 || x >= n || y >= m {
		return true
	}

	// 130x130. y00xd
	key := 10000*y + 10*x + dir
	if lines[y][x] == '#' || blocked == key/10 {
		return false
	}

	if visited[key] {
		ans[blocked] = true
		return true
	}

	visited[key] = true
	valid := p2dfs(lines, visited, ans, blocked, x+p2Dirs[dir][0], y+p2Dirs[dir][1], m, n, dir)
	for !valid {
		dir = (dir + 1) % 4
		valid = p2dfs(lines, visited, ans, blocked, x+p2Dirs[dir][0], y+p2Dirs[dir][1], m, n, dir)
	}
	return true
}
