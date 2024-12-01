// part2.go
package day01


func compute_similarity(l []int, occ map[int]int) int {
    similarity := 0
    for _, lval := range l {
        similarity += occ[lval] * lval
    }
    return similarity
}


func parseInput_part2(input []string) ([]int, map[int]int) {
    left_list := make([]int, len(input))
    occurences_right := make(map[int]int)
    for i, line := range input {
        left, right := parseLine(line)
        left_list[i] = left
        occurences_right[right]++
    }
    return left_list, occurences_right
}


func Part2(input []string) int {
    left_list, right_occ:= parseInput_part2(input)

    return compute_similarity(left_list, right_occ)
}
