// part1.go
package day12

import (
    "math"
    // "fmt"
    "adventofcode/utils"
)

type Field struct {
    row int
    col int
    regionRune rune
    region int
    merged bool
    neighbors int
}

type Pos struct {
    x int
    y int
}

func ParseInput(input []string) [][]Field {
    cur_region := 0
    regionMap := make([][]Field, len(input))
    for i, line := range input {
        regionMap[i] = make([]Field, len(line))
        for j, r := range line {
            regionMap[i][j] = Field{row: i, col: j, regionRune: r, region: cur_region, merged: false, neighbors: 0}
            cur_region += 1
        }
    }
    return regionMap
}

func (f Field) Neighbors(regionMap *[][]Field) []*Field {
    neighbors := []*Field{}
    coords := []Pos{
        {f.row-1, f.col},
        {f.row+1, f.col},
        {f.row, f.col-1},
        {f.row, f.col+1},
    }
    for _, n := range coords {
        if n.x >= 0 && n.x < len(*regionMap) && n.y >= 0 && n.y < len((*regionMap)[0]) {
            neighbors = append(neighbors, &(*regionMap)[n.x][n.y])
        }
    }
    return neighbors
}

func MergeTwoRegions(regionMap *[][]Field, newRegion int, oldRegion int) {
    for i := 0; i < len(*regionMap); i++ {
        for j := 0; j < len((*regionMap)[0]); j++ {
            if (*regionMap)[i][j].region == oldRegion {
                (*regionMap)[i][j].region = newRegion
            }
        }
    }
}

func JoinRegions(regionMap *[][]Field) {
    for i := 0; i < len(*regionMap); i++ {
        for j := 0; j < len((*regionMap)[0]); j++ {
            curField := &(*regionMap)[i][j]
            for _, n := range curField.Neighbors(regionMap) {
                if n.regionRune == curField.regionRune {
                    curField.neighbors += 1
                    minRegionID := int(math.Min(float64(curField.region), float64(n.region)))
                    maxRegionID := int(math.Max(float64(curField.region), float64(n.region)))
                    curField.region = minRegionID
                    MergeTwoRegions(regionMap, minRegionID, maxRegionID)
                }
            }
        }
    }
}

func ComputeRegionScore(regionMap *[][]Field, region int) int {
    perimeter := 0
    numFields := 0
    // regionRune := '.'
    for i := 0; i < len(*regionMap); i++ {
        for j := 0; j < len((*regionMap)[0]); j++ {
            if (*regionMap)[i][j].region == region {
                // regionRune = (*regionMap)[i][j].regionRune
                numFields += 1
                perimeter += 4 - (*regionMap)[i][j].neighbors
            }
        }
    }
    // fmt.Println("Region: ", region, "RegionRune: ", string(regionRune), "NumFields: ", numFields, "Perimeter: ", perimeter, "Score: ", numFields * perimeter)
    return numFields * perimeter
}

func GetUniqueRegions(regionMap *[][]Field) []int {
    uniqueRegions := []int{}
    for i := 0; i < len(*regionMap); i++ {
        for j := 0; j < len((*regionMap)[0]); j++ {
            region := (*regionMap)[i][j].region
            if !utils.ContainsInt(uniqueRegions, region) {
                uniqueRegions = append(uniqueRegions, region)
            }
        }
    }
    return uniqueRegions
}


func ComputeScore(regionMap *[][]Field) int {
    score := 0
    uniqueRegions := GetUniqueRegions(regionMap)
    for _, region := range uniqueRegions {
        score += ComputeRegionScore(regionMap, region)
    }
    return score
}

func Part1(input []string) int {
    regionMap := ParseInput(input)
    JoinRegions(&regionMap)
    return ComputeScore(&regionMap)
}
