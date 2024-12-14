// part1.go
package day10

type Point struct {
    x int
    y int
}

func ParseInput(input []string) [][]int {
    trailMap := make([][]int, len(input))
    for k, line := range input {
        row := make([]int, len(line))
        for i, c := range line {
            if c == '.' {
                row[i] = -1
            } else {
                row[i] = int(c - '0')
            }
        }
        trailMap[k] = append(trailMap[k], row...)

    }
    return trailMap
}

func computeTrailheads(trailMap [][]int) []Point{
    var trailheads []Point
    for y, row := range trailMap {
        for x, val := range row {
            if val == 0 {
                trailheads = append(trailheads, Point{x, y})
            }
        }
    }
    return trailheads
}

func computePaths(trailMap [][]int, trailhead Point) []Point {
    count := []Point{}
    cur_pos := trailhead
    cur_height := trailMap[cur_pos.y][cur_pos.x]
    bounds := Point{len(trailMap[0]), len(trailMap)}
    if cur_height == 9 {
        return []Point{cur_pos}
    }
    if cur_pos.x < 0 || cur_pos.x >= bounds.x || cur_pos.y < 0 || cur_pos.y >= bounds.y {
        return []Point{}
    }
    if cur_pos.y+1 < bounds.y && trailMap[cur_pos.y+1][cur_pos.x] == cur_height + 1 {
        count = append(count, computePaths(trailMap, Point{cur_pos.x, cur_pos.y+1})...)
    }
    if cur_pos.x > 0 && trailMap[cur_pos.y][cur_pos.x-1] == cur_height + 1 {
        count = append(count, computePaths(trailMap, Point{cur_pos.x-1, cur_pos.y})...)
    }
    if cur_pos.x+1 < bounds.x && trailMap[cur_pos.y][cur_pos.x+1] == cur_height + 1 {
        count = append(count, computePaths(trailMap, Point{cur_pos.x+1, cur_pos.y})...)
    }
    if cur_pos.y > 0 && trailMap[cur_pos.y-1][cur_pos.x] == cur_height + 1 {
        count = append(count, computePaths(trailMap, Point{cur_pos.x, cur_pos.y-1})...)
    }
    return count
}

func (p Point) Equal(q Point) bool {
    return p.x == q.x && p.y == q.y
}

func NumUnique(points []Point) int {
    numNotUnique := 0
    for i, p1 := range points {
        for _, p2 := range points[i+1:] {
            if p1.Equal(p2) {
                numNotUnique++
                break
            }
        }
    }
    return len(points) - numNotUnique
}

func Part1(input []string) int {
    trailMap := ParseInput(input)
    trailheads := computeTrailheads(trailMap)
    numPaths := 0
    for _, trailhead := range trailheads {

        reachableEnds := computePaths(trailMap, trailhead)
        numPaths += NumUnique(reachableEnds)
    }
    return numPaths
}
