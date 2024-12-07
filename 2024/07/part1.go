package main

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		lineSplit := strings.Split(line, ": ")
		if len(lineSplit) != 2 {
			continue
		}
		tc, _ := strconv.Atoi(lineSplit[0])
		numStrs := strings.Split(lineSplit[1], " ")
		nums := []int{}
		for _, i := range numStrs {
			j, _ := strconv.Atoi(i)
			nums = append(nums, j)
		}

		p1sum = 0
		p1dfs(nums, tc, nums[0], 1)
		sum += p1sum
	}
	return sum
}

var p1sum int

func p1dfs(nums []int, target, current, offset int) {
	if p1sum != 0 {
		return
	}
	if offset == len(nums) {
		if target == current {
			p1sum += target
		}
		return
	}

	val := nums[offset]
	add := val + current
	mult := val * current
	p1dfs(nums, target, add, offset+1)
	p1dfs(nums, target, mult, offset+1)
}
