// part2.go
package day02

import (
    "sync"
)

func Map(lines []string, fn func(string) int) []int {
	var wg sync.WaitGroup
	results := make([]int, len(lines))
	
	for i, v := range lines {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			results[i] = fn(v)
		}(v)
	}
	
	wg.Wait() // Wait for all goroutines to finish
	return results
}

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

func compute_line(line string) int {
    nums := ParseLine(line)
    perms := create_permutations(nums)
    for _, perm := range perms {
        if EvalSafe(perm) == 1 {
            return 1
        }
    }
    return 0
}
func SumArray(inputs []int) int {
	sum := 0
	for _, v := range inputs {
		sum += v
	}
	return sum
}

func Part2(input []string) int {
    results := Map(input, compute_line)
    return SumArray(results)
}
