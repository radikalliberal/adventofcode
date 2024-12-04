package main
import (
    "fmt"
    "adventofcode/utils"
    "adventofcode/2024/day01"
    "adventofcode/2024/day02"
    "adventofcode/2024/day03"
    "adventofcode/2024/day04"
)

func read_input(day int, part int) []string {
    input, err := utils.ReadFile("2024/day" + fmt.Sprintf("%02d", day) + "/input" + fmt.Sprintf("%d", part) + ".txt")
    if err != nil {
        fmt.Println("Error reading input: ", err)
    }
    if len(input) == 0 {
        fmt.Println("Error: empty input")
    }
    return input
}

func Solution() {
    fmt.Println("Advent of Code 2024")
    fmt.Println("  Day 1")
    fmt.Println("    Part 1: ", day01.Part1(read_input(1, 1)))
    fmt.Println("    Part 2: ", day01.Part2(read_input(1, 2)))
    fmt.Println("  Day 02")
    fmt.Println("    Part 1: ", day02.Part1(read_input(2, 1)))
    fmt.Println("    Part 2: ", day02.Part2(read_input(2, 2)))
    fmt.Println("  Day 03")
    fmt.Println("    Part 1: ", day03.Part1(read_input(3, 1)))
    fmt.Println("    Part 2: ", day03.Part2(read_input(3, 2)))
    fmt.Println("  Day 04")
    fmt.Println("    Part 1: ", day04.Part1(read_input(4, 1)))
    fmt.Println("    Part 2: ", day04.Part2(read_input(4, 2)))
}


func main() {
    Solution()
}
