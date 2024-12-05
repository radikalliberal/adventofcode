// part1.go
package day05

import (
    "regexp"
    "strconv"
    "slices"
)

type Requirement struct {
    nums_before []int
    nums_after []int
}

func ParseRequirement(line string, req *map[int]Requirement) {
    re := regexp.MustCompile(`\d+`)
    matches := re.FindAllString(line, -1)
    if len(matches) != 2 {
        panic("Invalid input")
    }
    num1, err1 := strconv.Atoi(matches[0])
    if err1 != nil {
        panic(err1)
    }
    num2, err := strconv.Atoi(matches[1])
    if err != nil {
        panic(err)
    }
    if r, ok := (*req)[num1]; ok {
        (*req)[num1] = Requirement{r.nums_before, append(r.nums_after, num2)}
    } else {
        (*req)[num1] = Requirement{[]int{}, []int{num2}}
    }
    if r, ok := (*req)[num2]; ok {
        (*req)[num2] = Requirement{ append(r.nums_before, num1), r.nums_after}
    } else {
        (*req)[num2] = Requirement{[]int{num1}, []int{}}
    }
}

func ParseUpdates(line string, updates *[]int) {
    re := regexp.MustCompile(`\d+`)
    matches := re.FindAllString(line, -1)
    new_update := make([]int, len(matches))
    for i, match := range matches {
        num, err := strconv.Atoi(match)
        if err != nil {
            panic(err)
        }
        new_update[i] = num
    }
    *updates = new_update
}

func CheckRequirements(i int, requirements map[int]Requirement, update []int) bool {
    num := update[i]
    req, ok := requirements[num]
    if !ok {
        return true 
    }
    for _, before := range req.nums_before {
        if !slices.Contains(update, before) {
            continue
        }
        if !slices.Contains(update[:i], before) {
            return false
        }
    }
    for _, after := range req.nums_after {
        if !slices.Contains(update, after) {
            continue
        }
        if !slices.Contains(update[i+1:], after) {
            return false
        }
    }
    return true
}

func ParseInput(input []string) (map[int]Requirement, [][]int) {
    requirements := make(map[int]Requirement)
    updates := make([][]int, 0)
    bParseRequirements := true
    nSizeRequirements := 0
    for lineIdx, line := range input {
        if line == "" {
            nSizeRequirements = lineIdx
            updates = make([][]int, len(input)-nSizeRequirements-1)
            bParseRequirements = false
            continue
        }
        if bParseRequirements {
            ParseRequirement(line, &requirements)
        } else {
            ParseUpdates(line, &updates[lineIdx-nSizeRequirements-1])
        }
    }
    return requirements, updates
}

func Part1(input []string) int {
    requirements, updates := ParseInput(input)
    numUpdatesGood := 0
    for _, update := range updates {
        is_ok := true
        for i := range update {
            is_ok = CheckRequirements(i, requirements, update) && is_ok
            if !is_ok {
                break
            }
        }
        if is_ok {
            numUpdatesGood += update[len(update) / 2]
        }
    }
    return numUpdatesGood
}
