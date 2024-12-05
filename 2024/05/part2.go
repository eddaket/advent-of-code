package main

import (
	"strconv"
	"strings"
)

func part2(input string) int {
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
				banRules[rule[1]] = map[string]bool{rule[0]: true}
			}
		} else if mode == "pages" {
			pages := strings.Split(line, ",")

			valid := true
			banned := map[string]bool{}
			set := map[string]bool{}
			for _, page := range pages {
				set[page] = true
				if banned[page] {
					valid = false
				}

				rules, ok := banRules[page]
				if ok {
					for key := range rules {
						banned[key] = true
					}
				}
			}

			if !valid {
				reqs := map[string]map[string]bool{}
				for _, page := range pages {
					reqs[page] = map[string]bool{}
					for ban := range banRules[page] {
						if set[ban] {
							reqs[page][ban] = true
						}
					}
				}

				newPages := []string{}
				for len(newPages) < len(pages) {
					real := ""
					for page, list := range reqs {
						if len(list) == 0 {
							real = page
							newPages = append(newPages, page)
							break
						}
					}
					delete(reqs, real)
					for _, list := range reqs {
						delete(list, real)
					}
				}

				mid := len(newPages) / 2
				val, _ := strconv.Atoi(newPages[mid])
				sum += val
			}
		}
	}
	return sum
}
