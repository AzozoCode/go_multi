package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Generate a random number from 1 - 100

	for {
		randomNumber := rand.Intn(100) + 1
		fmt.Printf("Random number:%d\n", randomNumber)
		time.Sleep(1 * time.Second)
	}

}
