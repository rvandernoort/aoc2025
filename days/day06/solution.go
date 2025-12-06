package day06

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Example: Sum all numbers in the input
	numbers := [][]int{}
	total := 0
	for _, line := range lines {
		numberStrings := strings.Fields(line)
		if numberStrings[0] == "*" || numberStrings[0] == "+" {
			for i, o := range numberStrings {
				if o == "*" {
					mul := 1
					for _, ls := range numbers {
						mul *= ls[i]
					}
					total += mul
				}
				if o == "+" {
					add := 0
					for _, ls := range numbers {
						add += ls[i]
					}
					total += add
				}
			}
			break
		}
		numberLine := []int{}
		for _, n := range numberStrings {
			nInt, err := strconv.Atoi(n)
			if err != nil {
				fmt.Printf("n error")
			}
			numberLine = append(numberLine, nInt)
		}
		numbers = append(numbers, numberLine)
	}

	return strconv.Itoa(total)
}

func (s *Solution) Part2(input string) string {
	lines := strings.Split(input, "\n")

	// Example: Multiply all numbers in the input
	total := 0
	currentNumber := string("")
	currentNumbers := []int{}
	for i := len(lines[0]) - 1; i >= 0; i-- {
		noOperator := true
		for j := 0; j < len(lines); j++ {
			lineList := strings.Split(lines[j], "")

			if lineList[i] == "*" {
				currentNumbers = addCurrentNumber(currentNumbers, currentNumber)
				mul := 1
				for _, c := range currentNumbers {
					mul *= c
				}
				total += mul
				noOperator = false
				currentNumber = string("")
				currentNumbers = []int{}
				i--
				break
			}

			if lineList[i] == "+" {
				currentNumbers = addCurrentNumber(currentNumbers, currentNumber)
				add := 0
				for _, c := range currentNumbers {
					add += c
				}
				total += add
				noOperator = false
				currentNumber = string("")
				currentNumbers = []int{}
				i--
				break
			}

			_, err := strconv.Atoi(lineList[i])
			if err != nil {
				continue
			}
			currentNumber = currentNumber + lineList[i]
		}

		if noOperator {
			currentNumbers = addCurrentNumber(currentNumbers, currentNumber)
			currentNumber = string("")
		}

	}

	return strconv.Itoa(total)
}

func addCurrentNumber(currentNumbers []int, currentNumber string) []int {
	currentNumberInt, err := strconv.Atoi(currentNumber)
	if err != nil {
		fmt.Printf("currentNumber error")
	}
	currentNumbers = append(currentNumbers, currentNumberInt)
	return currentNumbers
}
