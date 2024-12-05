package main

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	mode := "rules"
	banRules := map[string]map[string]bool{}
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			mode = "pages"
			continue
		}

		if mode == "rules" {
			rule := strings.Split(line, "|")

			set, ok := banRules[rule[1]]
			if ok {
				set[rule[0]] = true
			} else {
				banRules[rule[1]] = map[string]bool{
					rule[0]: true,
				}
			}
		} else if mode == "pages" {
			pages := strings.Split(line, ",")

			valid := true
			banned := map[string]bool{}
			for _, page := range pages {
				if banned[page] {
					valid = false
					break
				}

				rules, ok := banRules[page]
				if ok {
					for key := range rules {
						banned[key] = true
					}
				}
			}

			if valid {
				mid := len(pages) / 2
				val, _ := strconv.Atoi(pages[mid])
				sum += val
			}
		}
	}
	return sum
}
