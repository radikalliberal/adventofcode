// part2.go
package day05

import (
    "slices"
)

func Insert(num int, slice []int, index int) []int {
    // insert num at index in slice
    new_slice := append(slice, 0)
    copy(new_slice[index+1:], new_slice[index:])
    new_slice[index] = num
    // new_slice := append(slice[:index], num, slice[index:]...)
    return new_slice
}

func FindGoodOrdering(requirements map[int]Requirement, update []int) []int {
    var goodOrdering []int 
    for _, num := range update {
        req, ok := requirements[num]
        if !ok {
            goodOrdering = append(goodOrdering, num)
        } else {
            isInserted := false
            for k, goNum := range goodOrdering {
                if slices.Contains(req.nums_after, goNum) {
                    goodOrdering = Insert(num, goodOrdering, k)
                    isInserted = true
                    break
                }
            }
            if !isInserted {
                goodOrdering = append(goodOrdering, num)
            }
        }
    }
    return goodOrdering
}

func Part2(input []string) int {
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
        if !is_ok {
            goodOrderung := FindGoodOrdering(requirements, update)
            middleNum := goodOrderung[len(goodOrderung) / 2]
            numUpdatesGood += middleNum
        }
    }
    return numUpdatesGood
}
