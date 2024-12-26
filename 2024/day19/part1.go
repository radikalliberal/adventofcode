// part1.go
package day19

import (
    "adventofcode/utils"
    "regexp"
    "errors"
    "fmt"
)


const (
    green = 'g'
    white = 'w'
    red = 'r'
    blue = 'b'
)

type ColorPattern struct {
    colors []rune
    length int
}

func (cp *ColorPattern) ToString() string {
    return string(cp.colors)
}

type Patterns struct {
    patterns []ColorPattern
}

func (p *Patterns) ToString() string {
    s := ""
    for _, cp := range p.patterns {
        s += cp.ToString() + "\n"
    }
    return s
}

type Target struct {
    colors []rune
}

func (t *Target) ToString() string {
    if len(t.colors) == 0 {
        return ""
    }
    return string(t.colors)
}

type Targets []Target

func (t *Targets) ToString() string {
    s := ""
    for _, target := range *t {
        s += target.ToString() + "\n"
    }
    return s
}

func ParseInput(input []string) (Patterns, Targets) {
    patterns := Patterns{}
    targets := []Target{}
    patternRegex := regexp.MustCompile(`(\w+)`)
    for _, match := range patternRegex.FindAllStringSubmatch(input[0], -1) {
        colors := []rune(match[1])
        patterns.patterns = append(patterns.patterns, ColorPattern{colors, len(colors)})
    }
    targetRegex := regexp.MustCompile(`(\w+)`)
    for _, line := range input[1:] {
        for _, match := range targetRegex.FindAllStringSubmatch(line, -1) {
            targets = append(targets, Target{[]rune(match[1])})
        }
    }
    return patterns, targets
}

func (cp *ColorPattern) FindMatches(target *Target) []int {
    matches := []int{}
    for i := 0; i <= len(target.colors) - cp.length; i++ {
        if utils.RuneArrayEquals(target.colors[i:i+cp.length], cp.colors) {
            matches = append(matches, i)
        }
    }
    return matches
}

func (cp *ColorPattern) Split(target *Target, i int) (Target, Target, error) {
    if utils.RuneArrayEquals(target.colors[i:i+cp.length], cp.colors) {
        return Target{target.colors[:i]}, Target{target.colors[i+cp.length:]}, nil
    }
    return Target{}, Target{}, errors.New("No match")
}

var memo map[string]bool

func (p *Patterns) CheckTarget(target Target) bool {
    if len(target.colors) == 0 {
        return true
    }
    // fmt.Println("Checking ", target.ToString())
    for _, cp := range p.patterns {
        matches := cp.FindMatches(&target)
        for _, idxMatch := range matches {
            target1, target2, err := cp.Split(&target, idxMatch)
            // fmt.Println("Splitting ", target.ToString(), " at ", idxMatch, " with ", cp.ToString(), " into ", target1.ToString(), " and ", target2.ToString())
            if err == nil {
                if len(target1.colors) == 0  && len(target2.colors) == 0 {
                    return true
                }
                works1 := false
                works2 := false
                exists := false
                if works1, exists = memo[target1.ToString()]; !exists {
                    works1 = p.CheckTarget(target1)
                    memo[target1.ToString()] = works1
                }
                if works2, exists = memo[target2.ToString()]; !exists {
                    works2 = p.CheckTarget(target2)
                    memo[target2.ToString()] = works2
                }
                if works1 && works2 {
                    return true
                }
            }
        }
    }
    return false
}

func Part1(input []string) int {
    patterns, targets := ParseInput(input)
    validPatterns := 0
    for _, target := range targets {
        memo = make(map[string]bool)
        works := patterns.CheckTarget(target)
        if works {
            validPatterns++
            fmt.Println("Valid pattern: ", target.ToString())
        } else {
            fmt.Println("Invalid pattern: ", target.ToString())
        }

    }
    return validPatterns
}
