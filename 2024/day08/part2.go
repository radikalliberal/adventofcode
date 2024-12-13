// part2.go
package day08

import (
    "fmt"
)

func SetPossibleAntinodePart2(a, b loc, locMap [][]loc) {
    x_dist := a.p.x - b.p.x
    y_dist := a.p.y - b.p.y
    x_dist_step := x_dist
    y_dist_step := y_dist

    for {
        anode1 := pos{x: a.p.x + x_dist, y: a.p.y + y_dist}
        anode2 := pos{x: a.p.x - x_dist, y: a.p.y - y_dist}
        node1InMap := anode1.IsInsideMap(locMap)
        node2InMap := anode2.IsInsideMap(locMap)
        if node1InMap {
            locMap[anode1.x][anode1.y].antinode = true
        }

        if node2InMap {
            locMap[anode2.x][anode2.y].antinode = true
        }
        x_dist += x_dist_step
        y_dist += y_dist_step
        if !node1InMap && !node2InMap {
            break
        }
    }
}

func PrintMap(locMap [][]loc) {
    for _, row := range locMap {
        for _, l := range row {
            fmt.Print(l)
        }
        fmt.Println()
    }
}

func Part2(input []string) int {
    locMap, antennas := ParseInput(input)
    for _, ant := range antennas {
        for idx, a := range ant {
            // compare against all previous antennas
            for i := 0; i < idx; i++ {
                SetPossibleAntinodePart2(*a, *ant[i], locMap)
            }
            for i := idx + 1; i < len(ant); i++ {
                SetPossibleAntinodePart2(*a, *ant[i], locMap)
            }
        }
    }
    // PrintMap(locMap)
    return CountAntinodes(locMap)
}
