// day03_test.go
package day03

import (
    "testing"
    "adventofcode/utils"
)

func TestPart1(t *testing.T) {
    ex1, _ := utils.ReadFile("test1.txt")
    tests := []struct {
        name    string
        input []string
        expected int
    }{
        {
            name:     "part 1 example",
            input:     ex1,
            expected: 161,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Part1(tt.input)
            if result != tt.expected {
                t.Errorf("Expected %d, got %d", tt.expected, result)
            }
        })
    }
}

func TestPart2(t *testing.T) {
    ex2, _ := utils.ReadFile("test2.txt")
    inp2, _ := utils.ReadFile("input2.txt")
    tests := []struct {
        name    string
        input []string
        expected int
    }{
        {
            name:     "part 2 example",
            input:    ex2,
            expected: 48,
        },
        {
            name:     "dont",
            input:    []string{ "don't()'{+;when() ]mul(804,327)?@when()how()+from()mul(731,16)}when(632,481)^ mul(841,888)~:?{[<$+don't()'<&why()mul(668,253)mul(955,947) ' ^(*mul(227,561))^;*{ ,(mul(164,251)# (from()from(31,113)+]when(){mul(45,21)select()mul(180,486);: ?}from()/mul(411,320) ~{-when()){/+-mul(142,258)mul(895,918)#;#:mul(436,950)-{*what()where()select():mul(512,382)~#from(563,273)mul(180,971)(,mul(914,983)#who(){)where(773,354)mul(86,529)select()?$!:-mul(469,419)what()-!+why()#,&]mul(531,80)what()where()!where()" },
            expected: 0,
        },
        {
            name:     "part 2 input",
            input:    inp2,
            expected: 80570939,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Part2(tt.input)
            if result != tt.expected {
                t.Errorf("Expected %d, got %d", tt.expected, result)
            }
        })
    }
}

func BenchmarkPart1(b *testing.B) {
    input, _ := utils.ReadFile("input1.txt")
    for n := 0; n < b.N; n++ {
        Part1(input)
    }
}

func BenchmarkPart2(b *testing.B) {
    input, _ := utils.ReadFile("input2.txt")
    for n := 0; n < b.N; n++ {
        Part2(input)
    }
}
