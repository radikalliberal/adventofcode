// part1.go
package day08

import (
	"fmt"
)

type pos struct {
    x int
    y int
}

type loc struct {
    symbol rune
    antinode bool
    noAntenna bool
    p pos
}

func (l loc) String() string {
    if l.IsAntinode() {
        return fmt.Sprintf("\033[31m%c\033[0m", l.symbol)
    }
    return string(l.symbol)
}

func (l loc) IsAntinode() bool {
    return l.antinode
}

func (p pos) String() string {
    return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func (p pos) IsInsideMap(locMap [][]loc) bool {
    return p.x >= 0 && p.x < len(locMap) && p.y >= 0 && p.y < len(locMap[0])
}

func (p pos) Equals(p2 pos) bool {
    return p.x == p2.x && p.y == p2.y
}

func ParseInput(input []string) ([][]loc, map[rune][]*loc) {
    antennas := make(map[rune][]*loc)
    locMap := make([][]loc, len(input))
    for i, line := range input {
        locMap[i] = make([]loc, len(line))
        for j, c := range line {
            if c != '.' {
                locMap[i][j] = loc{symbol: c, antinode: false, noAntenna: false, p: pos{x: i, y: j}}
                antennas[c] = append(antennas[c], &locMap[i][j])
            } else {
        locMap[i][j] = loc{symbol: '.', antinode: false, noAntenna: true, p: pos{x: i, y: j}}
            }
        }
    }
    return locMap, antennas
}

func SetPossibleAntinode(a, b loc, locMap [][]loc) {
    x_dist := a.p.x - b.p.x
    y_dist := a.p.y - b.p.y
    anode1 := pos{x: a.p.x + x_dist, y: a.p.y + y_dist}
    if anode1.IsInsideMap(locMap) {
        if !anode1.Equals(b.p) {
            locMap[anode1.x][anode1.y].antinode = true
        }
    }

    anode2 := pos{x: a.p.x - x_dist, y: a.p.y - y_dist}
    if anode2.IsInsideMap(locMap) {
        if !anode2.Equals(b.p) {
            locMap[anode2.x][anode2.y].antinode = true
        }
    }
}

func CountAntinodes(locMap [][]loc) int {
    count := 0
    for _, row := range locMap {
        for _, l := range row {
            if l.antinode {
                count++
            }
        }
    }
    return count
}

func Part1(input []string) int {
    locMap, antennas := ParseInput(input)
    for _, ant := range antennas {
        for idx, a := range ant {
            // compare against all previous antennas
            for i := 0; i < idx; i++ {
                SetPossibleAntinode(*a, *ant[i], locMap)
            }
            for i := idx + 1; i < len(ant); i++ {
                SetPossibleAntinode(*a, *ant[i], locMap)
            }
        }
    }
    // PrintMap(locMap)
    return CountAntinodes(locMap)
}
