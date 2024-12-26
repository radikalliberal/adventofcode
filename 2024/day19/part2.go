// part2.go
package day19

import (
    "fmt"
)

type PickedPatterns []string

func (pp PickedPatterns) ToString() string {
    out := ""
    for _, p := range pp {
        out += p + "\n"
    }
    return out
}


func Unique(pp *[]PickedPatterns) []PickedPatterns {
    keys := make(map[string]bool)
    list := []PickedPatterns{}
    for _, entry := range *pp {
        key := entry.ToString()
        if _, value := keys[key]; !value {
            keys[key] = true
            list = append(list, entry)
        }
    }
    return list
}

func DiscardSubPatterns(picked []PickedPatterns, target Target) []PickedPatterns {
    out := []PickedPatterns{}
    for _, p := range picked {
        if len(p.ToString()) < len(target.ToString()) {
            out = append(out, p)
        }
    }
    return out
}

func CopyPicks(picks []PickedPatterns) []PickedPatterns {
    out := []PickedPatterns{}
    for _, p := range picks {
        out = append(out, p)
    }
    return out
}

func IsInvalidPick(picked PickedPatterns, target Target) bool {
    concat := ""
    for _, p := range picked {
        concat += p
    }
    return concat != target.ToString()
}

func (p *Patterns) FindAllPermutationsWithCache(target Target, picked PickedPatterns, memo *map[string][]PickedPatterns) []PickedPatterns {
    var cachePicked []PickedPatterns
    var exists bool
    if cachePicked, exists = (*memo)[target.ToString()]; !exists {
        cachePicked = p.FindAllPermutations(target, PickedPatterns{}, memo)
        (*memo)[target.ToString()] = Unique(&cachePicked)
    } else {
        // fmt.Println("Cache hit for target: ", target.ToString(), " with result: ", cachePicked)
        cachePicked = CopyPicks(cachePicked)
    }
    if len(picked) != 0 {
        if len(cachePicked) != 0 {
            for i := range cachePicked {
                cachePicked[i] = append(picked, cachePicked[i]...)
            }
        } else {
            cachePicked = []PickedPatterns{picked}
        }
    }
    // fmt.Println("FindAllPermutationsWithCache(\npattern:", p.ToString(), ",\ntarget:", target.ToString(), ",\npicked:", picked.ToString(), ")\n = ", cachePicked)
    return cachePicked
}

func (p *Patterns) FindAllPermutations(target Target, picked PickedPatterns, memo *map[string][]PickedPatterns) []PickedPatterns {
    if len(target.colors) == 0 {
        return []PickedPatterns{picked}
    }
    // if memo == nil {
    //     memo = &map[string][]PickedPatterns{}
    // }
    allPatterns := []PickedPatterns{}
    for _, cp := range p.patterns {
        matches := cp.FindMatches(&target)
        pattern := cp.ToString()
        for _, idxMatch := range matches {
            target1, target2, err := cp.Split(&target, idxMatch)
            if err != nil {
                continue
            }
            if memo == nil {
                picked1 := p.FindAllPermutations(target1, picked, memo)
                picked2 := p.FindAllPermutations(target2, PickedPatterns{pattern}, memo)
                perms := CombinePermutations(&picked1, &picked2)
                allPatterns = append(allPatterns, perms...)
            } else {
                picked1 := p.FindAllPermutationsWithCache(target1, picked, memo)
                picked2 := p.FindAllPermutationsWithCache(target2, PickedPatterns{pattern}, memo)
                perms := CombinePermutations(&picked1, &picked2)
                allPatterns = append(allPatterns, perms...)
            }
        }
    }
    // fmt.Println("All patterns: ", allPatterns, " for target: ", target.ToString())
    return Unique(&allPatterns)
    // return DiscardSubPatterns(Unique(allPatterns), target)
}

func CombinePermutations(picked1 , picked2 *[]PickedPatterns) []PickedPatterns {
    perms := []PickedPatterns{}
    for _, p1 := range *picked1 {
        for _, p2 := range *picked2 {
            perms = append(perms, append(p1, p2...))
        }
    }
    return perms
}

func Part2(input []string) int {
    patterns, targets := ParseInput(input)
    numAllPermutations := 0
    for _, target := range targets {
        if !patterns.CheckTarget(target) {
            fmt.Println("No match for target: ", target.ToString())
            continue
        }
        memo := map[string][]PickedPatterns{}
        perms:= patterns.FindAllPermutations(target, PickedPatterns{}, &memo)
        perms = Unique(&perms)
        filtered := []PickedPatterns{}
        for _, p := range perms {
            if !IsInvalidPick(p, target) {
                filtered = append(filtered, p)
            }
        }
        fmt.Println("Unique permutations: ", filtered, " for target: ", target.ToString())
        numAllPermutations += len(filtered)
    }
    return numAllPermutations
}
