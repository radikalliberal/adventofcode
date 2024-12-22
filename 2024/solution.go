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
    "adventofcode/2024/day12"
    "adventofcode/2024/day13"
    "adventofcode/2024/day14"
    "adventofcode/2024/day15"
    "adventofcode/2024/day16"
    "adventofcode/2024/day17"
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
        case 12:
            fmt.Println("  Day 12")
            fmt.Println("    Part 1: ", day12.Part1(read_input(12, 1)))
            fmt.Println("    Part 2: ", day12.Part2(read_input(12, 2)))
        case 13:
            fmt.Println("  Day 13")
            fmt.Println("    Part 1: ", day13.Part1(read_input(13, 1)))
            fmt.Println("    Part 2: ", day13.Part2(read_input(13, 2)))
        case 14:
            fmt.Println("  Day 14")
            fmt.Println("    Part 1: ", day14.Part1(read_input(14, 1)))
            fmt.Println("    Part 2: ", day14.Part2(read_input(14, 2)))
        case 15:
            fmt.Println("  Day 15")
            fmt.Println("    Part 1: ", day15.Part1(read_input(15, 1)))
            fmt.Println("    Part 2: ", day15.Part2(read_input(15, 2)))
        case 16:
            fmt.Println("  Day 16")
            fmt.Println("    Part 1: ", day16.Part1(read_input(16, 1)))
            fmt.Println("    Part 2: ", day16.Part2(read_input(16, 2)))
        case 17:
            fmt.Println("  Day 17")
            fmt.Println("    Part 1: ", day17.Part1(read_input(17, 1)))
            fmt.Println("    Part 2: ", day17.Part2(read_input(17, 2)))
        default:
            fmt.Println("  Day ", day, " not implemented")

    }
}


func main() {
    pDay := flag.Int("day", 0, "Day to run")
    flag.Parse()
    Solution(*pDay)
}
