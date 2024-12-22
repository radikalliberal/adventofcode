// part2.go
package day17

import (
    "fmt"
    // "math"
)

func (c *Computer) Reset() {
    c.pc = 0
    c.A = 0
    c.B = 0
    c.C = 0
    c.output = ""
}

func (c *Computer) OutputValid(target string) bool {
    for i := 0; i < len(c.output); i++ {
        if c.output[i] != target[i] {
            return false
        }
    }
    return true
}

func Contains(slice []int, val int) bool {
    for _, v := range slice {
        if v == val {
            return true
        }
    }
    return false
}

func (c *Computer) Serialize() string {
    return fmt.Sprintf("%d:%d:%d:%d:%s", c.pc, c.A, c.B, c.C, c.output)
}

func (c *Computer) Print() {
    fmt.Println("A:", c.A, "B:", c.B, "C:", c.C, "PC:", c.pc, "Output:", c.output)
}


func (c *Computer) CheckCandidateSub(lastCandidate int, wantedOutput string) int {

    for K := 0; K > -2; K-- {
        for M := 1; M < 100000; M++ {
            if M % 2 == 0 {
                continue
            }
            for N := 0; N < 32; N++ {
                c.Reset()
                candidate := lastCandidate - (1 << N) * M + K
                c.A = candidate
                steps := 0
                for c.pc < len(c.P) {
                    err := c.Step()
                    steps++
                    if err != nil || !c.OutputValid(wantedOutput) {
                        break
                    }
                    if len(c.output) == len(wantedOutput) {
                        // fmt.Println("\033[1;31m", candidate, "\033[0m")
                        // fmt.Println("N:", N, "M:", M, "K:", K, "Candidate:", lastCandidate, "-", (1 << N), "*", M, "+", K)
                        // c.Print()
                        return candidate
                    }
                }
            }
        }
    }
    return 0
}

func (c *Computer) CheckCandidate(bestCandidate string, lastCandidate int, wantedOutput string) int {

    // suggestion = lastCandidate + 2 ** N * M + K
    // for N in range(1, 1000):
    //    for M in range(1, 1000):
    //       for K in range(-1, 1):
    for K := 0; K > -2; K-- {
        for M := 1; M < 1000000; M++ {
            if M % 2 == 0 {
                continue
            }
            for N := 0; N < 50; N++ {
                c.Reset()
                candidate := lastCandidate + (1 << N) * M + K
                c.A = candidate
                steps := 0
                for c.pc < len(c.P) {
                    err := c.Step()
                    steps++
                    if err != nil || !c.OutputValid(wantedOutput) {
                        break
                    }
                    if len(c.output) > len(bestCandidate) {
                        bestCandidate = c.output
                        // fmt.Println("New best candidate:", bestCandidate, "at A =", candidate, "after", steps, "steps")
                        // // ansi colors for A
                        // fmt.Println("\033[1;31m", candidate, "\033[0m")
                        // fmt.Println("N:", N, "M:", M, "K:", K, "Candidate:", lastCandidate, "+", (1 << N), "*", M, "+", K)
                        // c.Print()
                        return candidate
                        // candidate = 
                    }
                }
            }
        }
    }
    return 0
}



func (c *Computer) GuessRegisterA() int {
    // program
    // 0 bst: b = b % 8
    // 1 bxl: b = b xor 2
    // 2 cdv: c = a / b
    // 3 adv: a = a / 3
    // 4 bxl: b = b xor 7
    // 5 bxc: b = b xor c
    // 6 out: o = b % 8
    // 7 jnz: goto 0
    wantedOutput := c.GetProgram()
    bestCandidate := ""
    lastCandidate := 7
    // var badPrograms = Set{}
    // current lowest 190615638981647

    for {
        lastCandidate = c.CheckCandidate(bestCandidate, lastCandidate, wantedOutput)
        if len(c.output) == len(wantedOutput) {
            // fmt.Println("best candidate:", c.output, "at A =", lastCandidate)
            // c.Print()
            return lastCandidate
        }
        if lastCandidate == 0 {
            // fmt.Println("No more candidates")
            return 0
        }
        bestCandidate = c.output
        c.Reset()
    }
}

func (c *Computer) GuessRegisterAFromBest(currentBest int) int {
    // program
    // 0 bst: b = b % 8
    // 1 bxl: b = b xor 2
    // 2 cdv: c = a / b
    // 3 adv: a = a / 3
    // 4 bxl: b = b xor 7
    // 5 bxc: b = b xor c
    // 6 out: o = b % 8
    // 7 jnz: goto 0
    wantedOutput := c.GetProgram()
    // var badPrograms = Set{}
    // current lowest 190615638981647
    prevBest := 0 
    for {
        currentBest = c.CheckCandidateSub(currentBest, wantedOutput)
        if len(c.output) == len(wantedOutput) {
            // fmt.Println("best candidate:", c.output, "at A =", currentBest)
            // c.Print()
        }
        if currentBest == prevBest || currentBest == 0 {
            // fmt.Println("No more candidates")
            return prevBest
        }
        prevBest = currentBest
        c.Reset()
    }
}

func (c *Computer) GetProgram() string {
    program := ""
    for i := 0; i < len(c.P); i++ {
        if i > 0 {
            program += ","
        }
        program += fmt.Sprintf("%d", c.P[i])
    }
    return program
}

func Part2(input []string) int {
    c := ParseInput(input)
    initalGuess := c.GuessRegisterA()
    return c.GuessRegisterAFromBest(initalGuess)
}
