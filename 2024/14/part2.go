package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2(input string) int {
	h := 103
	w := 101

	grid := make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)
	}

	robots := []*robot{}
	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, " ")
		posStr := strings.Split(pieces[0][2:], ",")
		velStr := strings.Split(pieces[1][2:], ",")
		px, _ := strconv.Atoi(posStr[0])
		py, _ := strconv.Atoi(posStr[1])
		vx, _ := strconv.Atoi(velStr[0])
		vy, _ := strconv.Atoi(velStr[1])

		robots = append(robots, &robot{
			x:  px,
			y:  py,
			dx: vx,
			dy: vy,
		})

		grid[py][px] += 1
	}

	count := 0
	for count < 10000 {
		count++
		for _, r := range robots {
			newX := r.x + r.dx
			newY := r.y + r.dy

			for newX < 0 {
				newX += w
			}

			for newY < 0 {
				newY += h
			}

			newX %= w
			newY %= h

			grid[r.y][r.x]--
			r.x = newX
			r.y = newY
			grid[r.y][r.x]++
		}
		printGrid(grid, w, h, count)
	}

	return -1
}

type robot struct {
	x  int
	y  int
	dx int
	dy int
}

func printGrid(grid [][]int, w, h, count int) {

	baseImage := image.NewRGBA(image.Rect(0, 0, w, h))

	file, err := os.Create(fmt.Sprintf("2024/14/out/%d.png", count))
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	for i, row := range grid {
		for j, col := range row {
			if col > 0 {
				baseImage.Set(j, i, color.Black)
			} else {
				baseImage.Set(j, i, color.White)
			}
		}
	}

	if err := png.Encode(file, baseImage); err != nil {
		log.Fatalf("Error encoding image: %v", err)
	}
}
