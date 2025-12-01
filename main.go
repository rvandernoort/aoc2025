package main

import (
	"fmt"
	"os"
	"strconv"

	"aoc2025/days/day00"
	"aoc2025/days/day01"
)

type Solution interface {
	Part1(input string) string
	Part2(input string) string
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <day> <part> [input_file]")
		fmt.Println("Example: go run main.go 1 1 inputs/day00/input.txt")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid day: %s\n", os.Args[1])
		os.Exit(1)
	}

	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid part: %s\n", os.Args[2])
		os.Exit(1)
	}

	inputFile := fmt.Sprintf("inputs/day%02d/input.txt", day)
	if len(os.Args) >= 4 {
		inputFile = os.Args[3]
	}

	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	var solution Solution
	switch day {
	case 0:
		solution = &day00.Solution{}
	case 1:
		solution = &day01.Solution{}
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
		os.Exit(1)
	}

	var result string
	switch part {
	case 1:
		result = solution.Part1(string(input))
	case 2:
		result = solution.Part2(string(input))
	default:
		fmt.Printf("Invalid part: %d (must be 1 or 2)\n", part)
		os.Exit(1)
	}

	fmt.Println(result)
}
