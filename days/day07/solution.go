package day07

import (
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := [][]string{}
	startIndexX := 0
	for i, line := range lines {
		lineList := strings.Split(line, "")
		if i == 0 {
			for j, s := range lineList {
				if s == "S" {
					startIndexX = j
				}
			}
		}
		grid = append(grid, lineList)
	}

	beams := make(map[int]bool)
	indexY := 1
	beams[startIndexX] = true
	splits := 0

	for indexY < len(lines)-1 {
		nextBeams := make(map[int]bool)
		for key := range beams {
			if grid[indexY+1][key] == "^" {
				nextBeams[key-1] = true
				nextBeams[key+1] = true
				splits++
			} else {
				nextBeams[key] = true
			}
		}
		beams = nextBeams
		indexY++
	}

	return strconv.Itoa(splits)
}

func (s *Solution) Part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := [][]string{}
	startIndexX := 0
	for i, line := range lines {
		lineList := strings.Split(line, "")
		if i == 0 {
			for j, s := range lineList {
				if s == "S" {
					startIndexX = j
				}
			}
		}
		grid = append(grid, lineList)
	}

	beams := make(map[int]int)
	indexY := 1
	beams[startIndexX] = 1
	splits := 0

	for indexY < len(lines)-1 {
		nextBeams := make(map[int]int)
		for key := range beams {
			if grid[indexY+1][key] == "^" {
				nextBeams[key-1] += beams[key]
				nextBeams[key+1] += beams[key]
				splits++
			} else {
				nextBeams[key] += beams[key]
			}
		}
		beams = nextBeams
		indexY++
	}

	total := 0
	for _, count := range beams {
		total += count
	}

	return strconv.Itoa(total)
}
