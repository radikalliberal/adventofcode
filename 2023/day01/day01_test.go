package day01

import (
	"testing"
	"adventofcode/utils"
)

func TestPart1(t *testing.T) {
	ex1, _ := utils.ReadFile("test1.txt")
	tests := []struct {
		name    string
		input []string
		expected int
	}{
		{
			name:	 "2 lines 1 number",
			input:    []string{"Example 1", "Example 2"},
			expected: 33,
		},
		{
			name:	 "day1 example",
			input:    ex1,
			expected: 142,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Part1(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	ex2, _ := utils.ReadFile("test2.txt")
	tests := []struct {
		name    string
		input []string
		expected int
	}{
		{
			name:	 "2 lines 1 number",
			input:    []string{"Example 1", "Example 2"},
			expected: 33,
		},
		{
			name:	 "day1 example",
			input:    ex2,
			expected: 281,
		},
		{
			name:	 "day1 input2 p1",
			input:    []string{"fivefourhjckmndtdp98jzgqvrhbhxeighttdxkjltrdq"},
			expected: 58,
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Part2(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
