package utils

import (
	"bufio"
	"os"
	"fmt"
)

// read input via pipe

func ReadPipe() ([]string, error) {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input, nil
}

func ReadFile(filename string) ([]string, error) {
	file, e := os.Open(filename)
	defer file.Close()
	if file == nil {
		return []string{}, fmt.Errorf("could not open file %s", filename)
	}
	if e != nil {
		return []string{}, e
	}

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input, nil
}
