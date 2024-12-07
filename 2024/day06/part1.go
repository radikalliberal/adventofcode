// part1.go
package day06

const (
    OPEN = '.'
    BLOCKED = '#'
    G_UP = '^'
    G_DOWN = 'v'
    G_LEFT = '<'
    G_RIGHT = '>'
    VISITED = 'X'
    NEW_BLOCK = 'B'
)

func ParseInput(input []string) [][]rune {
    nRows := len(input)
    nCols := len(input[0])
    field := make([][]rune, nRows)
    for i, row := range input {
        field[i] = make([]rune, nCols)
        for j, r := range row {
            field[i][j] = r
        }
    }
    return field
}

func PrintField(field [][]rune) {
    for _, row := range field {
        for _, r := range row {
            print(string(r))
        }
        println()
    }
}

func Step(field [][]rune, bGuardLeft *bool) [][]rune {
    for i, row := range field {
        for j, r := range row {
            switch r {
            case G_UP:
                field[i][j] = VISITED
                if i > 0 {
                    if field[i-1][j] != BLOCKED {
                        field[i-1][j] = G_UP
                    } else {
                        field[i][j] = G_RIGHT
                    }
                } else {
                    *bGuardLeft = true
                }
            case G_DOWN:
                field[i][j] = VISITED
                if i < len(field)-1 {
                    if field[i+1][j] != BLOCKED { 
                        field[i+1][j] = G_DOWN
                    } else {
                        field[i][j] = G_LEFT
                    }
                } else {
                    *bGuardLeft = true
                }
            case G_LEFT:
                field[i][j] = VISITED
                if j > 0 {
                    if field[i][j-1] != BLOCKED {
                        field[i][j-1] = G_LEFT
                    } else {
                        field[i][j] = G_UP
                    }
                } else {
                    *bGuardLeft = true
                }
            case G_RIGHT:
                field[i][j] = VISITED
                if j+1 < len(row) {
                    if field[i][j+1] != BLOCKED {
                        field[i][j+1] = G_RIGHT
                    } else {
                        field[i][j] = G_DOWN
                    }
                } else {
                    *bGuardLeft = true
                }
            }
        }
    }
    return field
}

func Equal(field1, field2 [][]rune) bool {
    for i, row := range field1 {
        for j, r := range row {
            if r != field2[i][j] {
                return false
            }
        }
    }
    return true
}

func ComputeVisited(field [][]rune) int {
    visited := 0
    for _, row := range field {
        for _, r := range row {
            if r == VISITED {
                visited++
            }
        }
    }
    return visited
}

func HasGuard(field [][]rune) bool {
    for _, row := range field {
        for _, r := range row {
            if r == G_UP || r == G_DOWN || r == G_LEFT || r == G_RIGHT {
                return true
            }
        }
    }
    return false
}


func Part1(input []string) int {
    field := ParseInput(input)
    for bGuardLeft := false; !bGuardLeft; {
        Step(field, &bGuardLeft)
        if bGuardLeft {
            break
        }
    }
    return ComputeVisited(field)
}
