// part2.go
package day16

import (
    "container/heap"
)

func Fitness(costMatrix [][]int) int {
    sum := 0
    for i := 0; i < len(costMatrix); i++ {
        for j := 0; j < len(costMatrix[i]); j++ {
            sum += costMatrix[i][j]
        }
    }
    return sum
}


func (r Reindeer) MarkUsedFields(visited *[][]bool) {
    r.row = r.m.start.row
    r.col = r.m.start.col
    r.orientation = EAST
    (*visited)[r.row][r.col] = true
    prev := EAST
    for _, move := range r.path {
        if move != prev {
            prev = move
            continue
        }
        prev = move
        r = r.Move(move)
        (*visited)[r.row][r.col] = true
    }
}

func SumVisitedFields(visited [][]bool) int {
    traveledFields := 0
    for i := 0; i < len(visited); i++ {
        for j := 0; j < len(visited[i]); j++ {
            if visited[i][j] {
                traveledFields++
            }
        }
    }
    return traveledFields
}

func ComputeTraveledFields(reindeers []*Reindeer) int {
    minLen := 10000000
    var minPaths []*Reindeer
    for _, rd := range reindeers {
        if len(rd.path) < minLen {
            minLen = len(rd.path)
        }
    }
    for _, rd := range reindeers {
        if len(rd.path) == minLen {
            minPaths = append(minPaths, rd)
        }
    }
    visited := make([][]bool, reindeers[0].m.rows)
    for i := 0; i < reindeers[0].m.rows; i++ {
        visited[i] = make([]bool, reindeers[0].m.cols)
    }
    for _, rd := range minPaths {
        rd.MarkUsedFields(&visited)
        // fmt.Printf("Path %s visited: %d\n", string(rd.path), SumVisitedFields(visited))
    }
    return SumVisitedFields(visited)
}

func (m *Map) ComputeMazePart2() int {
    costMap := make([][]int, m.rows)
    for i := 0; i < m.rows; i++ {
        costMap[i] = make([]int, m.cols)
        for j := 0; j < m.cols; j++ {
            costMap[i][j] = 10000000
        }
    }
    var valid_solutions []*Reindeer
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    rd := &m.r
    item := &Item{rd, 0, 0}
    heap.Push(&pq, item)
    for len(pq) > 0 {
        for _, move := range rd.PossibleMoves() {
            cost := move.score
            if costMap[move.row][move.col] + 1001 < cost {
                continue
            }
            if costMap[move.row][move.col] > cost {
                costMap[move.row][move.col] = cost
            }
            // fmt.Println("Move", string(move.path), "Score", cost)
            heap.Push(&pq, &Item{ r: &move, priority: -cost})
        }

        item = heap.Pop(&pq).(*Item)
        rd = item.r

        for rd.row == m.end.row && rd.col == m.end.col {
            valid_solutions = append(valid_solutions, rd)
            item = heap.Pop(&pq).(*Item)
            rd = item.r
        }
    }

    numFields := ComputeTraveledFields(valid_solutions)
    rd.PrintPath()
    return numFields
}

func Part2(input []string) int {
    m := ParseInputPart1(input)
    // m.Print()
    return m.ComputeMazePart2()
}
