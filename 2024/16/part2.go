package main

import (
	"math"
	"strings"
)

func part2(input string) int {
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
	dirMap := map[int]string{
		0: ">",
		1: "v",
		2: "<",
		3: "^",
	}
	revMap := map[string]int{
		">": 0,
		"v": 1,
		"<": 2,
		"^": 3,
	}

	min := math.MaxInt
	all := []string{}
	sub := [][]int{}
	sub2 := [][]int{}

	q := [][]p2{{{x, y, 0, 0, 0, ""}}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		for _, pair := range sub2 {
			grid[pair[0]][pair[1]] = "X"
		}

		sub2 = sub
		sub = [][]int{}

		next := []p2{}

		for len(p) > 0 {
			c := p[0]
			p = p[1:]
			for grid[c.y][c.x] != "#" && grid[c.y][c.x] != "X" {
				if grid[c.y][c.x] == "E" {
					if c.p < min {
						min = c.p
						all = []string{c.path}
					} else if c.p == min {
						all = append(all, c.path)
					}
				}

				potentialSub := []int{c.y, c.x}

				// add left and right to q
				ld := (c.d + 3) % 4
				lx := c.x + dirs[ld][0]
				ly := c.y + dirs[ld][1]
				lp := c.p + 1001
				next = append(next, p2{lx, ly, ld, lp, c.t + 1, c.path + dirMap[ld]})

				rd := (c.d + 1) % 4
				rx := c.x + dirs[rd][0]
				ry := c.y + dirs[rd][1]
				rp := c.p + 1001
				next = append(next, p2{rx, ry, rd, rp, c.t + 1, c.path + dirMap[rd]})

				c.x += dirs[c.d][0]
				c.y += dirs[c.d][1]
				c.p++
				c.path += dirMap[c.d]

				if grid[c.y][c.x] != "#" {
					sub = append(sub, potentialSub)
				}
			}
		}

		if len(next) > 0 {
			q = append(q, next)
		}
	}

	// reset grid
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = string(lines[i][j])
		}
	}

	count := 0
	for _, path := range all {
		cx, cy := x, y
		for _, r := range path {
			if grid[cy][cx] != "O" {
				count++
				grid[cy][cx] = "O"
			}

			d := revMap[string(r)]
			cx += dirs[d][0]
			cy += dirs[d][1]
		}

		if grid[cy][cx] != "O" {
			count++
			grid[cy][cx] = "O"
		}
	}

	return count
}

type p2 struct {
	x    int
	y    int
	d    int
	p    int
	t    int
	path string
}
