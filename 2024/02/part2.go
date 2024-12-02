package main

import (
	"strconv"
	"strings"
)

func part2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		nums := []int{}
		for _, numStr := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		for j := -1; j < len(nums); j++ {
			var numsRemoved []int
			if j == -1 {
				numsRemoved = append(numsRemoved, nums...)
			} else {
				numsRemoved = append(numsRemoved, nums[:j]...)
				numsRemoved = append(numsRemoved, nums[j+1:]...)
			}

			dir := 0
			if numsRemoved[0] < numsRemoved[1] {
				dir = 1
			} else {
				dir = -1
			}

			safe := true
			for i, num := range numsRemoved {
				if i == 0 {
					continue
				}

				diff := num - numsRemoved[i-1]
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
				break
			}
		}
	}
	return sum
}
