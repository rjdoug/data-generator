package io

import (
	"bufio"
	"fmt"
	"os"
)

func WriteFile(file string, lines []string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("")
	}
	defer f.Close()

	for _, line := range lines {
		_, err := f.WriteString(line)

		if err != nil {
			return fmt.Errorf("writing to file: %v", err)
		}
	}
	fmt.Printf("%s written successfully\n", file)

	return nil
}

func ReadTextFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading file: %v", err)
	}

	return lines, nil
}
