package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	wordCount := 0

	for scanner.Scan() {
		word := strings.Fields(scanner.Text())
		wordCount += len(word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Total word: %d\n", wordCount)

}
