// part1.go
package day01

import (
    "strconv"
    "strings"
    "sort"
    "math"
)

func parseLine(line string) (int, int) {
    parts := strings.Split(line, " ")
    left, _ := strconv.Atoi(parts[0])
    var right int
    for i := 1; i < len(parts); i++ {
        if parts[i] == "" {
            continue
        }
        right, _ = strconv.Atoi(parts[i])
    }
    return left, right
}

func parseInput(input []string) ([]int, []int) {
    left_list := make([]int, len(input))
    right_list := make([]int, len(input))
    for i, line := range input {
        left_list[i], right_list[i] = parseLine(line)
    }
    sort.Ints(left_list)
    sort.Ints(right_list)
    return left_list, right_list
}

func compute_distance(l []int, r []int) int {
    distance := 0
    for i := 0; i < len(l); i++ {
        distance += int(math.Abs(float64(r[i] - l[i])))
    }
    return distance
}

func Part1(input []string) int {
    l, r := parseInput(input)
    return compute_distance(l, r)
}
