package adventofcode

import (
	"bufio"
	"os"
)

// read input via pipe

func ReadInput() []string {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
