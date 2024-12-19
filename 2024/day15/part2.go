// part2.go
package day15

// import (
//     "fmt"
// )

const (
    LEFT_BOX = '['
    RIGHT_BOX = ']'
)

type Push struct {
    start Pos
    end Pos
}

func (m *Map) Expand() Map {
    mNew := make([][]rune, m.rows)
    newMap := Map{rows: m.rows, cols: m.cols * 2}
    for i := 0; i < m.rows; i++ {
        mNew[i] = make([]rune, m.cols * 2)
        for j := 0; j < m.cols; j++ {
            if m.GetField(Pos{i,j}) == BOX {
                mNew[i][j*2] = LEFT_BOX
                mNew[i][j*2 + 1] = RIGHT_BOX
            } else if m.m[i][j] == ROBOT {
                mNew[i][j*2] = ROBOT
                mNew[i][j*2 + 1] = EMPTY
                newMap.r = Pos{i, j * 2}
            } else {
                mNew[i][j*2] = m.m[i][j]
                mNew[i][j*2+1] = m.m[i][j]
            }
        }
    }
    newMap.m = mNew
    return newMap
}

func ParseInputPart2(input []string) (Map, string) {
    m, r := ParseInput(input)
    return m.Expand(), r
}

func (m *Map) CheckPath(move rune, pos Pos) ([]Pos, bool){
    direction := dir[move]
    field := m.GetField(pos)
    switch field {
    case ROBOT:
        p, ok := m.CheckPath(move, pos.Move(direction))
        if !ok {
            return nil, false
        }
        return append(p, pos), true
    case WALL:
        return nil, false
    case EMPTY:
        return []Pos{pos}, true
    case LEFT_BOX:
        lp, lok := m.CheckPath(move, pos.Move(direction))
        lp = append(lp, pos)
        if !lok {
            return nil, false
        }
        if direction.col == -1 {
            return lp, true
        }
        rpos := Pos{row: pos.row, col: pos.col + 1}
        rp, rok := m.CheckPath(move, rpos.Move(direction))
        rp = append(rp, rpos)
        if !rok {
            return nil, false
        }
        return append(lp, rp...), true
    case RIGHT_BOX:
        rp, rok := m.CheckPath(move, pos.Move(direction))
        rp = append(rp, pos)
        if !rok {
            return nil, false
        }
        if direction.col == 1 {
            return rp, true
        }
        lpos := Pos{ pos.row, pos.col - 1}
        lp, lok := m.CheckPath(move, lpos.Move(direction))
        lp = append(lp, lpos)
        if !lok {
            return nil, false
        }
        return append(lp, rp...), true
    }
    panic("unreachable")
}

func Contains(p []Pos, pos Pos) bool {
    for _, e := range p {
        if e == pos {
            return true
        }
    }
    return false
}

func FilterPaths(paths []Pos, direction Pos) []Pos {
    // remove duplicates
    filtered := make([]Pos, 0, len(paths))
    for _, p := range paths {
        if !Contains(filtered, p) {
            filtered = append(filtered, p)
        }
    }
    return filtered
}

func (m *Map) ExtraxtBoxes(paths []Pos) []Pos {
    boxes := make([]Pos, 0)
    for _, pos := range paths {
        if m.GetField(pos) == LEFT_BOX {
            boxes = append(boxes, pos)
        }
    }
    return boxes
}


func (m *Map) UpdateMap(paths []Pos, direction Pos) {
    paths = FilterPaths(paths, direction)
    boxes := m.ExtraxtBoxes(paths)
    right := Pos{0, 1}
    for _, box := range boxes {
        m.SetField(box, EMPTY)
        rBox := box.Move(right)
        m.SetField(rBox, EMPTY)
    }
    for _, box := range boxes {
        m.SetField(box.Move(direction), LEFT_BOX)
        rBox := box.Move(right)
        m.SetField(rBox.Move(direction), RIGHT_BOX)
    }
}


func (m *Map) CountBoxes() int {
    score    := 0
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            if m.GetField(Pos{i,j}) == LEFT_BOX {
                score += 100 * i + j
            }
        }
    }
    return score
}

func (m *Map) CheckIntegrity() bool {
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            if m.GetField(Pos{i,j}) == LEFT_BOX {
                if m.GetField(Pos{i, j+1}) != RIGHT_BOX {
                    return false
                }
            }
        }
    }
    return true
}

func Part2(input []string) int {
    m, moves := ParseInputPart2(input)
    m.Print()
    for _, move := range moves {
        // fmt.Printf("Move: %s\n", string(move))
        paths, ok := m.CheckPath(move, m.r)
        if !ok {
            // fmt.Println("Invalid move")
            continue
        }
        m.UpdateMap(paths, dir[move])
        m.SetField(m.r, EMPTY)
        m.r = m.r.Move(dir[move])
        m.SetField(m.r, ROBOT)
        if !m.CheckIntegrity() {
            m.Print()
            panic("Integrity check failed")
        }
        // m.Print()
    }

    m.Print()
    return m.CountBoxes()
}
