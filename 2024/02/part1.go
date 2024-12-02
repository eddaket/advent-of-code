package main

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		nums := []int{}
		for _, numStr := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		dir := 0
		if nums[0] < nums[1] {
			dir = 1
		} else {
			dir = -1
		}

		safe := true
		for i, num := range nums {
			if i == 0 {
				continue
			}

			diff := num - nums[i-1]
			if diff == 0 || diff > 3 || diff < -3 {
				safe = false
				break
			}

			if diff > 0 && dir < 0 {
				safe = false
				break
			}

			if diff < 0 && dir > 0 {
				safe = false
				break
			}
		}

		if safe {
			sum++
		}
	}
	return sum
}
