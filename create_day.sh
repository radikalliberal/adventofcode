#!/bin/bash

# Define color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if the argument is provided
if [ -z "$1" ]; then
echo 
  echo -e "${RED}Usage: $0 <day_number>${NC}"
  exit 1
fi

# Pad the day number with a leading zero if necessary
DAY=$(printf "%02d" "$1")
PREV_DAY=$(printf "%02d" "$(($1-1))")

# Create the directory
DIR="day${DAY}"
mkdir -p "$DIR"

# Create the files with the required content
cat <<EOL > "${DIR}/day${DAY}_test.go"
// day${DAY}_test.go
package day${DAY}

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
            expected: 1,
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
    tests := []struct {
        name    string
        input []string
        expected int
    }{
        {
            name:     "part 2 example",
            input:    ex2,
            expected: 1,
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
EOL
cat <<EOL > "${DIR}/part1.go"
// part1.go
package day${DAY}

func Part1(input []string) int {
    // Your code for part 1 here
    return 0
}
EOL

cat <<EOL > "${DIR}/part2.go"
// part2.go
package day${DAY}

func Part2(input []string) int {
    // Your code for part 2 here
    return 0
}
EOL

# Create the empty files
touch "${DIR}/input1.txt"
touch "${DIR}/input2.txt"
touch "${DIR}/test1.txt"
touch "${DIR}/test2.txt"

# Add the import statement
IMPORT_STATEMENT="\"adventofcode/2024/day$DAY\""

sed -i "/^)$/i \    $IMPORT_STATEMENT" solution.go
# Add the lines to the Solution() function
NEW_LINES=$(cat <<EOL
    case $DAY:\\
        fmt.Println("  Day $DAY")\\
        fmt.Println("    Part 1: ", day$DAY.Part1(read_input($1, 1)))\\
        fmt.Println("    Part 2: ", day$DAY.Part2(read_input($1, 2)))
EOL
)

# Insert the new lines before the closing brace of the Solution() function
line_number=$(grep -n "fmt.Println(\"    Part 2: \", day$PREV_DAY.Part2" solution.go | cut -d: -f1)
sed -i "${line_number}a\\$NEW_LINES" solution.go
# sed -i "/fmt.Println\(\"\ \ \ \ Part\ 2:\ \",\ day$PREV_DAY./a\ \    $NEW_LINES" solution.go


echo -e "${GREEN}Directory and files for day${DAY} created successfully.${NC}"
