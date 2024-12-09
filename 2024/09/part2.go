package main

import (
	"strconv"
)

func part2(input string) int {
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

	i := len(val) - 1
	for i >= 0 {
		if val[i] < 0 {
			i--
			continue
		}

		j := i - 1
		for j >= 0 && val[j] == val[i] {
			j--
		}
		block := i - j

		k := 0
		for k < i {
			if val[k] >= 0 {
				k++
				continue
			}

			l := k + 1
			for l < i && val[l] == val[k] {
				l++
			}
			eBlock := l - k
			if block <= eBlock {
				for k < l && i > j {
					tmp := val[i]
					val[i] = val[k]
					val[k] = tmp
					i--
					k++
				}
				break
			}
			k = l
		}
		i = j
	}

	sum := 0
	for i, id := range val {
		if id >= 0 {
			sum += i * id
		}
	}
	return sum
}
