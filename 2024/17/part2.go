package main

import (
	"math"
	"strconv"
	"strings"
)

func part2(input string) int {
	lines := strings.Split(input, "\n")

	rb, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	rc, _ := strconv.Atoi(strings.Split(lines[2], ": ")[1])
	pStr := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	p := []int{}
	for _, s := range pStr {
		x, _ := strconv.Atoi(s)
		p = append(p, x)
	}

	l := len(p)
	pow := make([][]int, l)
	for i := range pow {
		pow[i] = []int{-1}
	}

	for x := 0; x < l; x++ {
		for _, combo := range allCombos(pow) {
			ra := 0
			for i := 0; i < x; i++ {
				ra += combo[i] * int(math.Pow(2, float64(3*(x-i))))
			}

			for z := 0; z < 8; z++ {
				out := p2instruct(ra, rb, rc, p)

				all := true
				for y := 0; y <= x; y++ {
					if out[len(out)-1-y] != p[len(p)-1-y] {
						all = false
						break
					}
				}

				if all {
					if x == l-1 {
						return ra
					}
					if pow[x][0] == -1 {
						pow[x][0] = z
					} else {
						inc := false
						for _, m := range pow[x] {
							if m == z {
								inc = true
								break
							}
						}
						if !inc {
							pow[x] = append(pow[x], z)
						}
					}
				}
				ra++
			}
		}
	}

	return -1
}

func allCombos(a [][]int) [][]int {
	if len(a) == 0 {
		return [][]int{{0}}
	}

	c := allCombos(a[1:])
	r := [][]int{}
	for _, x := range a[0] {
		for _, y := range c {
			n := []int{x}
			n = append(n, y...)
			r = append(r, n)
		}
	}
	return r
}

func p2instruct(a, b, c int, p []int) []int {
	i := 0
	o := 1

	out := []int{}
	k := 0
	valid := true

	for i < len(p) && o < len(p) && k < len(p) {
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
			out = append(out, n%8)
			k++
		case 6:
			n := a
			d := []int{0, 1, 2, 3, a, b, c}[p[o]]
			b = n / int(math.Pow(2, float64(d)))
		case 7:
			n := a
			d := []int{0, 1, 2, 3, a, b, c}[p[o]]
			c = n / int(math.Pow(2, float64(d)))
		}

		if !valid {
			break
		}

		if !j {
			i += 2
			o += 2
		}
	}

	return out
}
