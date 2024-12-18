package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for more := scanner.Scan(); more; more = scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
