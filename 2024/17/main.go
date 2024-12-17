package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimRight(input, "\n")
	fmt.Printf("Part 1: %s\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}
