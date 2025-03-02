package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			wg.Done()
			ch <- i
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Result ", val)
	}

	fmt.Println("Done...")
}
