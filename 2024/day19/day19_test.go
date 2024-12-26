// day19_test.go
package day19

import (
    "testing"
    "errors"
    "adventofcode/utils"
)

func TestParseInputPart1(t *testing.T) {
    ex1, _ := utils.ReadFile("test1.txt")
    patterns, targets := ParseInput(ex1)
    expectedPatterns := `r
wr
b
g
bwu
rb
gb
br
`
    expectedTargets := `brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`
    if patterns.ToString() != expectedPatterns {
        t.Errorf("Expected %s, got %s", expectedPatterns, patterns.ToString())
    }
    if targets.ToString() != expectedTargets {
        t.Errorf("Expected %s, got %s", expectedTargets, targets.ToString())
    }
}

func TestFindMatches(t *testing.T) {
    p := ColorPattern{colors: []rune{'r'}, length: 1}

    tests := []struct {
        name    string
        input Target
        expected []int
    }{
        {
            name:     "single match",
            input:    Target{colors: []rune{'r'}},
            expected: []int{0},
        },
        {
            name:     "no match",
            input:    Target{colors: []rune{'g'}},
            expected: []int{},
        },
        {
            name:     "multiple matches",
            input:    Target{colors: []rune{'r', 'r'}},
            expected: []int{0, 1},
        },
        {
            name:     "multiple matches",
            input:    Target{colors: []rune{'r', 'b', 'g', 'r'}},
            expected: []int{0, 3},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := p.FindMatches(&tt.input)
            if !utils.IntArrayEquals(result, tt.expected) {
                t.Errorf("Expected %v, got %v", tt.expected, result)
            }
        })
    }
}

func TestSplitTarget(t *testing.T) {
    tests := []struct {
        name    string
        cp ColorPattern
        index int
        target Target
        expected Targets
        expectedError error
    }{
        {
            name:     "single match",
            cp:       ColorPattern{colors: []rune{'r'}, length: 1},
            index:    0,
            target:   Target{colors: []rune{'r'}},
            expected: Targets{{colors: []rune{}}, {colors: []rune{}}},
            expectedError: nil,
        },
        {
            name:     "no match",
            cp:       ColorPattern{colors: []rune{'r'}, length: 1},
            index:    0,
            target:   Target{colors: []rune{'g'}},
            expected: nil,
            expectedError: errors.New("No match"),
        },
        {
            name:     "multiple matches",
            cp:       ColorPattern{colors: []rune{'r', 'b'}, length: 2},
            index:    1,
            target:   Target{colors: []rune{'r', 'r', 'b', 'b'}},
            expected: Targets{{colors: []rune{'r'}}, {colors: []rune{'b'}}},
            expectedError: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result1, result2, err := tt.cp.Split(&tt.target, tt.index)
            if tt.expectedError != nil && err == nil {
                t.Errorf("Expected %v, got %v", tt.expectedError, err)
            }
            if tt.expectedError == nil && err != nil {
                t.Errorf("Expected %v, got %v", tt.expectedError, err)
            }
            if err == nil {
                if result1.ToString() != tt.expected[0].ToString() {
                    t.Errorf("Expected %s, got %s", tt.expected[0].ToString(), result1.ToString())
                }
                if result2.ToString() != tt.expected[1].ToString() {
                    t.Errorf("Expected %s, got %s", tt.expected[1].ToString(), result2.ToString())
                }
            }
        })
    }

}

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
            expected: 6,
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

func TestCombinePermutations(t *testing.T) {
    p1 := []PickedPatterns{
        {"r", "b"},
        {"r", "g"},
    }
    p2 := []PickedPatterns{
        {"w", "r"},
        {"b", "r"},
    }
    expected := []PickedPatterns{
        {"r", "b", "w", "r"},
        {"r", "b", "b", "r"},
        {"r", "g", "w", "r"},
        {"r", "g", "b", "r"},
    }
    result := CombinePermutations(&p1,&p2)
    for i, r := range result {
        if r.ToString() != expected[i].ToString() {
            t.Errorf("Expected %s, got %s", expected[i].ToString(), r.ToString())
        }
    }
}

func TestFindAllPermutations1(t *testing.T) {
    patterns := Patterns{
        patterns: []ColorPattern{
            {colors: []rune{'r'}, length: 1},
        },
    }

    target := Target{colors: []rune{'r'}}
    picked := PickedPatterns{}
    expected := []PickedPatterns{
        {"r"},
    }
    result := patterns.FindAllPermutations(target, picked, nil)
    if len(result) != len(expected) { 
        t.Errorf("Expected %d, got %d", len(expected), len(result))
    }
    for i, r := range result {
        if r.ToString() != expected[i].ToString() {
            t.Errorf("Expected %s, got %s", expected[i].ToString(), r.ToString())
        }
    }
    memo := map[string][]PickedPatterns{}
    result = patterns.FindAllPermutations(target, picked, &memo)
    if len(result) != len(expected) {
        t.Errorf("Expected %d, got %d", len(expected), len(result))
    }
    for i, r := range result {
        if r.ToString() != expected[i].ToString() {
            t.Errorf("Expected %s, got %s", expected[i].ToString(), r.ToString())
        }
    }
}

func TestFindAllPermutations2(t *testing.T) {
    patterns := Patterns{
        patterns: []ColorPattern{
            {colors: []rune{'r'}, length: 1},
            {colors: []rune{'w', 'r'}, length: 2},
            {colors: []rune{'b'}, length: 1},
        },
    }

    target := Target{colors: []rune{'b', 'r'}}
    picked := PickedPatterns{}
    expected := []PickedPatterns{
        {"b", "r"},
    }
    result := patterns.FindAllPermutations(target, picked, nil)
    if len(result) != len(expected) { 
        t.Errorf("Expected %d, got %d", len(expected), len(result))
    }
    for i, r := range result {
        if r.ToString() != expected[i].ToString() {
            t.Errorf("Expected %s, got %s", expected[i].ToString(), r.ToString())
        }
    }
    memo := map[string][]PickedPatterns{}
    result = patterns.FindAllPermutations(target, picked, &memo)
    if len(result) != len(expected) {
        t.Errorf("Expected %d, got %d", len(expected), len(result))
    } else {
        for i, r := range result {
            if r.ToString() != expected[i].ToString() {
                t.Errorf("Expected %s, got %s", expected[i].ToString(), r.ToString())
            }
        }
    }

}

func TestFindAllPermutationsWithCacheFail(t *testing.T) {
    patterns := Patterns{
        patterns: []ColorPattern{
            {colors: []rune{'r'}, length: 1},
            {colors: []rune{'w', 'r'}, length: 2},
            {colors: []rune{'b'}, length: 1},
        },
    }

    target := Target{colors: []rune{'r'}}
    picked := PickedPatterns{"b"}
    expected := []PickedPatterns{
        {"b", "r"},
    }
    memo := map[string][]PickedPatterns{}
    // result := patterns.FindAllPermutations(target, picked, nil)
    result := patterns.FindAllPermutationsWithCache(target, picked, &memo)
    if len(result) != len(expected) {
        t.Errorf("Expected %d, got %d", len(expected), len(result))
    } else {
        for i, r := range result {
            if r.ToString() != expected[i].ToString() {
                t.Errorf("Expected %s, got %s", expected[i].ToString(), r.ToString())
            }
        }
    }
}

func TestPart2CachedVsUncached(t *testing.T) {
    ex2, _ := utils.ReadFile("test2.txt")
    patterns, targets := ParseInput(ex2)
    for _, target := range targets {
        memo := map[string][]PickedPatterns{}
        result1 := patterns.FindAllPermutations(target, PickedPatterns{}, nil)
        result2 := patterns.FindAllPermutations(target, PickedPatterns{}, &memo)
        for i, r := range result1 {
            if len(result2) <= i-1 {
                continue
            }
            if r.ToString() != result2[i].ToString() {
                t.Errorf("Expected %s, got %s", r.ToString(), result2[i].ToString())
            }
        }
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
            expected: 16,
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
