package day03

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Example: Sum all numbers in the input
	sum := 0
	for _, line := range lines {
		numbers := strings.Split(line, "")
		firstHighest := 0
		latestIndex := 0
		for i := 0; i < len(numbers)-1; i++ {
			currentNumber, err := strconv.Atoi(numbers[i])
			if err != nil {
				fmt.Println("E", err)
			}
			if currentNumber > firstHighest {
				firstHighest = currentNumber
				latestIndex = i
			}
		}
		secondHighest := 0
		for i := latestIndex + 1; i < len(numbers); i++ {
			currentNumber, err := strconv.Atoi(numbers[i])
			if err != nil {
				fmt.Println("E", err)
			}
			if currentNumber > secondHighest {
				secondHighest = currentNumber
			}
		}
		sum += firstHighest*10 + secondHighest
	}

	return strconv.Itoa(sum)
}

func (s *Solution) Part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Example: Sum all numbers in the input
	sum := 0
	for _, line := range lines {
		numbers := strings.Split(line, "")
		lastIndex := 0
		endIndex := len(numbers) - 12 + 1
		highestString := ""
		for i := 12; i >= 2; i-- {
			highest, index := findLargestNumber(numbers[lastIndex:endIndex]) //count the skipped
			lastIndex += index + 1
			endIndex += 1
			highestString += strconv.Itoa(highest)
		}

		highest, _ := findLargestNumber(numbers[lastIndex:])
		highestString += strconv.Itoa(highest)

		highest, err := strconv.Atoi(highestString)
		if err != nil {
			fmt.Println("E", err)
		}
		sum += highest
	}

	return strconv.Itoa(sum)
}

func findLargestNumber(numbers []string) (int, int) {
	highest := 0
	latestIndex := 0
	for i := range numbers {
		currentNumber, err := strconv.Atoi(numbers[i])
		if err != nil {
			fmt.Println("E", err)
		}
		if currentNumber > highest {
			highest = currentNumber
			latestIndex = i
		}
	}
	return highest, latestIndex
}
