package main

import (
	"fmt"
	"sync"
	"time"
)

var ready = false
var mu sync.Mutex

var cond = sync.NewCond(&mu)

func worker(id int) {
	mu.Lock()

	for !ready {
		cond.Wait()
	}

	mu.Unlock()
	fmt.Printf("worker %d is running...", id)

}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	//start multiple workers
	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}

	//simulate some work

	time.Sleep(2 * time.Second)

	//signal workers to start
	mu.Lock()
	ready = true
	cond.Broadcast() //wake up all running goroutines
	mu.Unlock()

	wg.Wait()
	fmt.Println("All workers completed")
}
