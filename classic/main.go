package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the Cli")

	greeting(os.Stdout, "Joseph")
	greeting2("Janet")
}

func greeting(out io.Writer, name string) {
	fmt.Fprintf(out, "Hello,%s\n", name)
}

func greeting2(name string) {
	fmt.Printf("Hello,%s\n", name)
}
