package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(input string) int {
	blinks := 25
	stones := strings.Split(input, " ")

	for i := 0; i < blinks; i++ {
		newStones := []string{}
		for _, s := range stones {
			n := len(s)
			if s == "0" {
				newStones = append(newStones, "1")
			} else if n%2 == 0 {
				left := s[:n/2]
				right := strings.TrimLeft(s[n/2:], "0")
				if right == "" {
					right = "0"
				}

				newStones = append(newStones, left)
				newStones = append(newStones, right)
			} else {
				i, _ := strconv.Atoi(s)
				newStones = append(newStones, fmt.Sprint(i*2024))
			}
		}
		stones = newStones
	}
	return len(stones)
}
