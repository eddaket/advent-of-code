package main

import (
	"strings"
)

func part1(input string) int {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	x := -1
	y := -1

	grid := make([][]string, n)
	for i := range grid {
		grid[i] = make([]string, m)
		for j := range grid[i] {
			grid[i][j] = string(lines[i][j])
			if lines[i][j] == 'S' {
				x, y = j, i
			}
		}
	}

	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	q := []p1{{x, y, 0, 0}}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		for grid[c.y][c.x] != "#" {
			if grid[c.y][c.x] == "E" {
				return c.p
			}

			grid[c.y][c.x] = "#"

			// add left and right to q
			ld := (c.d + 3) % 4
			lx := c.x + dirs[ld][0]
			ly := c.y + dirs[ld][1]
			lp := c.p + 1001
			q = append(q, p1{lx, ly, ld, lp})

			rd := (c.d + 1) % 4
			rx := c.x + dirs[rd][0]
			ry := c.y + dirs[rd][1]
			rp := c.p + 1001
			q = append(q, p1{rx, ry, rd, rp})

			c.x += dirs[c.d][0]
			c.y += dirs[c.d][1]
			c.p++
		}
	}
	return -1
}

type p1 struct {
	x int
	y int
	d int
	p int
}
