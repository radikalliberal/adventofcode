// day01_test.go
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
            name:     "part 1 example",
            input:     ex1,
            expected: 11,
        },
        {
            name:     "min ex",
            input:    []string{"0   2"},
            expected: 2,
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
            name:     "part 2 example",
            input:    ex2,
            expected: 31,
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
