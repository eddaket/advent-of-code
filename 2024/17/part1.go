package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(input string) string {
	lines := strings.Split(input, "\n")

	a, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	b, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	c, _ := strconv.Atoi(strings.Split(lines[2], ": ")[1])

	pStr := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	p := []int{}
	for _, s := range pStr {
		x, _ := strconv.Atoi(s)
		p = append(p, x)
	}

	i := 0
	o := 1

	out := []string{}

	for i < len(p) && o < len(p) {
		j := false
		switch p[i] {
		case 0:
			n := a
			d := []int{0, 1, 2, 3, a, b, c}[p[o]]
			a = n / int(math.Pow(2, float64(d)))
		case 1:
			b ^= p[o]
		case 2:
			n := []int{0, 1, 2, 3, a, b, c}[p[o]]
			b = n % 8
		case 3:
			if a != 0 {
				i = p[o]
				o = i + 1
				j = true
			}
		case 4:
			b ^= c
		case 5:
			n := []int{0, 1, 2, 3, a, b, c}[p[o]]
			out = append(out, fmt.Sprint(n%8))
		case 6:
			n := a
			d := []int{0, 1, 2, 3, a, b, c}[p[o]]
			b = n / int(math.Pow(2, float64(d)))
		case 7:
			n := a
			d := []int{0, 1, 2, 3, a, b, c}[p[o]]
			c = n / int(math.Pow(2, float64(d)))
		}

		if !j {
			i += 2
			o += 2
		}
	}

	return strings.Join(out, ",")
}
