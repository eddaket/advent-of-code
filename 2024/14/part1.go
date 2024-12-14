package main

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	s := 100
	h := 103
	mh := h / 2
	w := 101
	mw := w / 2

	quad := make([]int, 4)
	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, " ")
		posStr := strings.Split(pieces[0][2:], ",")
		velStr := strings.Split(pieces[1][2:], ",")
		px, _ := strconv.Atoi(posStr[0])
		py, _ := strconv.Atoi(posStr[1])
		vx, _ := strconv.Atoi(velStr[0])
		vy, _ := strconv.Atoi(velStr[1])

		x := px + s*vx
		y := py + s*vy

		for x < 0 {
			x += w
		}

		for y < 0 {
			y += h
		}

		x %= w
		y %= h

		if x == mw || y == mh {
			continue
		}

		i := 0
		if x > mw {
			i += 2
		}

		if y > mh {
			i += 1
		}

		quad[i]++
	}

	prod := 1
	for _, q := range quad {
		prod *= q
	}
	return prod
}
