// part2.go
package day03
import (
    "regexp"
    "strconv"
)

func Part2(input []string) int {
    result := 0
    r := regexp.MustCompile(`(do\(\)|don't\(\)|mul\(\d+,\d+\))`)
    enabled := true
    for _, line := range input {
        matches := r.FindAllString(line, -1)
        for _, match := range matches {
            nums := regexp.MustCompile(`\d+`).FindAllString(match, -1)
            if match == "do()" {
                enabled = true
            } else if match == "don't()" {
                enabled = false
            } else if enabled {
                n1, _ := strconv.Atoi(nums[0])
                n2, _ := strconv.Atoi(nums[1])
                result += n1 * n2
            }
        }
    }
    return result
}
