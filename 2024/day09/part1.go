// part1.go
package day09

import (
    "strconv"
)

func ParseInput(input string) []int {
    nums := make([]int, len(input))
    for i, c := range input {
        nums[i], _ = strconv.Atoi(string(c))
    }
    return nums
}

func SumOfDense(denRep []int) int {
    sum := 0
    for _, num := range denRep {
        sum += num
    }
    return sum
}

func ComputeExpandedRep(denRep []int) []int {
    lenExp := SumOfDense(denRep)
    expandedRep := make([]int, lenExp)

    for i := 0; i < lenExp; i++ {
        expandedRep[i] = -1
    }
    current_index := 0
    for idx, num := range denRep {
        if idx % 2 == 0 {
            for i := 0; i < num; i++ {
                expandedRep[current_index] = idx / 2
                current_index++
            }
        } else {
            current_index += num
        }
    }
    return expandedRep
}

func CompactRep(expandedRep []int) []int {
    for i:=len(expandedRep)-1; i>=0; i-- {
        if expandedRep[i] != -1 {
            for k:=0; k<i; k++ {
                if expandedRep[k] == -1 {
                    expandedRep[k] = expandedRep[i]
                    expandedRep[i] = -1
                    break
                }
            }
        }
    }
    return expandedRep
}

func ComputeChecksum(filesystem []int) int {
    checksum := 0
    for i, n := range filesystem {
        if n == -1 {
            continue
        }
        checksum += n * i
    }
    return checksum
}

func Part1(input []string) int {
    denRep := ParseInput(input[0])
    expandedRep := ComputeExpandedRep(denRep)
    compactRep := CompactRep(expandedRep)
    return ComputeChecksum(compactRep)
}
