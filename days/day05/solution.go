package day05

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

type Solution struct{}

func (s *Solution) Part1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	fresh := [][]int{}
	freshInput := true
	count := 0
	for _, line := range lines {
		if freshInput {
			if line == "" {
				freshInput = false
				continue
			}
			freshRange := strings.Split(line, "-")
			left, err := strconv.Atoi(freshRange[0])
			if err != nil {
				fmt.Printf("left error")
			}
			right, err := strconv.Atoi(freshRange[1])
			if err != nil {
				fmt.Printf("right error")
			}
			fresh = append(fresh, []int{left, right})
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("id error")
			}
			for _, freshRange := range fresh {
				if id >= freshRange[0] && id <= freshRange[1] {
					count += 1
					break
				}
			}
		}
	}

	return strconv.Itoa(count)
}

// func (s *Solution) Part1(input string) string {
// 	lines := strings.Split(strings.TrimSpace(input), "\n")

// 	fresh := make(map[int]bool)
// 	freshInput := true
// 	count := 0
// 	for _, line := range lines {
// 		if freshInput {
// 			if line == "" {
// 				freshInput = false
// 				continue
// 			}
// 			freshRange := strings.Split(line, "-")
// 			left, err := strconv.Atoi(freshRange[0])
// 			if err != nil {
// 				fmt.Printf("left error")
// 			}
// 			right, err := strconv.Atoi(freshRange[1])
// 			if err != nil {
// 				fmt.Printf("right error")
// 			}
// 			for i := left; i <= right; i++ {
// 				fresh[i] = true
// 			}
// 		} else {
// 			id, err := strconv.Atoi(line)
// 			if err != nil {
// 				fmt.Printf("id error")
// 			}
// 			if fresh[id] {
// 				count += 1
// 			}
// 		}
// 	}

// 	return strconv.Itoa(count)
// }

func (s *Solution) Part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	sort.Strings(lines) //thanks irene

	total := 0
	fresh := [][]int{}
	for _, line := range lines {
		if line == "" {
			break
		}
		freshRange := strings.Split(line, "-")
		left, err := strconv.Atoi(freshRange[0])
		if err != nil {
			fmt.Printf("left error")
		}
		right, err := strconv.Atoi(freshRange[1])
		if err != nil {
			fmt.Printf("right error")
		}

		// for i := left; i <= right; i++ {
		// 	duplicate := false
		// 	for _, freshRange := range fresh {
		// 		if i >= freshRange[0] && i <= freshRange[1] {
		// 			duplicate = true
		// 			break
		// 		}
		// 	}
		// 	if !duplicate {
		// 		count++
		// 	}
		// }

		addRange := true
		retry := true
		for retry {
			retry = false
			for _, freshRange := range fresh {
				if left >= freshRange[0] && left <= freshRange[1] {
					if right <= freshRange[1] {
						addRange = false
						break
					}
					left = freshRange[1] + 1
					retry = true
					break
				}
			}
		}

		if !addRange {
			continue
		}

		retry = true
		for retry {
			retry = false
			for _, freshRange := range fresh {
				if right >= freshRange[0] && right <= freshRange[1] {
					if left >= freshRange[0] {
						addRange = false
						break
					}
					right = freshRange[0] - 1
					retry = true
					break
				}
			}
		}

		if addRange {
			fresh = append(fresh, []int{left, right})
			total += right - left + 1
		}
	}

	return strconv.Itoa(total)
}
