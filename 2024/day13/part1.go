// part1.go
package day13

import (
    "regexp"
    "strings"
    "strconv"
    "fmt"
    "errors"
)

type Pos struct {
    x int
    y int
}

type Claw struct {
    button rune
    x int
    y int
    cost int
    tries int
}

type ClawMachine struct {
    prize Pos
    claws map[rune]*Claw
}

func ParseClaw(s string) Claw {
    re := regexp.MustCompile(`^Button\s(\w):\sX\+(\d+),\sY\+(\d+)`)
    matches := re.FindStringSubmatch(s)
    x, _ := strconv.Atoi(matches[2])
    y, _ := strconv.Atoi(matches[3])
    button := matches[1][0]
    if button == 'A' {
        return Claw{ button: rune(button), x: x, y: y, cost: 3, tries: 0 }
    }
    return Claw{ button: rune(button), x: x, y: y, cost: 1, tries: 0}
}

func ParseInput(input []string, isPart2 bool) []ClawMachine {
    var cms = []ClawMachine{}
    var cm ClawMachine
    for _, line := range input {
        if line == "" {
            continue
        }
        if strings.HasPrefix(line, "Prize") {
            re := regexp.MustCompile(`^Prize:\sX=(\d+),\sY=(\d+)`)
            matches := re.FindStringSubmatch(line)
            x, _ := strconv.Atoi(matches[1])
            y, _ := strconv.Atoi(matches[2])
            if isPart2 {
                cm.prize = Pos{ x: x+10000000000000, y: y+10000000000000 }
            } else {

                cm.prize = Pos{ x: x, y: y }
            }
            cms = append(cms, cm)
            cm = ClawMachine{}
        } else {
            if cm.claws == nil {
                cm.claws = map[rune]*Claw{}
            }
            c := ParseClaw(line)
            cm.claws[c.button] = &c
        }
    }
    return cms
}

func (cm ClawMachine) Print() {
    machineStr := "Claw Machine\n"
    for _, claw := range cm.claws {
        machineStr += string(claw.button)
        machineStr += fmt.Sprintf(": X+%d, Y+%d cost: %d tries: %d\n", claw.x, claw.y, claw.cost, claw.tries)
    }
    machineStr += fmt.Sprintf("Prize: X=%d, Y=%d\n", cm.prize.x, cm.prize.y)
    machineStr += fmt.Sprintf("Cost: %d\n", cm.ComputeCost())
    fmt.Println(machineStr)
}

func (pos Pos) AddClaw(c Claw) Pos {
    return Pos{ x: pos.x + c.x, y: pos.y + c.y }
}

func (pos Pos) Add(other Pos) Pos {
    return Pos{ x: pos.x + other.x, y: pos.y + other.y }
}

func (pos Pos) Sub(other Pos) Pos {
    return Pos{ x: pos.x - other.x, y: pos.y - other.y }
}

func (claw Claw) HasPath(pos Pos, isPart2 bool) (int, error) {
    if pos.x == 0 || pos.y == 0 {
        return 0, errors.New("No path")
    }
    if (pos.x / claw.x == pos.y / claw.y) && pos.x % claw.x == 0  && pos.y % claw.y == 0 {
        steps := pos.x / claw.x
        if (steps > 100 && !isPart2) || steps < 0 {
            return 0, errors.New("No path")
        }
        return steps, nil
    }
    return 0, errors.New("No path")
}

func (cm ClawMachine) FindSolutions(isPart2 bool) []ClawMachine {
    solutions := []ClawMachine{}
    position := Pos{ x: 0, y: 0 }
    clawA := cm.claws['A']
    clawB := cm.claws['B']
    for {
        position = position.AddClaw(*clawA)
        clawA.tries += 1
        if clawA.tries > 100 && !isPart2 {
            return solutions
        }
        if clawA.x > cm.prize.x || clawA.y > cm.prize.y {
            return solutions
        }

        if position.x > cm.prize.x || position.y > cm.prize.y {
            return solutions
        }
        steps, err := (*clawB).HasPath(cm.prize.Sub(position), isPart2)
        if err == nil {
            copy_cm := cm.Copy()
            copy_cm.claws['B'].tries = steps
            if err := copy_cm.CheckSolution(); err == nil {
                solutions = append(solutions, copy_cm)
            }
        }
        if position == cm.prize {
            solutions = append(solutions, cm)
            return solutions
        }
    }
}

func (cm ClawMachine) Copy() ClawMachine {
    copy_cm := ClawMachine{}
    copy_cm.prize = cm.prize
    copy_cm.claws = map[rune]*Claw{}
    for k, v := range cm.claws {
        cpy := *v
        copy_cm.claws[k] = &cpy
    }
    return copy_cm
}

func (cm ClawMachine) ComputeCost() int {
    totalCost := 0
    for _, claw := range cm.claws {
        totalCost += claw.cost * claw.tries
    }
    return totalCost
}

func (cm ClawMachine) Play(isPart2 bool) (ClawMachine, error) {
    solutions := cm.FindSolutions(isPart2)
    if len(solutions) == 0 {
        return cm, errors.New("No solution")
    }
    // fmt.Println("Solutions: ", len(solutions))
    // for _, solution := range solutions {
        // solution.Print()
    // }
    bestSolution := solutions[0]
    for _, solution := range solutions[1:] {
        cost := solution.ComputeCost()
        if cost < bestSolution.ComputeCost() {
            bestSolution = solution
        }
    }
    return bestSolution, nil
}

func (cm ClawMachine) CheckSolution() error {
    if cm.prize.x == 0 || cm.prize.y == 0 {
        return errors.New("No prize")
    }
    if cm.claws['A'] == nil || cm.claws['B'] == nil {
        return errors.New("Missing claw")
    }
    pos := Pos{ x: 0, y: 0 }
    for _, claw := range cm.claws {
        for i := 0; i < claw.tries; i++ {
            pos = pos.AddClaw(*claw)
            if pos == cm.prize {
                return nil
            }
        }
    }
    return errors.New("Solution is not valid")
}


func Part1(input []string) int {
    isPart2 := false
    cms := ParseInput(input, isPart2)
    cost := 0
    for _, cm := range cms {
        // cm.Print()
        bestSolution, err := cm.Play(isPart2)
        if err != nil {
            // fmt.Println(err)
        } else {
            if err := bestSolution.CheckSolution(); err != nil {
                fmt.Println(err)
                bestSolution.Print()
            }
            // bestSolution.Print()
            cost += bestSolution.ComputeCost()
        }
    }
    return cost
}
