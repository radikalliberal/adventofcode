// part2.go
package day02

func create_permutations(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{nums}
    }
    permutations := [][]int{}
    for i, _ := range nums {
        rest := make([]int, len(nums)-1)
        copy(rest[:i], nums[:i])
        copy(rest[i:], nums[i+1:])
        permutations = append(permutations, rest)
    }
    return permutations
}

func Part2(input []string) int {
    num_safe := 0
    for _, line := range input {
        nums := ParseLine(line)
        perms := create_permutations(nums)
        for _, perm := range perms {
            if EvalSafe(perm) == 1 {
                num_safe += 1
                break
            }
        }
    }
    return num_safe
}
