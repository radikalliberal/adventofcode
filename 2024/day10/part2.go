// part2.go
package day10

func Part2(input []string) int {
    trailMap := ParseInput(input)
    trailheads := computeTrailheads(trailMap)
    numPaths := 0
    for _, trailhead := range trailheads {

        reachableEnds := computePaths(trailMap, trailhead)
        numPaths += len(reachableEnds)
    }
    return numPaths
}
