// part1.go
package day15

import "fmt"


const (
    ROBOT = '@'
    BOX = 'O'
    EMPTY = '.'
    WALL = '#'
    UP = '^'
    RIGHT = '>'
    DOWN = 'v'
    LEFT = '<'
)

type Pos struct {
    row int
    col int
}

type Map struct {
    rows int
    cols int
    m [][]rune
    r Pos // Robot Position
}

var dir = map[rune]Pos{ // direction
    '<': { 0, -1 },
    '^': { -1, 0 },
    '>': { 0, 1 },
    'v': { 1, 0 },
}

var rdir = map[rune]Pos{ // reverse direction
    '<': { 0, 1 },
    '^': { 1, 0 },
    '>': { 0, -1 },
    'v': { -1, 0 },
}

func (p *Pos) Move(dir Pos) Pos {
    return Pos{p.row + dir.row, p.col + dir.col}
}

func ParseInput(input []string) (Map, string) {
    // parse map till first blank line
    rows := 0
    for i := 0; i < len(input); i++ {
        if input[i] == "" {
            rows = i
            break
        }
    }
    var m [][]rune
    cols := 0
    var rpos Pos
    for row, line := range input[:rows] {
        cols = len(line)
        parsedLine := make([]rune, cols)
        for col, c := range line {
            parsedLine[col] = c
            if c == ROBOT {
                rpos = Pos{row, col}
            }
        }
        m = append(m, parsedLine)
    }
    movements := ""
    for _, line := range input[rows+1:] {
        movements += line
    }
    return Map{rows, cols, m, rpos}, movements
}

func (m *Map) PushBoxes(endPos, reverseDir Pos, numBoxes int) {
    for range numBoxes {
        m.SetField(endPos, BOX)
        endPos = endPos.Move(reverseDir)
    }
    m.SetField(endPos, ROBOT)
    m.r = endPos
    endPos = endPos.Move(reverseDir)
    m.SetField(endPos, EMPTY)
}

func (m *Map) Print() {
    for _, row := range m.m {
        for _, field := range row {
            // use ansi colors
            switch field {
            case ROBOT:
                fmt.Printf("\033[1;31m%c\033[0m", field)
            case BOX:
                fmt.Printf("\033[1;32m%c\033[0m", field)
            case LEFT_BOX:
                fmt.Printf("\033[1;32m%c\033[0m", field)
            case RIGHT_BOX:
                fmt.Printf("\033[1;32m%c\033[0m", field)
            case WALL:
                fmt.Printf("\033[1;34m%c\033[0m", field)
            default:
                fmt.Printf("%c", field)
            }
        }
        fmt.Println()
    }
    fmt.Println()
}

func (m *Map) SetField(pos Pos, field rune) {
    m.m[pos.row][pos.col] = field
}

func (m *Map) GetField(pos Pos) rune {
    return m.m[pos.row][pos.col]
}

func (m *Map) Update(move rune) {
    current_pos := m.r
    direction := dir[move]
    numBoxes := 0
    for {
        current_pos = current_pos.Move(direction)
        field := m.GetField(current_pos)
        switch field {
        case WALL:
            return
        case EMPTY:
            m.PushBoxes(current_pos, rdir[move], numBoxes)
            return
        case BOX:
            numBoxes++
        }
    }
}

func (m *Map) ScoreBoxes() int {
    score := 0
    for r, row := range m.m {
        for c, field := range row {
            if field == BOX {
                score += r * 100 + c
            }
        }
    }
    return score
}

func Part1(input []string) int {
    m, moves := ParseInput(input)
    // m.Print()
    for i := range moves {
        move := rune(moves[i])
        // fmt.Println("Move: ", string(move))
        m.Update(move)
        // m.Print()
    }
    return m.ScoreBoxes()
}
