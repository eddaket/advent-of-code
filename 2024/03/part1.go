package main

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	sum := 0
	for i, piece := range strings.Split(input, "mul(") {
		if i == 0 {
			continue
		}

		piece = strings.Split(piece, ")")[0]
		pieces := strings.Split(piece, ",")
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
	return sum
}
