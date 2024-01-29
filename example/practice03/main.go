package main

import (
	"fmt"
	"sync"
)

// goroutine池
var (
	wg sync.WaitGroup
)

func worker(id int, job, result chan int) {
	defer wg.Done()
	for v := range job {
		fmt.Printf("jobID=%d, v=%d\n", id, v)
		result <- v * 2
	}
}

func main() {
	job := make(chan int, 100)
	result := make(chan int, 100)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, job, result)
	}
	for i := 0; i < 10; i++ {
		job <- i
	}
	close(job)
	for i := 0; i < 10; i++ {
		fmt.Println(<-result)
	}
	wg.Wait()
}
