// part1.go
package day03

import (
    "regexp"
    "strconv"
)

func Part1(input []string) int {
    result := 0
    r := regexp.MustCompile(`mul\(\d+,\d+\)`)
    for _, line := range input {
        matches := r.FindAllString(line, -1)
        for _, match := range matches {
            nums := regexp.MustCompile(`\d+`).FindAllString(match, -1)

            n1, _ := strconv.Atoi(nums[0])
            n2, _ := strconv.Atoi(nums[1])
            result += n1 * n2
        }
    }
    return result
}
