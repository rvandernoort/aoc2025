package day04

import (
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := [][]string{}
	x_len := len(lines[0])
	y_len := len(lines)

	line := []string{}
	line = append(line, ".")
	for range x_len {
		line = append(line, ".")
	}
	line = append(line, ".")
	grid = append(grid, line)

	for _, line := range lines {
		addedLine := []string{}
		addedLine = append(addedLine, ".")
		parsedLine := strings.Split(line, "")
		for _, char := range parsedLine {
			addedLine = append(addedLine, char)
		}
		addedLine = append(addedLine, ".")
		grid = append(grid, addedLine)
	}

	line = []string{}
	line = append(line, ".")
	for range x_len {
		line = append(line, ".")
	}
	line = append(line, ".")
	grid = append(grid, line)

	accessible := 0

	for i := 1; i < x_len+1; i++ {
		for j := 1; j < y_len+1; j++ {
			if grid[i][j] == string('@') && CheckAdjacent(grid, i, j) {
				accessible += 1
			}
		}
	}

	return strconv.Itoa(accessible)
}

func CheckAdjacent(grid [][]string, x int, y int) bool {
	adjacent := 0

	if grid[x-1][y-1] == string('@') {
		adjacent += 1
	}
	if grid[x][y-1] == string('@') {
		adjacent += 1
	}
	if grid[x+1][y-1] == string('@') {
		adjacent += 1
	}
	if grid[x-1][y] == string('@') {
		adjacent += 1
	}
	if grid[x+1][y] == string('@') {
		adjacent += 1
	}
	if grid[x-1][y+1] == string('@') {
		adjacent += 1
	}
	if grid[x][y+1] == string('@') {
		adjacent += 1
	}
	if grid[x+1][y+1] == string('@') {
		adjacent += 1
	}

	return adjacent < 4
}

func (s *Solution) Part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := [][]string{}
	x_len := len(lines[0])
	y_len := len(lines)

	line := []string{}
	line = append(line, ".")
	for range x_len {
		line = append(line, ".")
	}
	line = append(line, ".")
	grid = append(grid, line)

	for _, line := range lines {
		addedLine := []string{}
		addedLine = append(addedLine, ".")
		parsedLine := strings.Split(line, "")
		for _, char := range parsedLine {
			addedLine = append(addedLine, char)
		}
		addedLine = append(addedLine, ".")
		grid = append(grid, addedLine)
	}

	line = []string{}
	line = append(line, ".")
	for range x_len {
		line = append(line, ".")
	}
	line = append(line, ".")
	grid = append(grid, line)

	accessible := 1
	total_accessible := 0
	removedPaper := [][]int{}

	for accessible > 0 {
		accessible = 0
		for i := 1; i < x_len+1; i++ {
			for j := 1; j < y_len+1; j++ {
				if grid[i][j] == string('@') && CheckAdjacent(grid, i, j) {
					accessible += 1
					removedPaper = append(removedPaper, []int{i, j})
				}
			}
		}
		for _, cords := range removedPaper {
			grid[cords[0]][cords[1]] = "."
		}
		total_accessible += accessible
	}

	return strconv.Itoa(total_accessible)
}
