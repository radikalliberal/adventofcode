package main
import (
	"fmt"
	"adventofcode/utils"
	"adventofcode/2023/day01"
)

func Solution() {
	fmt.Println("Advent of Code 2023 - Day 1")
	input1, err := utils.ReadFile("2023/day01/input1.txt")
	if err != nil {
		fmt.Println("Error reading input: ", err)
	}
	fmt.Println("Part 1: ", day01.Part1(input1))
	input2, err := utils.ReadFile("2023/day01/input2.txt")
	if err != nil {
		fmt.Println("Error reading input: ", err)
	}
	fmt.Println("Part 2: ", day01.Part2(input2))
}


func main() {
	Solution()
}
