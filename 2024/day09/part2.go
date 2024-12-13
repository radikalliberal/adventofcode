// part2.go
package day09

import (
     "fmt"
)


func PrintFiles2(filesystem []int) {
    printString := ""
    for _, f := range filesystem {
        if f == -1 {
            printString += "."
        } else {
            printString += fmt.Sprintf("%d", f)
        }
    }
    fmt.Println(printString)
}

func CompactFiles(filesystem []int) []int {
    // PrintFiles2(filesystem)
    currentSize := 1
    for i:=len(filesystem)-1; i>=0; i-- {
        fid := filesystem[i]
        if fid == -1 {
            continue
        }
        if i > 0 && filesystem[i-1] == fid {
            currentSize++
            continue
        }
        for j:=0; j<i; j++ {
            if filesystem[j] == -1 {
                if SumOfDense(filesystem[j:j+currentSize]) == -currentSize {
                    for k:=0; k<currentSize; k++ {
                        filesystem[j+k] = fid
                        filesystem[i+k] = -1
                    }
                    currentSize = 1
                    // PrintFiles2(filesystem)
                    break
                }
            }
        }
        currentSize = 1
    }
    return filesystem
}

func Part2(input []string) int {
    filesystem := ParseInput(input[0])
    expandedRep := ComputeExpandedRep(filesystem)
    files := CompactFiles(expandedRep)
    // PrintFiles(files)
    checksum := ComputeChecksum(files)
    return checksum
}
