package day01

import (
	"fmt"
	"strings"
	// "log/slog"
)

func sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func get_min_idx(nums []int) int {
	min := 0
	for idx, num := range nums {
		if nums[min] == -1 {
			min = idx
		}
		if num < nums[min] && num != -1 {
			min = idx
		}
	}
	return min
}

func reverse(s string) string {
	var reversed string
	for _, char := range s {
		reversed = string(char) + reversed
	}
	return reversed
}

func substitute_first(line string, nums []Number) string {
	indecies := []int{}
	for _, num := range nums {
		indecies = append(indecies, strings.Index(line, num.name))
	}
	min_idx := get_min_idx(indecies)
	line = strings.Replace(line, nums[min_idx].name, fmt.Sprintf("%d", nums[min_idx].num), 1)
	return line
}

type Number struct {
	name string
	num int
}

func reverse_num(num Number) Number {
	return Number{ reverse(num.name), num.num }
}

func replace(line string) string {
	nums := []Number{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}
	result := substitute_first(line, nums)
	var reversed_nums []Number
	for _, num := range nums {
		reversed_nums = append(reversed_nums, reverse_num(num))
	}
	reversed_line := reverse(line)

	reversed_line = substitute_first(reversed_line, reversed_nums)
	reversed_line = reverse(reversed_line)
	// slog.Info(line)
	// slog.Info(result)
	return result + reversed_line
}

func Part2(input []string) int {
	var sum int
	for _, line := range input {
		sum += compute_line(replace(line))
	}
	return sum
}
