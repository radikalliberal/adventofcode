package day01

import (
	"unicode"
)

func compute_line(line string) int {
	var numbers []int
	for _, char := range line {
		if unicode.IsDigit(char) {
			numbers = append(numbers, int(char - '0'))
		}
	}
	return numbers[0] * 10 + numbers[len(numbers)-1]
}

func Part1(input []string) int {
	var sum int
	for _, line := range input {
		sum += compute_line(line)
	}
	return sum
}
