// part2.go
package day07

import (
    "strconv"
)

func Atoi(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return n
}

func Concat(a, b int) int {
    return Atoi(strconv.Itoa(a) + strconv.Itoa(b))
}

var concat = Operation{"concat", '&', func(a, b int) int { return Concat(a, b) }}

func Part2(input []string) int {
    equations := ParseInput(input)
    nSum := 0
    for _, e := range equations {
        e, err := FindPossibleEquation(e, []Operation{add, mul, concat})
        if err == nil {
            // PrintEquation(e)
            nSum += e.result
        }
    }
    return nSum
}
