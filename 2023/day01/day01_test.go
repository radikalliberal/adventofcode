package day01

import (
	"testing"
)

func TestPart1(t *testing.T) {
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
			input:    []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
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
			input:    []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen" },
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
