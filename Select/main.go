package main

import "fmt"

func main() {
	go func(i int) {
		fmt.Println("Background process running....")
	}(1)

	select {} //blocks forever
}
