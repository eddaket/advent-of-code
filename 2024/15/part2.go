package main

import (
	"fmt"
	"strings"
)

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	gridStr := strings.Split(parts[0], "\n")
	inst := strings.Join(strings.Split(parts[1], "\n"), "")

	n := len(gridStr)
	m := len(gridStr[0]) * 2
	grid := make([][]string, n)
	for i := range grid {
		grid[i] = make([]string, m)
	}

	x := -1
	y := -1
	for i := range gridStr {
		for j := range gridStr[i] {
			switch string(gridStr[i][j]) {
			case "O":
				grid[i][j*2] = "["
				grid[i][j*2+1] = "]"
			case "@":
				grid[i][j*2] = "@"
				grid[i][j*2+1] = "."
			default:
				grid[i][j*2] = string(gridStr[i][j])
				grid[i][j*2+1] = string(gridStr[i][j])
			}

			if gridStr[i][j] == '@' {
				x = j * 2
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

		if dir[1] == 0 {
			// left/right is the same as part 1 so use that
			moved := p1slide(grid, x, y, dir)
			if moved {
				x += dir[0]
				y += dir[1]
			}
			continue
		}

		moved := p2slide(grid, map[int]bool{x: true}, y, dir)
		if moved {
			x += dir[0]
			y += dir[1]
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	sum := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "[" {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func p2slide(grid [][]string, allX map[int]bool, y int, dir []int) bool {
	for x := range allX {
		if grid[y][x] == "#" {
			return false
		}

		if grid[y][x] == "." {
			delete(allX, x)
		}

		if grid[y][x] == "]" {
			allX[x-1] = true
		}

		if grid[y][x] == "[" {
			allX[x+1] = true
		}
	}

	if len(allX) == 0 {
		return true
	}

	yp := y + dir[1]

	allXC := make(map[int]bool)
	for x := range allX {
		allXC[x] = true
	}

	p2slide(grid, allXC, yp, dir)

	allDot := true
	for x := range allX {
		if grid[yp][x] != "." {
			allDot = false
			break
		}
	}
	if allDot {
		for x := range allX {
			grid[yp][x] = grid[y][x]
			grid[y][x] = "."
		}
		return true
	}
	return false
}
