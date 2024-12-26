// part1.go
package day18

import (
    "container/heap"
    "log/slog"
    "regexp"
    "math"
    "strconv"
)

type Coord struct {
    row int
    col int
}

type Node struct {
    cost int
    pos Coord
    isCorrupted bool
}

type Path struct {
    _nodes *Map
    history []Coord
    pos Coord
    cost int
}

type CostMatrix [][]int

func CreateCostMatrix(rows, cols, defaultValue int) CostMatrix {
    matrix := make([][]int, rows)
    for row := 0; row < rows; row++ {
        matrix[row] = make([]int, cols)
        for col := 0; col < cols; col++ {
            matrix[row][col] = defaultValue
        }
    }
    return matrix
}

func (c *CostMatrix) Get(pos Coord) int {
    return (*c)[pos.row][pos.col]
}

func (c *CostMatrix) Set(pos Coord, value int) {
    (*c)[pos.row][pos.col] = value
}

func ParseInput(input []string) (Map, []Coord) {
    // small input == 12 lines
    maxRows := 7
    maxCols := 7
    readMax := 12
    if len(input) > 1000 {
        maxRows = 71
        maxCols = 71
        readMax = 1024
    }
    corruptedMemory := make([]Coord, 0)
    for _, line := range input {
        re := regexp.MustCompile("([0-9]+),([0-9]+)")
        matches := re.FindStringSubmatch(line)
        x, _ := strconv.Atoi(matches[1])
        y, _ := strconv.Atoi(matches[2])
        corruptedMemory = append(corruptedMemory, Coord{y,x})
    }
    nodes := make([][]Node, maxRows)
    for row := 0; row < maxRows; row++ {
        nodes[row] = make([]Node, maxCols)
        for col := 0; col < maxCols; col++ {
            found := false
            for _, coord := range corruptedMemory[0:readMax] {
                if coord.row == row && coord.col == col {
                    found = true
                    break
                }
            }
            if found {
                nodes[row][col] = Node{
                    cost: 9999,
                    pos: Coord{row, col},
                    isCorrupted: true}
            } else {
                nodes[row][col] = Node{
                    cost: maxRows - row + maxCols - col,
                    pos: Coord{row, col},
                    isCorrupted: false}
            }
        }
    }
    return nodes, corruptedMemory[readMax:]
}

func (path Path) PossiblePaths() []Path {
    paths := make([]Path, 0)
    for _, dRow := range []int{-1, 0, 1} {
        for _, dCol := range []int{-1, 0, 1} {
            if dRow == 0 && dCol == 0 {
                continue
            }
            if math.Abs(float64(dRow)) + math.Abs(float64(dCol)) == 2 {
                continue
            }
            newRow := path.pos.row + dRow
            newCol := path.pos.col + dCol
            newPos := Coord{newRow, newCol}
            if path.InHistory(newPos) {
                continue
            }
            if newRow >= 0 && newRow < len(*path._nodes) && newCol >= 0 && newCol < len((*path._nodes)[0]) {
                if (*path._nodes)[newRow][newCol].isCorrupted {
                    continue
                }
                newPath := path.NewPath(newPos)

                if !newPath.Check() {
                    panic("Invalid path")
                }
                paths = append(paths, newPath)
            }
        }
    }
    return paths
}

func (p *Path) Copy() Path {
    history := make([]Coord, len(p.history))
    copy(history, p.history)
    return Path{
        _nodes: p._nodes,
        history: history,
        pos: p.pos,
        cost: p.cost,
    }
}

func (p *Path) NewPath(pos Coord) Path {
    copy := p.Copy()
    copy.history = append(copy.history, pos)
    copy.pos = pos
    copy.cost = (*p._nodes)[pos.row][pos.col].cost
    return copy
}

func (p *Path) Check() bool {
    for idx, coord := range p.history {
        if idx == 0 {
            continue
        }
        prev := p.history[idx - 1]
        if math.Abs(float64(coord.row - prev.row)) + math.Abs(float64(coord.col - prev.col)) != 1 {
            return false
        }
    }
    if p.pos.row != p.history[len(p.history) - 1].row || p.pos.col != p.history[len(p.history) - 1].col {
        return false
    }
    return true
}


func FindShortestPath(nodes *Map, start Coord, end Coord) int {
    startPath := Path{
        _nodes: nodes,
        history: []Coord{start},
        pos: start,
        cost: (*nodes)[start.row][start.col].cost,
    }
    cm := CreateCostMatrix(len(*nodes), len((*nodes)[0]), 9999999)
    startPath.Print()
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    item := &Item{&startPath, 0, 0}
    heap.Push(&pq, item)
    iteration := 0
    for len(pq) > 0 {
        item = heap.Pop(&pq).(*Item)
        path := item.p
        if !path.Check() {
            panic("Invalid path")
        }
        if path.pos.row == end.row && path.pos.col == end.col {
            path.Print()
            return len(path.history) - 1
        }
        paths := path.PossiblePaths()
        for _, p := range paths {
            if cm.Get(p.pos) >= len(p.history) {
                heap.Push(&pq, &Item{ p: &p, priority: -(p.cost * len(p.history))})
                cm.Set(p.pos, len(p.history))
            }
        }
        iteration += 1
        if iteration % 100000 == 0 {
            path.Print()
        }
    }
    return -1
}

func (p *Path) InHistory(coord Coord) bool {
    for _, c := range p.history {
        if c.row == coord.row && c.col == coord.col {
            return true
        }
    }
    return false
}

func (p *Path) Print() {
    out := ""
    m := p._nodes
    for r, row := range *m {
        for c, node := range row {
            pos := Coord{r, c}
            if node.isCorrupted {
                if p.InHistory(pos) {
                    panic("Corrupted node in path")
                }
                // ansi colors
                out += "\033[31m#\033[0m"
            } else {
                if p.InHistory(pos) {
                    out += "\033[32mO\033[0m"
                } else {
                    out += "."
                }
            }
        }
        out += "\n"
    }
    slog.Debug(out)
}

func Part1(input []string) int {
    nodes, _ := ParseInput(input)
    start := Coord{0,0}
    end := Coord{len(nodes) - 1, len(nodes[0]) - 1}
    return FindShortestPath(&nodes, start, end)
}
