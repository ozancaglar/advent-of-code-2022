package utils

import (
	"bufio"
	"log"
	"os"
)

// streamLines returns a scanner to stream a file
func StreamLines(path string) bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return *bufio.NewScanner(file)
}
