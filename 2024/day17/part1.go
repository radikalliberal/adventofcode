// part1.go
package day17

import (
    "math"
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

type Computer struct {
    P []int
    output string
    pc int
    A int
    B int
    C int
}
// Combo operands 0 through 3 represent literal values 0 through 3.
// Combo operand 4 represents the value of register A.
// Combo operand 5 represents the value of register B.
// Combo operand 6 represents the value of register C.
// Combo operand 7 is reserved and will not appear in valid programs.
func (c *Computer) ComboOperand(operand int) (int, error) {
    if operand < 4 {
        return operand, nil
    }
    switch operand {
    case 4:
        return c.A, nil
    case 5:
        return c.B, nil
    case 6:
        return c.C, nil
    }
    return 0, fmt.Errorf("Invalid combo operand %d", operand)
}

// opcode
// The adv instruction (opcode 0) performs division. The numerator is the value in the A register. The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of the division operation is truncated to an integer and then written to the A register.

func (c *Computer) adv(operand int) error {
    newOperand, err := c.ComboOperand(operand)
    if err != nil {
        return err
    }
    c.A = c.A / int(math.Pow(2, float64(newOperand)))
    return nil
}

// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand, then stores the result in register B.

func (c *Computer) bxl(operand int) {
    c.B = c.B ^ operand
}

// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits), then writes that value to the B register.

func (c *Computer) bst(operand int) error {
    newB, err := c.ComboOperand(operand)
    if err != nil {
        return err
    }
    c.B = newB % 8
    return nil
}

// The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero, it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.

func (c *Computer) jnz(operand int) bool {
    // return true if jump
    if c.A != 0 {
        c.pc = operand
        return true
    }
    return false
}

// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C, then stores the result in register B. (For legacy reasons, this instruction reads an operand but ignores it.)

func (c *Computer) bxc(operand int) {
    c.B = c.B ^ c.C
}

// The out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value. (If a program outputs multiple values, they are separated by commas.)

func (c *Computer) out(operand int) error {
    val, err := c.ComboOperand(operand)
    if err != nil {
        return err
    }
    val = val % 8
    if c.output == "" {
        c.output += fmt.Sprintf("%d", val)
    } else {
        c.output += fmt.Sprintf(",%d", val)
    }
    return nil
}

// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the B register. (The numerator is still read from the A register.)

func (c *Computer) bdv(operand int) error {
    newOperand, err := c.ComboOperand(operand)
    if err != nil {
        return err
    }
    c.B = c.A / int(math.Pow(2, float64(newOperand)))
    return nil
}

// The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the C register. (The numerator is still read from the A register.)

func (c *Computer) cdv(operand int) error {
    newOperand, err := c.ComboOperand(operand)
    if err != nil {
        return err
    }
    d := int(math.Pow(2, float64(newOperand)))
    if d == 0 {
        return fmt.Errorf("Division by zero")
    }
    c.C = c.A / d
    return nil
}

// operand

func ParseInput(input []string) *Computer {
    c := Computer{}
    c.P = []int{}
    for line, reg := range []string{"A", "B", "C"} {
        re := regexp.MustCompile(fmt.Sprintf("Register %s: (\\d+)", reg))
        matches := re.FindStringSubmatch(input[line])
        val, _ := strconv.Atoi(matches[1])
        switch reg {
        case "A":
            c.A = val
        case "B":
            c.B = val
        case "C":
            c.C = val
        }
    }
    re := regexp.MustCompile("Program: (.+)")
    matches := re.FindStringSubmatch(input[4])
    for _, val := range strings.Split(matches[1], ",") {
        i, _ := strconv.Atoi(val)
        c.P = append(c.P, i)
    }
    return &c
}

func (c *Computer) Step() error {
    opcode := c.P[c.pc]
    operand := c.P[c.pc + 1]
    switch opcode {
    case 0:
        err := c.adv(operand)
        if err != nil {
            return err
        }
    case 1:
        c.bxl(operand)
    case 2:
        err := c.bst(operand)
        if err != nil {
            return err
        }
    case 3:
        if c.jnz(operand) {
            return nil
        }
    case 4:
        c.bxc(operand)
    case 5:
        err := c.out(operand)
        if err != nil {
            return err
        }
    case 6:
        err := c.bdv(operand)
        if err != nil {
            return err
        }
    case 7:
        error := c.cdv(operand)
        if error != nil {
            return error
        }
    }
    c.pc += 2
    return nil
}

func (c *Computer) Run() {
    for c.pc < len(c.P) {
        err := c.Step()
        if err != nil {
            fmt.Println(err)
            return
        }
    }
}

func Part1(input []string) string {
    c := ParseInput(input)
    c.Run()
    return c.output
}
