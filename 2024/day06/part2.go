// part2.go
package day06

type Coord struct {
    row, col int
}

type Direction struct {
    nd rune
    dx int
    dy int
}

type GameState struct {
    field [][]rune
    field_rows int
    field_cols int
    guard Coord
    numVisited int
    bTerminated bool
    bLooped bool
    numNoNewVisits int
    numSteps int
    nMaxSteps int
}

var dir = map[rune]Direction {
    G_UP: { G_RIGHT, 0, -1 },
    G_RIGHT: { G_DOWN, 1, 0 },
    G_DOWN: { G_LEFT, 0, 1 },
    G_LEFT: { G_UP, -1, 0 },
}

func GetGuardDirection(g *GameState) rune {
    return g.field[g.guard.row][g.guard.col]
}

func GetNextField(g *GameState) rune {
    return g.field[g.guard.row + dir[GetGuardDirection(g)].dy][g.guard.col + dir[GetGuardDirection(g)].dx]
}

func isNextStepInField(g *GameState) bool {
    dx, dy := dir[GetGuardDirection(g)].dx, dir[GetGuardDirection(g)].dy
    nrow := g.guard.row + dy
    ncol := g.guard.col + dx
    return nrow >= 0 && nrow < g.field_rows && ncol >= 0 && ncol < g.field_cols
}

func StepGuard(g *GameState) {
    guardDir := GetGuardDirection(g)
    if isNextStepInField(g) {
        nfield := GetNextField(g)
        if nfield != BLOCKED {
            g.field[g.guard.row][g.guard.col] = VISITED
            if nfield != VISITED {
                g.numVisited++
            } else {
                g.numNoNewVisits++
            }
            g.guard.row += dir[guardDir].dy
            g.guard.col += dir[guardDir].dx
            g.field[g.guard.row][g.guard.col] = guardDir
        } else {
            g.field[g.guard.row][g.guard.col] = dir[guardDir].nd
        }
    } else {
        g.field[g.guard.row][g.guard.col] = VISITED
        g.bTerminated = true
    }
    g.numSteps++
    if g.numNoNewVisits > g.numVisited {
        g.bLooped = true
    }
}

func CopyField(field [][]rune) [][]rune {
    newField := make([][]rune, len(field))
    for i, row := range field {
        newField[i] = make([]rune, len(row))
        copy(newField[i], row)
    }
    return newField
}

func CopyGameState(g GameState) GameState {
    newField := CopyField(g.field)
    newGuard := Coord{g.guard.row, g.guard.col}
    return GameState{
        newField,
        g.field_rows,
        g.field_cols,
        newGuard,
        g.numVisited,
        g.bTerminated,
        g.bLooped,
        0,
        0,
        g.nMaxSteps,
    }
}

func ComputeGame(g *GameState) {
    for {
        StepGuard(g)
        if g.bTerminated || g.bLooped {
            return
        }
    }
}

func GuardPosition(field [][]rune) Coord {
    for i, row := range field {
        for j, r := range row {
            if r == G_UP || r == G_DOWN || r == G_LEFT || r == G_RIGHT {
                return Coord{i, j}
            }
        }
    }
    panic("Guard not found")
}

func ParseGameState(input []string) GameState {
    field := ParseInput(input)
    startPosition := GuardPosition(field)
    return GameState{field, len(field), len(field[0]), startPosition, 0, false, false, 0, 0, 0}
}

func Part2(input []string) int {
    gbase := ParseGameState(input)
    ComputeGame(&gbase)
    g := ParseGameState(input)
    nFoundBlocks := 0
    for i, row := range gbase.field {
        for j, r := range row {
            if r == VISITED {
                if i == g.guard.row && j == g.guard.col {
                    continue
                }
                current_g := CopyGameState(g)
                current_g.field[i][j] = BLOCKED
                ComputeGame(&current_g)
                if current_g.bLooped {
                    nFoundBlocks++
                    current_g.field[i][j] = NEW_BLOCK
                }
            }
        }
    }
    return nFoundBlocks 
}
