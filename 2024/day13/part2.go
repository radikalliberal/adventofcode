// part2.go
package day13

import (
    "fmt"
    "errors"
    // "math/big"
)
func (cm ClawMachine) FindSolutionsPart2() []ClawMachine {
    solutions := []ClawMachine{}
    position := Pos{ x: 0, y: 0 }
    clawA := cm.claws['A']
    clawB := cm.claws['B']
    for {
        position = position.AddClaw(*clawA)
        clawA.tries += 1
        if clawA.x > cm.prize.x || clawA.y > cm.prize.y {
            return solutions
        }

        if position.x > cm.prize.x || position.y > cm.prize.y {
            return solutions
        }
        steps, err := (*clawB).HasPath(cm.prize.Sub(position), true)
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

func (cm ClawMachine) PlayPart2() (ClawMachine, error) {
    solutions := cm.FindSolutionsPart2()
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

func Part2(input []string) int {
    isPart2 := true
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
