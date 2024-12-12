package main

import (
	"strings"
)

func part2(input string) int {
	lines := strings.Split(input, "\n")
	n := len(lines)
	m := len(lines[0])

	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, m)
	}

	zone := 1
	areaMap := make(map[int]int)
	sideMap := make(map[int]int)
	for i := range lines {
		for j := range lines[i] {
			r := lines[i][j]
			area := p2dfs(lines, visited, r, i, j, n, m, zone)
			if area > 0 {
				areaMap[zone] = area
				sideMap[zone] = 0
				zone++
			}
		}
	}

	sideMap[-1] = 0
	for i := -1; i < len(visited); i++ {
		last := []int{-1, -1}
		for j := range visited {
			s1 := -1
			if i >= 0 && i < len(visited) {
				s1 = visited[i][j]
			}

			s2 := -1
			if i < len(visited)-1 {
				s2 = visited[i+1][j]
			}

			if s1 == s2 {
				last = []int{-1, -1}
				continue
			}

			if s1 == last[0] && s2 == last[1] {
				continue
			}

			if s1 != last[0] {
				sideMap[s1]++
			}

			if s2 != last[1] {
				sideMap[s2]++
			}

			last = []int{s1, s2}
		}
	}

	for j := -1; j < len(visited[0]); j++ {
		last := []int{-1, -1}
		for i := range visited {
			s1 := -1
			if j >= 0 && j < len(visited[i]) {
				s1 = visited[i][j]
			}

			s2 := -1
			if j < len(visited[i])-1 {
				s2 = visited[i][j+1]
			}

			if s1 == s2 {
				last = []int{-1, -1}
				continue
			}

			if s1 == last[0] && s2 == last[1] {
				continue
			}

			if s1 != last[0] {
				sideMap[s1]++
			}

			if s2 != last[1] {
				sideMap[s2]++
			}

			last = []int{s1, s2}
		}
	}

	price := 0
	for key := range areaMap {
		price += areaMap[key] * sideMap[key]
	}
	return price
}

func p2dfs(lines []string, visited [][]int, r byte, i, j, n, m, zone int) int {
	if i < 0 || j < 0 || i >= n || j >= m {
		return 0
	}

	if lines[i][j] != r {
		return 0
	}

	if visited[i][j] != 0 {
		return 0
	}

	visited[i][j] = zone
	area := 1
	for _, dir := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		area += p2dfs(lines, visited, r, i+dir[0], j+dir[1], n, m, zone)
	}
	return area
}
