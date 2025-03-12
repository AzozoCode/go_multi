package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("home", "somepath")
	dir := filepath.Base("./..")

	fmt.Println("Path", path)

	fmt.Println("dir", dir)
}
