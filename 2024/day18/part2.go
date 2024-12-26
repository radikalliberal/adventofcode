// part2.go
package day18

import (
    "container/heap"
    "fmt"
    "log/slog"
)

type History []Coord
type Map [][]Node

func MazeIsPossible(nodes *Map, start Coord, end Coord) (History, bool) {
    startPath := Path{
        _nodes: nodes,
        history: []Coord{start},
        pos: start,
        cost: (*nodes)[start.row][start.col].cost,
    }
    cm := CreateCostMatrix(len(*nodes), len((*nodes)[0]), 9999999)
    // startPath.Print()
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
            return path.history, true
        }
        paths := path.PossiblePaths()
        for _, p := range paths {
            if cm.Get(p.pos) >= len(p.history) {
                heap.Push(&pq, &Item{ p: &p, priority: -(p.cost * len(p.history))})
                cm.Set(p.pos, len(p.history))
            }
        }
        iteration += 1
     //    if iteration % 100000 == 0 {
	    // path.Print()
     //    }
    }
    return nil, false
}

func (h *History) InHistory(coord Coord) bool {
    for _, c := range *h {
	if c.row == coord.row && c.col == coord.col {
	    return true
	}
    }
    return false
}

func (m *Map) AddCorruption(coord Coord) {
    (*m)[coord.row][coord.col].cost = 9999999
    (*m)[coord.row][coord.col].isCorrupted = true
}

func (h *History) GetPrecedingCoord(coord Coord) Coord {
    for i, c := range *h {
	if c.row == coord.row && c.col == coord.col {
	    return (*h)[i - 1]
	}
    }
    return Coord{-1, -1}
}

func (c Coord) Print() string {
    return fmt.Sprintf("%d,%d", c.col, c.row)
}

func Part2(input []string) string {
    nodes, moreCorruption := ParseInput(input)
    start := Coord{0,0}
    end := Coord{len(nodes) - 1, len(nodes[0]) - 1}
    var lastCorruption Coord
    lastHist, _ := MazeIsPossible(&nodes, start, end)
    for _, corruptNode := range moreCorruption {
	lastCorruption = corruptNode
	nodes.AddCorruption(corruptNode)
	if lastHist.InHistory(corruptNode) {
	    // start = lastHist.GetPrecedingCoord(corruptNode)
	    hist, isPossible := MazeIsPossible(&nodes, start, end)
	    lastHist = hist
	    if !isPossible {
		break
	    }
	}
	slog.Debug(fmt.Sprintf("Corruption at %s is possible", corruptNode.Print()))

    }
    return lastCorruption.Print()
}
