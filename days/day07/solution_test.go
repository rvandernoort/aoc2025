package day07

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day07/example.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part1(string(input))
	expected := "21"

	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day07/example.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part2(string(input))
	expected := "40"

	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day07/input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part1(string(input))
	expected := "150" // 10+20+30+40+50 = 150

	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day07/input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part2(string(input))
	expected := "12000000" // 10*20*30*40*50 = 12000000

	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}
