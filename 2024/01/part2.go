package main

import (
	"strconv"
	"strings"
)

func part2(input string) int {
	left := []int{}
	right := map[int]int{}
	for _, line := range strings.Split(input, "\n") {
		lineSplit := strings.Split(line, "   ")
		l, _ := strconv.Atoi(lineSplit[0])
		r, _ := strconv.Atoi(lineSplit[1])
		left = append(left, l)
		right[r] = right[r] + 1
	}

	sum := 0
	for _, i := range left {
		sum += i * right[i]
	}
	return sum
}
