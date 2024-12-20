// part1.go
package day16

import (
    // "fmt"
    "container/heap"
)

const (
    EMPTY = '.'
    WALL = '#'
    START = 'S'
    END = 'E'
    NORTH = '^'
    SOUTH = 'v'
    EAST = '>'
    WEST = '<'
)

var color = map[rune]string{
    EMPTY: "\033[0m",
    WALL: "\033[31m",
    START: "\033[32m",
    END: "\033[33m",
    NORTH: "\033[34m",
    SOUTH: "\033[35m",
    EAST: "\033[36m",
    WEST: "\033[37m",
}

type Reindeer struct {
    row int
    col int
    orientation rune
    path []rune
    m *Map
    score int
}

type Path struct {
    p []rune
    score int
    pos Pos
}

type Pos struct {
    row int
    col int
}

type Map struct {
    m [][]rune
    rows int
    cols int
    r Reindeer
    start Pos
    end Pos
}

func CalculateCost(path []rune) int {
    score := 0
    prev := EAST
    for _, r := range path {
        if r != prev {
            score += 1000
        } else {
            score += 1
        }
        prev = r
    }
    return score
}

func ParseInputPart1(input []string) Map {
    m := make([][]rune, len(input))
    rd := Reindeer{}
    start := Pos{}
    end := Pos{}
    for row, line := range input {
        m[row] = []rune(line)
        for col, r := range m[row] {
            if r == START {
                start = Pos{row: row, col: col}
                rd = Reindeer{row: row, col: col, orientation: EAST, path: []rune{}, m: nil}
            }
            if r == END {
                end = Pos{row: row, col: col}
            }
        }
    }
    newMap := Map{m: m, rows: len(input), cols: len(input[0]), r: rd, start: start, end: end}
    rd.m = &newMap
    newMap.r = rd
    return newMap
}

func (m *Map) Print() {
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            c := m.m[i][j]
            print(color[m.m[i][j]], string(c))
        }
        println()
    }
    print(color[EMPTY])
}

func (r *Reindeer) Copy() Reindeer {
    path := make([]rune, len(r.path))
    copy(path, r.path)
    return Reindeer{row: r.row, col: r.col, orientation: r.orientation, path: path, m: r.m, score: r.score}
}

func (rOrig *Reindeer) Move(move rune) Reindeer {
    r := rOrig.Copy()
    r.path = append(r.path, move)
    switch move {
        case NORTH:
            r.row--
        case SOUTH:
            r.row++
        case EAST:
            r.col++
        case WEST:
            r.col--
    }
    r.score += 1
    return r
}

func (r *Reindeer) Turn(turn rune) Reindeer {
    cpy := r.Copy()
    switch turn {
    case 'L':
        switch r.orientation {
        case NORTH:
            cpy.path = append(cpy.path, WEST)
        case SOUTH:
            cpy.path = append(cpy.path, EAST)
        case EAST:
            cpy.path = append(cpy.path, NORTH)
        case WEST:
            cpy.path = append(cpy.path, SOUTH)
        }
    case 'R':
        switch r.orientation {
        case NORTH:
            cpy.path = append(cpy.path, EAST)
        case SOUTH:
            cpy.path = append(cpy.path, WEST)
        case EAST:
            cpy.path = append(cpy.path, SOUTH)
        case WEST:
            cpy.path = append(cpy.path, NORTH)
        }
    }
    cpy.orientation = cpy.path[len(cpy.path) - 1]
    cpy.score += 1000
    return cpy
}

func (r *Reindeer) PossibleMoves() []Reindeer {
    moves := []Reindeer{}

    switch r.orientation {
    case NORTH:
        if r.m.GetField(Pos{r.row - 1, r.col}) != WALL {
            moves = append(moves, r.Move(NORTH))
        }
    case SOUTH:
        if r.m.GetField(Pos{r.row + 1, r.col}) != WALL {
            moves = append(moves, r.Move(SOUTH))
        }
    case EAST:
        if r.m.GetField(Pos{r.row, r.col + 1}) != WALL {
            moves = append(moves, r.Move(EAST))
        }
    case WEST:
        if r.m.GetField(Pos{r.row, r.col - 1}) != WALL {
            moves = append(moves, r.Move(WEST))
        }
    }
    moves = append(moves, r.Turn('L'))
    moves = append(moves, r.Turn('R'))
    return moves
}

func (m *Map) GetField(p Pos) rune {
    if p.row < 0 || p.row >= m.rows || p.col < 0 || p.col >= m.cols {
        return WALL
    }
    return m.m[p.row][p.col]
}

// func (m *Map) GetReindeer(p Path) Reindeer {
//     r := Reindeer{row: m.start.row, col: m.start.col, orientation: EAST, path: p, m: m}
//     pos := r.ComputePosition(p.p)
//     r.row = pos.row
//     r.col = pos.col
//     if len(p.p) > 0 {
//         r.orientation = p.p[len(p.p) - 1]
//     }
//     return r
// }

func (m *Map) ComputeMaze() {
    costMap := make([][]int, m.rows)
    for i := 0; i < m.rows; i++ {
        costMap[i] = make([]int, m.cols)
        for j := 0; j < m.cols; j++ {
            costMap[i][j] = 10000000
        }
    }
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    rd := m.r
    item := &Item{&rd, 0, 0}
    heap.Push(&pq, item)
    iteration := 0
    for rd.row != m.end.row || rd.col != m.end.col {
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
        rd = *(item.r)
        // rd.PrintPath()
        iteration++
        // if iteration > 10000 && iteration % 1000 == 0 {
        //     fmt.Println("Iteration", iteration, "Queue size", len(pq), "Score", item.priority)
        // }
    }
    m.r = *(item.r)
    m.r.PrintPath()
}

func Contains(positions []Pos, p Pos) bool {
    for _, pos := range positions {
        if pos.row == p.row && pos.col == p.col {
            return true
        }
    }
    return false
}

func CheckPath(p Path) bool {
    prev := EAST
    pos := Pos{row: 0, col: 0}
    history := []Pos{}
    for _, r := range p.p {
        if r != prev {
        } else {
            switch r {
            case NORTH:
                pos.row--
            case SOUTH:
                pos.row++
            case EAST:
                pos.col++
            case WEST:
                pos.col--
            }
            if !Contains(history, pos) {
                history = append(history, pos)
            } else {
                return false
            }
        }
        prev = r
    }
    return true
}

func (m *Map) Copy() Map {
    mNew := make([][]rune, m.rows)
    for i := 0; i < m.rows; i++ {
        mNew[i] = make([]rune, m.cols)
        copy(mNew[i], m.m[i])
    }
    return Map{m: mNew, rows: m.rows, cols: m.cols, r: m.r, start: m.start, end: m.end}
}

func (m *Map) SetField(p Pos, r rune) {
    m.m[p.row][p.col] = r
}

func (r *Reindeer) ComputePosition(path []rune) Pos {
    pos := Pos{row: r.row, col: r.col}
    prevMove := EAST
    for _, p := range path {
        if p != prevMove {
            prevMove = p
            continue
        }
        switch p {
        case NORTH:
            pos.row--
        case SOUTH:
            pos.row++
        case EAST:
            pos.col++
        case WEST:
            pos.col--
        }
        prevMove = p
    }
    return pos
}

func (r *Reindeer) PrintPath() {
    cpMap := r.m.Copy()
    prevMove := '.'
    for _, p := range r.path {
        if p == prevMove {
            cpMap.r = cpMap.r.Move(p)
        }
        cpMap.SetField(Pos{cpMap.r.row, cpMap.r.col}, p)
        prevMove = p
    }
    // println("Path: Score", CalculateCost(cpMap.r.path))
    cpMap.Print()
}


func Part1(input []string) int {
    m := ParseInputPart1(input)
    // m.Print()
    m.ComputeMaze()
    return CalculateCost(m.r.path)
}
