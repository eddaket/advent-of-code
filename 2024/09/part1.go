package main

import (
	"strconv"
)

func part1(input string) int {
	val := []int{}
	id := 0
	on := true
	for _, r := range input {
		count, _ := strconv.Atoi(string(r))
		for count > 0 {
			if on {
				val = append(val, id)
			} else {
				val = append(val, -1)
			}
			count--
		}
		if on {
			id++
			on = false
		} else {
			on = true
		}
	}

	sum := 0
	l := 0
	r := len(val) - 1
	for l <= r {
		if val[l] >= 0 {
			sum += val[l] * l
			l++
			continue
		}

		tmp := val[l]
		val[l] = val[r]
		val[r] = tmp

		sum += val[l] * l
		l++
		for val[r] < 0 {
			r--
		}
	}
	return sum
}
