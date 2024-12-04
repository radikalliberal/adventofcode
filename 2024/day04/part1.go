// part1.go
package day04

func Rotate90(input []string) []string {
    nMaxCol := len(input[0])
    nMaxRows := len(input)
    output := make([]string, nMaxCol)
    for j := nMaxRows - 1; j >= 0; j-- {
        for i := 0; i < nMaxCol; i++ {
            output[i] += string(input[j][i])
        }
    }
    return output
}

func checkFrame(input []string) int {
    nMaxCol := len(input[0])
    nMaxRows := len(input)
    nFoundXmas := 0
    for row, line := range input {
        for col, char := range line {
            if char == 'X' {
                if col < nMaxCol-3 && line[col:col+4] == "XMAS" {
                    nFoundXmas++
                }
                if row < nMaxRows-3 && col < nMaxCol-3 {
                    if input[row+1][col+1] == 'M' && input[row+2][col+2] == 'A' && input[row+3][col+3] == 'S' {
                        nFoundXmas++
                    }
                }
            }
        }
    }
    return nFoundXmas
}

func Part1(input []string) int {

    nFoundXmas :=  0
    nFoundXmas += checkFrame(input)
    for i := 0; i < 3; i++ {
        input = Rotate90(input)
        nFoundXmas += checkFrame(input)
    }
    return nFoundXmas
}
