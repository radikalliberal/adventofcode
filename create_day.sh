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

# Create the directory
DIR="day${DAY}"
mkdir -p "$DIR"

# Create the files with the required content
cat <<EOL > "${DIR}/day${DAY}_test.go"
// day${DAY}_test.go
package day${DAY}

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		name    string
		input []string
		expected int
	}{
		{
			name:	 "2 lines 1 number",
			input:    []string{"Example 1", "Example 2"},
			expected: 33,
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
	tests := []struct {
		name    string
		input []string
		expected int
	}{
		{
			name:	 "2 lines 1 number",
			input:    []string{"Example 1", "Example 2"},
			expected: 33,
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

echo -e "${GREEN}Directory and files for day${DAY} created successfully.${NC}"
