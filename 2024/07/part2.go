package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part2(input string) int {
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

		p2sum = 0
		p2dfs(nums, tc, nums[0], 1)
		sum += p2sum
	}
	return sum
}

var p2sum int

func p2dfs(nums []int, target, current, offset int) {
	if p2sum != 0 {
		return
	}
	if offset == len(nums) {
		if target == current {
			p2sum += target
		}
		return
	}

	val := nums[offset]
	add := val + current
	mult := val * current
	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, val))
	p2dfs(nums, target, add, offset+1)
	p2dfs(nums, target, mult, offset+1)
	p2dfs(nums, target, concat, offset+1)
}
