package helper

import (
	"bufio"
	"strings"
)

func Read(scanner *bufio.Scanner) int {
	var count int

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		count += len(words)
	}
	return count
}
