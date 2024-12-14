// day10_test.go
package day10

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
            input:     []string{
                "0123",
                "1234",
                "8765",
                "9876",
            },
            expected: 1,
        },
        {
            name:     "part 1 example",
            input:     []string{
                "...0...",
                "...1...",
                "...2...",
                "6543456",
                "7.....7",
                "8.....8",
                "9.....9",
            },
            expected: 2,
        },
        {
            name:     "part 1 example",
            input:     []string{
                "..90..9",
                "...1.98",
                "...2..7",
                "6543456",
                "765.987",
                "876....",
                "987....",
            },
            expected: 4,
        },
        {
            name:     "part 1 example",
            input:     []string{
                "10..9..",
                "2...8..",
                "3...7..",
                "4567654",
                "...8..3",
                "...9..2",
                ".....01",
            },
            expected: 3,
        },
        {
            name:     "part 1 example",
            input:     ex1,
            expected:  36,
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
            name:     "part 1 example",
            input:     []string{
                "..90..9",
                "...1.98",
                "...2..7",
                "6543456",
                "765.987",
                "876....",
                "987....",
            },
            expected: 13,
        },       {
            name:     "part 1 example",
            input:     []string{
                ".....0.",
                "..4321.",
                "..5..2.",
                "..6543.",
                "..7..4.",
                "..8765.",
                "..9....",
            },
            expected: 3,
        },{
            name:     "part 1 example",
            input:     []string{
                "89010123",
                "78121874",
                "87430965",
                "96549874",
                "45678903",
                "32019012",
                "01329801",
                "10456732",
            },
            expected: 81,
        },
        {
            name:     "part 2 example",
            input:    ex2,
            expected: 227,
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
