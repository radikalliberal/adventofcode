package main
import (
    "fmt"
    "flag"
    "adventofcode/utils"
    "adventofcode/2024/day01"
    "adventofcode/2024/day02"
    "adventofcode/2024/day03"
    "adventofcode/2024/day04"
    "adventofcode/2024/day05"
    "adventofcode/2024/day06"
    "adventofcode/2024/day07"
    "adventofcode/2024/day08"
    "adventofcode/2024/day09"
    "adventofcode/2024/day10"
    "adventofcode/2024/day11"
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

func Solution(day int) {
    fmt.Println("Advent of Code 2024")
    if day == 0 {
        fmt.Println("  All days")
        for i := 1; i <= 11; i++ {
            Solution(i)
        }
        return
    }

    switch day {
        case 1:
            fmt.Println("  Day 1")
            fmt.Println("    Part 1: ", day01.Part1(read_input(1, 1)))
            fmt.Println("    Part 2: ", day01.Part2(read_input(1, 2)))
        case 2:
            fmt.Println("  Day 02")
            fmt.Println("    Part 1: ", day02.Part1(read_input(2, 1)))
            fmt.Println("    Part 2: ", day02.Part2(read_input(2, 2)))
        case 3:
            fmt.Println("  Day 03")
            fmt.Println("    Part 1: ", day03.Part1(read_input(3, 1)))
            fmt.Println("    Part 2: ", day03.Part2(read_input(3, 2)))
        case 4:
            fmt.Println("  Day 04")
            fmt.Println("    Part 1: ", day04.Part1(read_input(4, 1)))
            fmt.Println("    Part 2: ", day04.Part2(read_input(4, 2)))
        case 5:
            fmt.Println("  Day 05")
            fmt.Println("    Part 1: ", day05.Part1(read_input(5, 1)))
            fmt.Println("    Part 2: ", day05.Part2(read_input(5, 2)))
        case 6:
            fmt.Println("  Day 06")
            fmt.Println("    Part 1: ", day06.Part1(read_input(6, 1)))
            fmt.Println("    Part 2: ", day06.Part2(read_input(6, 2)))
        case 7:
            fmt.Println("  Day 07")
            fmt.Println("    Part 1: ", day07.Part1(read_input(7, 1)))
            fmt.Println("    Part 2: ", day07.Part2(read_input(7, 2)))
        case 8:
            fmt.Println("  Day 08")
            fmt.Println("    Part 1: ", day08.Part1(read_input(8, 1)))
            fmt.Println("    Part 2: ", day08.Part2(read_input(8, 2)))
        case 9:
            fmt.Println("  Day 09")
            fmt.Println("    Part 1: ", day09.Part1(read_input(9, 1)))
            fmt.Println("    Part 2: ", day09.Part2(read_input(9, 2)))
            fmt.Println("    Part 2_2: ", day09.Part2_2(read_input(9, 2)))
        case 10:
            fmt.Println("  Day 10")
            fmt.Println("    Part 1: ", day10.Part1(read_input(10, 1)))
            fmt.Println("    Part 2: ", day10.Part2(read_input(10, 2)))
        case 11:
            fmt.Println("  Day 11")
            fmt.Println("    Part 1: ", day11.Part1(read_input(11, 1)))
            fmt.Println("    Part 2: ", day11.Part2(read_input(11, 2)))
        default:
            fmt.Println("  Day ", day, " not implemented")

    }
}


func main() {
    pDay := flag.Int("day", 0, "Day to run")
    flag.Parse()
    Solution(*pDay)
}
