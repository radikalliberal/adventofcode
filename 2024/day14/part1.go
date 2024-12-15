// part1.go
package day14

import (
    "fmt"
    "regexp"
    "strconv"
)

type Robot struct {
    x int
    y int
    vx int
    vy int
}

type Map struct {
    rows int
    cols int
}

func (r *Robot) Move(m Map) Robot {
    r.x += r.vx
    r.y += r.vy
    if r.x < 0 {
        r.x += m.cols
    }
    if r.x >= m.cols {
        r.x -= m.cols
    }
    if r.y < 0 {
        r.y += m.rows
    }
    if r.y >= m.rows {
        r.y -= m.rows
    }
    return *r
}

func (r Robot) Quadrant(m Map) int {
    if r.x < m.cols / 2 {
        if r.y < m.rows / 2 {
            return 0
        }
        if r.y >= m.rows / 2 + 1 {
            return 2
        }
    }
    if r.x >= m.cols / 2 + 1 {
        if r.y < m.rows / 2 {
            return 1
        }
        if r.y >= m.rows / 2 + 1 {
            return 3
        }
    }
    return -1
}

func ParseLine(line string) Robot {
    re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
    matches := re.FindStringSubmatch(line)
    x, _ := strconv.Atoi(matches[1])
    y, _ := strconv.Atoi(matches[2])
    vx, _ := strconv.Atoi(matches[3])
    vy, _ := strconv.Atoi(matches[4])
    return Robot{x: x, y: y, vx: vx, vy: vy}
}

func MaxDims(robots []Robot) (int, int) {
    max_cols := 0
    max_rows := 0
    for _, robot := range robots {
        if robot.x > max_cols {
            max_cols = robot.x
        }
        if robot.y > max_rows {
            max_rows = robot.y
        }
    }
    return max_rows, max_cols
}

func ParseInput(input []string) ([]Robot, Map) {
    robots := []Robot{}
    for _, line := range input {
        robots = append(robots, ParseLine(line))
    }
    max_rows, max_cols:= MaxDims(robots)
    // 101, 103 if large
    // 7, 11 if small
    if max_cols <= 11 && max_rows <= 7 {
        return robots, Map{7, 11}
    }
    return robots, Map{103, 101}
}

func ComputeRobotScore(robots []Robot, m Map) int {
    score := 1
    quadrants := []int{0, 0, 0, 0}
    for _, robot := range robots {
        q := robot.Quadrant(m)
        if q != -1 {
            quadrants[q]++
        }
    }
    for _, q := range quadrants {
        score *= q
    }
    return score
}

func PrintRobots(robots []Robot, m Map) {
    Map := make([][]int, m.rows)
    for y := 0; y < m.rows; y++ {
        Map[y] = make([]int, m.cols)
        for x := 0; x < m.cols; x++ {
            for _, robot := range robots {
                if robot.x == x && robot.y == y {
                    Map[y][x] += 1
                }
            }
        }
    }
    for y := 0; y < m.rows; y++ {
        for x := 0; x < m.cols; x++ {
            if Map[y][x] > 0 {
                fmt.Printf("%d", Map[y][x])
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
    fmt.Println()
}

func Part1(input []string) int {
    robots, m := ParseInput(input)
    // PrintRobots(robots, m)
    for range 100 {
        for i := range robots {
            robots[i] = robots[i].Move(m)
        }
    }
    // PrintRobots(robots, m)
    return ComputeRobotScore(robots, m)
}
