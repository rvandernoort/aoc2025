package day02

import (
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day02/example.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part1(string(input))
	expected := "1227775554"

	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day02/example.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part2(string(input))
	expected := "4174379265"

	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day02/input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part1(string(input))
	expected := "5398419778"

	if result != expected {
		t.Errorf("Part1() = %v, want %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("../../inputs/day02/input.txt")
	if err != nil {
		t.Fatalf("Failed to read input: %v", err)
	}

	solution := &Solution{}
	result := solution.Part2(string(input))
	expected := "15704845910" // 10*20*30*40*50 = 12000000

	if result != expected {
		t.Errorf("Part2() = %v, want %v", result, expected)
	}
}
