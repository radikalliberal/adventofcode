// part1.go
package day02
import (
    "strconv"
    "strings"
    "math"
)

func ParseLine(line string) []int {
    parts := strings.Split(line, " ")
    nums := make([]int, len(parts))
    for i, part := range parts {
        nums[i], _ = strconv.Atoi(part)
    }
    return nums
}

func DistSave(num1 int, num2 int) bool {
    abs_first_dist := int(math.Abs(float64(num1 - num2)))
    return abs_first_dist >= 1 && abs_first_dist <= 3
}

func EvalSafe(nums []int) int {
    first_dist := nums[1] - nums[0]
    if !DistSave(nums[0], nums[1]) {
        return 0
    }
    is_increasing := first_dist > 0
    last_num := nums[0]
    for _, num := range nums[1:] {
        dist := num - last_num
        if is_increasing && dist < 0 || !is_increasing && dist > 0 {
            return 0
        }
        if !DistSave(last_num, num) {
            return 0
        }
        last_num = num
    }
    return 1
}

func Part1(input []string) int {
    num_safe := 0
    for _, line := range input {
        num_safe += EvalSafe(ParseLine(line))
    }
    return num_safe
}
