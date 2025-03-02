package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	prompter()
}

func prompter() {
	fmt.Println(">>")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Printf("<- %s\n", line)
		fmt.Println(">>")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

}
