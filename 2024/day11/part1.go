// part1.go
package day11

import (
    "regexp"
    "strconv"
    "fmt"
    "math"
)

type Stone struct {
    val int
}

func PrintStones(stones []Stone) {
    var line string
    for _, stone := range stones {
        line += fmt.Sprintf("%d ", stone.val)
    }
    fmt.Println(line)
}

func (s Stone) Digits() int {
    // return len(strconv.Itoa(s.val))
    return int(math.Log10(float64(s.val))) + 1
}

func (s Stone) Split() []Stone {
    digits := s.Digits()
    left := Stone{val: s.val / int(math.Pow10(digits/2))}
    right := Stone{val: s.val % int(math.Pow10(digits/2))}
    return []Stone{left, right}
}

func (s Stone) Update() []Stone {
    // If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
    if s.val == 0 {
        s.val = 1
        return []Stone{s}
    }
    // If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
    if s.Digits() % 2 == 0 {
        return s.Split()
    }
    // If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.<F11>
    s.val *= 2024
    return []Stone{s}
}

func ParseInput(input string) []Stone{
    var stones []Stone
    re := regexp.MustCompile(`(\d+)`)
    for _, match := range re.FindAllString(input, -1) {
        val, _ := strconv.Atoi(match)
        stones = append(stones, Stone{val: val})
    }
    return stones
}

func Part1(input []string) int {
    stones := ParseInput(input[0])
    // PrintStones(stones)
    for i := 0; i < 25; i++ {
        newStones := []Stone{}
        for _, stone := range stones {
            newStones = append(newStones, stone.Update()...)
        }
        stones = newStones
        // PrintStones(stones)
    }
    return len(stones)
}
