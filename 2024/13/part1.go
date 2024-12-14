package main

import (
	"math"
	"strconv"
	"strings"
)

func part1(input string) int {
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

		min := 101*3 + 101*1
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				if i*ax+j*bx == px && i*ay+j*by == py {
					min = int(math.Min(float64(min), float64(i*3+j*1)))
				}
			}
		}

		if min <= 400 {
			price += min
		}
	}
	return price
}
