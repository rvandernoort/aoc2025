package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Part1(input string) string {
	lines := strings.Split(input, ",")

	// Example: Sum all numbers in the input
	sum := 0
	for _, line := range lines {
		ids := strings.Split(line, "-")
		firstID, err := strconv.Atoi(ids[0])
		if err != nil {
			fmt.Println("E", err)
		}
		secondID, err := strconv.Atoi(ids[1])
		if err != nil {
			fmt.Println("E", err)
		}
		for i := firstID; i <= secondID; i++ {
			id := strconv.Itoa(i)
			if len(id)%2 != 0 {
				continue
			}
			mid := len(id) / 2
			firstHalf := id[:mid]
			secondHalf := id[mid:]
			if firstHalf == secondHalf {
				sum += i
			}
		}
	}

	return strconv.Itoa(sum)
}

func (s *Solution) Part2(input string) string {
	lines := strings.Split(input, ",")

	// Example: Sum all numbers in the input
	sum := 0
	set := make(map[int]bool)

	add := func(num int) {
		if !set[num] {
			set[num] = true
			sum += num
		}
	}

	for _, line := range lines {
		ids := strings.Split(line, "-")
		firstID, err := strconv.Atoi(ids[0])
		if err != nil {
			fmt.Println("E", err)
		}
		secondID, err := strconv.Atoi(ids[1])
		if err != nil {
			fmt.Println("E", err)
		}
		for i := firstID; i <= secondID; i++ {
			id := strconv.Itoa(i)
			mid := len(id) / 2
			for j := 1; j <= mid; j++ {
				if len(id) % j != 0 {
					continue
				}
				pattern := id[:j]
				checked := true
				for k := 0; k < len(id); k = k + j {
					if pattern != id[k:k+j] {
						checked = false
						break
					}
				}
				if checked {
					add(i)
				}
			}
		}
	}

	return strconv.Itoa(sum)
}
