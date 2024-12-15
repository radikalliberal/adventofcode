// day14_test.go
package day14

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
            expected: 12,
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
            expected: 1,
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

func BenchmarkPart1(b *testing.B) {
    input, _ := utils.ReadFile("input1.txt")
    for n := 0; n < b.N; n++ {
        Part1(input)
    }
}

func BenchmarkPart2(b *testing.B) {
    input, _ := utils.ReadFile("input2.txt")
    for n := 0; n < b.N; n++ {
        Part2(input)
    }
}
