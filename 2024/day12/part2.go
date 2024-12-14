// part2.go
package day12

import (
    "fmt"
)

func (f Field) UpNeighbor(regionMap *[][]Field) *Field {
    if f.row-1 >= 0 && f.region == (*regionMap)[f.row-1][f.col].region {
        return &(*regionMap)[f.row-1][f.col]
    }
    return nil
}

func (f Field) DownNeighbor(regionMap *[][]Field) *Field {
    if f.row+1 < len(*regionMap) && f.region == (*regionMap)[f.row+1][f.col].region {
        return &(*regionMap)[f.row+1][f.col]
    }
    return nil
}

func (f Field) LeftNeighbor(regionMap *[][]Field) *Field {
    if f.col-1 >= 0 && f.region == (*regionMap)[f.row][f.col-1].region {
        return &(*regionMap)[f.row][f.col-1]
    }
    return nil
}

func (f Field) RightNeighbor(regionMap *[][]Field) *Field {
    if f.col+1 < len((*regionMap)[0]) && f.region == (*regionMap)[f.row][f.col+1].region {
        return &(*regionMap)[f.row][f.col+1]
    }
    return nil
}


func ComputePerimeterForRegion(regionMap *[][]Field, region int) int {
    // comute horizontal perrowimeter
    numPerimeterRow := 0
    for col := 0; col < len((*regionMap)[0]); col++ {
        runningPerimeterLeft := false
        runningPerimeterRight := false
        for row := 0; row < len(*regionMap); row++ {
            curField := (*regionMap)[row][col]
            if curField.region == region {
                if !runningPerimeterLeft && curField.LeftNeighbor(regionMap) == nil {
                    numPerimeterRow += 1
                    runningPerimeterLeft = true
                }
                if runningPerimeterLeft && curField.LeftNeighbor(regionMap) != nil {
                    runningPerimeterLeft = false
                }
                if !runningPerimeterRight && curField.RightNeighbor(regionMap) == nil {
                    numPerimeterRow += 1
                    runningPerimeterRight = true
                }
                if runningPerimeterRight && curField.RightNeighbor(regionMap) != nil {
                    runningPerimeterRight = false
                }
            } else {
                runningPerimeterLeft = false
                runningPerimeterRight = false
            }

        }
    }
    numPerimeterCol := 0
    for row := 0; row < len((*regionMap)); row++ {
        runningPerimeterUp := false
        runningPerimeterDown := false
        for col := 0; col < len((*regionMap)[0]); col++ {
            if (*regionMap)[row][col].region == region {
                if !runningPerimeterUp && (*regionMap)[row][col].UpNeighbor(regionMap) == nil {
                    numPerimeterCol += 1
                    runningPerimeterUp = true
                }
                if runningPerimeterUp && (*regionMap)[row][col].UpNeighbor(regionMap) != nil {
                    runningPerimeterUp = false
                }
                if !runningPerimeterDown && (*regionMap)[row][col].DownNeighbor(regionMap) == nil {
                    numPerimeterCol += 1
                    runningPerimeterDown = true
                }
                if runningPerimeterDown && (*regionMap)[row][col].DownNeighbor(regionMap) != nil {
                    runningPerimeterDown = false
                }
            } else {
                runningPerimeterUp = false
                runningPerimeterDown = false
            }
        }
    }
    fmt.Printf("Perimeter for region %c: (row %d + col %d) * count %d = score %d\n", GetRuneForRegion(regionMap, region), numPerimeterRow, numPerimeterCol, CountFiledsForRegion(regionMap, region), (numPerimeterRow + numPerimeterCol) * CountFiledsForRegion(regionMap, region))
    return numPerimeterRow + numPerimeterCol
}

func CountFiledsForRegion(regionMap *[][]Field, region int) int {
    count := 0
    for _, row := range *regionMap {
        for _, field := range row {
            if field.region == region {
                count += 1
            }
        }
    }
    return count
}

func GetRuneForRegion(regionMap *[][]Field, region int) rune {
    for _, row := range *regionMap {
        for _, field := range row {
            if field.region == region {
                return field.regionRune
            }
        }
    }
    return ' '
}


func ComputeScorePart2(regionMap *[][]Field) int {
    score := 0
    uniqueRegions := GetUniqueRegions(regionMap)
    for _, region := range uniqueRegions {
        peri := ComputePerimeterForRegion(regionMap, region)
        count := CountFiledsForRegion(regionMap, region)
        score += peri * count
        // fmt.Printf("Region %d: %c, count: %d, peri: %d, score: %d\n", region, GetRuneForRegion(regionMap, region), count, peri, peri*count)
    }
    return score
}

func Part2(input []string) int {
    regionMap := ParseInput(input)
    JoinRegions(&regionMap)
    return ComputeScorePart2(&regionMap)
}
