package day05

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day05/example.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part1(string(input))
	expected := "3"

	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day05/example.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part2(string(input))
	expected := "14"

	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day05/input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part1(string(input))
	expected := "529"

	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day05/input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part2(string(input))
	expected := "344260049617193"

	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}
