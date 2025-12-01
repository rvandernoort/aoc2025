package day01

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

type Direction string

const (
	Left  Direction = "L"
	Right Direction = "R"
)

func ParseLine(line string) (Direction, int) {
	direction := Direction(line[0])
	amount, err := strconv.Atoi(line[1:])
	if err != nil {
		return "", 0
	}
	return direction, amount
}

func (s *Solution) Part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	zeros := 0
	dail := 50
	for _, line := range lines {
		direction, amount := ParseLine(line)
		amount = amount % 100
		switch direction {
		case Left:
			if amount == dail {
				dail = 0
			} else if amount >= dail {
				dail = 100 - (amount - dail)
			} else {
				dail -= amount
			}
		case Right:
			if amount+dail > 99 {
				dail = amount + dail - 100
			} else {
				dail += amount
			}
		default:
			return "Failed"
		}
		fmt.Print(strconv.Itoa(dail) + " - ")
		if dail == 0 {
			zeros += 1
		}
	}

	return strconv.Itoa(zeros)
}

func (s *Solution) Part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	zeros := 0
	dail := 50
	for _, line := range lines {
		direction, amount := ParseLine(line)
		zeros += amount / 100
		amount = amount % 100
		switch direction {
		case Left:
			if dail == 0 {
				dail = 100 - amount
			} else if amount > dail {
				dail = 100 - (amount - dail)
				zeros += 1
			} else if amount == dail {
				dail = 0
				zeros += 1
			} else {
				dail = dail - amount
			}
		case Right:
			if amount+dail > 99 {
				dail = (amount + dail) - 100
				zeros += 1
			} else if (amount + dail) == 100 {
				dail = 0
				zeros += 1
			} else {
				dail = dail + amount
			}
		default:
			return "Failed"
		}
	}

	return strconv.Itoa(zeros)
}
