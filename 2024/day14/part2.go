// part2.go
package day14

import (
    "fmt"
    "math"
)

func ComputeNeighbours(robots []Robot) int {
    neighbours := 0
    for i := 0; i < len(robots); i++ {
        for j := i + 1; j < len(robots); j++ {
            x_dist := int(math.Abs(float64(robots[i].x - robots[j].x)))
            y_dist := int(math.Abs(float64(robots[i].y - robots[j].y)))
            if x_dist <= 1 && y_dist <= 1 {
                neighbours++
            }
        }
    }
    return neighbours
}

func Part2(input []string) int {
    robots, m := ParseInput(input)
    max_neighbours := 0
    for i := range 8280 {
        for i := range robots {
            robots[i] = robots[i].Move(m)
        }
        neighbours := ComputeNeighbours(robots)
        if neighbours > max_neighbours {
            fmt.Println("After", i+1, "neighbours:", neighbours)
            PrintRobots(robots, m)
            fmt.Println()
            max_neighbours = neighbours
        }
    }
    return max_neighbours
}
