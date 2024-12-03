package main

import (
	"strconv"
	"strings"
)

func part2(input string) int {
	sum := 0
	for _, piece := range strings.Split(input, "do()") {
		piece = strings.Split(piece, "don't()")[0]
		for _, piece2 := range strings.Split(piece, "mul(") {
			piece2 = strings.Split(piece2, ")")[0]
			pieces := strings.Split(piece2, ",")
			if len(pieces) < 2 {
				continue
			}

			num1, err1 := strconv.Atoi(pieces[0])
			num2, err2 := strconv.Atoi(pieces[1])
			if err1 != nil || err2 != nil {
				continue
			}

			sum += num1 * num2
		}
	}
	return sum
}
