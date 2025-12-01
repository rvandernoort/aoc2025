package day00

import (
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Example: Sum all numbers in the input
	sum := 0
	for _, line := range lines {
		if num, err := strconv.Atoi(strings.TrimSpace(line)); err == nil {
			sum += num
		}
	}

	return strconv.Itoa(sum)
}

func (s *Solution) Part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Example: Multiply all numbers in the input
	product := 1
	for _, line := range lines {
		if num, err := strconv.Atoi(strings.TrimSpace(line)); err == nil {
			product *= num
		}
	}

	return strconv.Itoa(product)
}
