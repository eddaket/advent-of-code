package main

import (
	"strings"
)

func part1(input string) int {
	parts := strings.Split(input, "\n\n")
	gridStr := strings.Split(parts[0], "\n")
	inst := strings.Join(strings.Split(parts[1], "\n"), "")

	n := len(gridStr)
	m := len(gridStr[0])
	grid := make([][]string, n)
	for i := range grid {
		grid[i] = make([]string, m)
	}

	x := -1
	y := -1
	for i := range gridStr {
		for j := range gridStr[i] {
			grid[i][j] = string(gridStr[i][j])
			if gridStr[i][j] == '@' {
				x = j
				y = i
			}
		}
	}

	dirs := map[rune][]int{
		'^': {0, -1},
		'>': {1, 0},
		'v': {0, 1},
		'<': {-1, 0},
	}

	for _, r := range inst {
		dir := dirs[r]

		moved := p1slide(grid, x, y, dir)
		if moved {
			x += dir[0]
			y += dir[1]
		}
	}

	sum := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "O" {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func p1slide(grid [][]string, x, y int, dir []int) bool {
	if grid[y][x] == "#" {
		return false
	}

	if grid[y][x] == "." {
		return true
	}

	xp := x + dir[0]
	yp := y + dir[1]

	p1slide(grid, xp, yp, dir)

	if grid[yp][xp] == "." {
		grid[yp][xp] = grid[y][x]
		grid[y][x] = "."
		return true
	}
	return false
}
