package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func part1(input string) int {
	left := []int{}
	right := []int{}
	for _, line := range strings.Split(input, "\n") {
		lineSplit := strings.Split(line, "   ")
		l, _ := strconv.Atoi(lineSplit[0])
		r, _ := strconv.Atoi(lineSplit[1])
		left = append(left, l)
		right = append(right, r)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	sum := 0

	for i := range left {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	return sum
}
