// part1.go
package day07

import (
    "strings"
    "strconv"
    "regexp"
    "errors"
)

type Operation struct {
    name string
    sign rune
    op func(int, int) int
}

func (o Operation) Apply(a, b int) int {
    return o.op(a, b)
}

var mul = Operation{"mul", '*', func(a, b int) int { return a * b }}
var add = Operation{"add", '+', func(a, b int) int { return a + b }}


type Equation struct {
    result int
    nums []int
    ops []Operation
}

func PrintEquation(e Equation) {
    if len(e.nums) != len(e.ops) + 1 {
        panic("Invalid equation cant be printed")
    }
    print("Equation: ", e.result, " = ")
    for i, n := range e.nums {
        print(" ", n, " ")
        if i < len(e.ops) {
            print(string(e.ops[i].sign))
        }
    }
    println()
}

func ParseInput(input []string) []Equation {
    equations := make([]Equation, len(input))
    for i, line := range input {
        parts := strings.Split(line, ":")
        if len(parts) != 2 {
            panic("Invalid input")
        }
        result, err := strconv.Atoi(parts[0])
        if err != nil {
            panic(err)
        }
        matches := regexp.MustCompile(`(\d+)`).FindAllString(parts[1], -1)
        nums := make([]int, len(matches))
        for idx, m := range matches {
            n, err := strconv.Atoi(m)
            if err != nil {
                panic(err)
            }
            nums[idx] = n
        }
        equations[i] = Equation{result, nums, []Operation{}}
    }
    return equations
}

func CopyEquation(e Equation) Equation {
    newNums := make([]int, len(e.nums))
    copy(newNums, e.nums)
    newOps := make([]Operation, len(e.ops))
    copy(newOps, e.ops)
    return Equation{e.result, newNums, newOps}
}

func FindPossibleEquation(e Equation, ops []Operation) (Equation, error) {
    if len(e.nums) == 1 {
        return e, nil
    }
    currentValue := e.nums[0]
    if len(e.ops) > 0 {
        for i, op := range e.ops {
            currentValue = op.Apply(currentValue, e.nums[i+1])
        }
    } else {
        currentValue = 0
    }

    if currentValue == e.result {
        if len(e.ops) == len(e.nums) - 1 {
            return e, nil
        }
    }
    if currentValue > e.result {
        return e, errors.New("No equation found, result too high")
    }
    if len(e.ops) >= len(e.nums) - 1 {
        return e, errors.New("No equation found, no more operations")
    }
    for _, op := range ops {
        newEquation := CopyEquation(e)
        newEquation.ops = append(newEquation.ops, op)
        newEquation, err := FindPossibleEquation(newEquation, ops)
        if err == nil {
            return newEquation, nil
        }
    }
    return e, errors.New("No possible equation found")
}

func Part1(input []string) int {
    equations := ParseInput(input)
    nSum := 0
    for _, e := range equations {
        e, err := FindPossibleEquation(e, []Operation{add, mul})
        if err == nil {
            // PrintEquation(e)
            nSum += e.result
        }
    }
    return nSum
}
