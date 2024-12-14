// part2.go
package day11

var memo = make(map[[2]int]int)

func (s Stone) UpdateRecurse(depth int) int {
    numStones := 0
    val, ok :=memo[[2]int{s.val, depth}]
    if ok {
        return val
    }
    if depth == 0 {
        return 1
    }
    // If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
    if s.val == 0 {
        s.val = 1
        res := s.UpdateRecurse(depth-1)
        memo[[2]int{s.val, depth-1}] = res
        numStones += res
    } else if s.Digits() % 2 == 0 {
    // If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
        for _, stone := range s.Split() {
            res := stone.UpdateRecurse(depth-1)
            memo[[2]int{stone.val, depth-1}] = res
            numStones += res
        }
    } else {
    // If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.<F11>
        s.val *= 2024
        res := s.UpdateRecurse(depth-1)
        memo[[2]int{s.val, depth-1}] = res
        numStones += res
    }
    return numStones
}

func Part2(input []string) int {
    stones := ParseInput(input[0])
    // PrintStones(stones)
    numStones := 0
    totalSteps := 75
    for _, stone := range stones {
        numStones += stone.UpdateRecurse( totalSteps)
    }
    return numStones
}
