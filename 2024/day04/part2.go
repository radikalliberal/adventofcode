// part2.go
package day04

func checkFrameP2(input []string) int {
    nMaxCol := len(input[0])
    nMaxRows := len(input)
    nFoundXmas := 0
    for row, line := range input {
        for col, char := range line {
            if char == 'A' {
                if col > 0 && col < nMaxCol-1 && row > 0 && row < nMaxRows-1 {
                    surrounding := string([]byte {
                        input[row-1][col-1],
                        input[row-1][col+1],
                        input[row+1][col+1],
                        input[row+1][col-1],
                    })
                    // check if all runes are 'M' or 'S'
                    if surrounding == "MMSS" || surrounding == "SMMS" || surrounding == "SSMM" || surrounding == "MSSM" {
                        nFoundXmas++
                    }
                }
            }
        }
    }
    return nFoundXmas
}

func Part2(input []string) int {

    nFoundXmas := checkFrameP2(input)
    return nFoundXmas
}
