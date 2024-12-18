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
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %s\n", part2(input))
}
