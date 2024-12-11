package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part2(input string) int {
	blinks := 75
	stones := strings.Split(input, " ")

	dp := map[string]int{}
	sum := 0
	for _, stone := range stones {
		sum += p2dfs(stone, 0, blinks, dp)
	}

	return sum
}

func p2dfs(stone string, current, total int, dp map[string]int) int {
	if current == total {
		return 1
	}

	key := fmt.Sprintf("%s_%d", stone, current)
	if memo, ok := dp[key]; ok {
		return memo
	}

	sum := 0
	if stone == "0" {
		sum += p2dfs("1", current+1, total, dp)
	} else if len(stone)%2 == 0 {
		l := len(stone) / 2
		left := stone[:l]
		right := strings.TrimLeft(stone[l:], "0")
		if right == "" {
			right = "0"
		}
		sum += p2dfs(left, current+1, total, dp)
		sum += p2dfs(right, current+1, total, dp)
	} else {
		stoneI, _ := strconv.Atoi(stone)
		sum += p2dfs(fmt.Sprint(stoneI*2024), current+1, total, dp)
	}

	dp[key] = sum
	return sum
}
