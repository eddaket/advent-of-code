package main

import (
	"strconv"
	"strings"
)

func part2(input string) int {
	price := 0
	for _, machine := range strings.Split(input, "\n\n") {
		lines := strings.Split(machine, "\n")
		a := lines[0][12:]
		b := lines[1][12:]
		p := lines[2][9:]

		as := strings.Split(a, ", Y+")
		bs := strings.Split(b, ", Y+")
		ps := strings.Split(p, ", Y=")

		ax, _ := strconv.Atoi(as[0])
		ay, _ := strconv.Atoi(as[1])
		bx, _ := strconv.Atoi(bs[0])
		by, _ := strconv.Atoi(bs[1])
		px, _ := strconv.Atoi(ps[0])
		py, _ := strconv.Atoi(ps[1])

		px += 10000000000000
		py += 10000000000000

		// a*ax + b*bx = px
		// a*ay + b*by = py

		// a = (px - b*bx) / ax
		// a = (py - b*by) / ay

		// (px - b*bx) / ax = (py - b*by) / ay
		// ay*px - ay*b*bx = ax*py - ax*b*by
		// b*ax*by - b*ay*bx = ax*py - ay*px
		// b*(ax*by - ay*bx) = ax*py - ay*px
		// b = (ax*py - ay*px) / (ax*by - ay*bx)

		bPress := (ax*py - ay*px) / (ax*by - ay*bx)
		aPress := (px - bPress*bx) / ax

		if aPress*ax+bPress*bx == px && aPress*ay+bPress*by == py {
			price += aPress*3 + bPress
		}
	}
	return price
}
