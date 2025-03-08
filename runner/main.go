package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	input := "ls"

	args := []string{}

	ls := exec.Command(input, args...)

	output, err := ls.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
