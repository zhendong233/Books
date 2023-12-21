package main

import (
	"fmt"
	"sync"
)

func printOddNumber(ch chan int, number int) {
	defer wg.Done()
	for i := 1; i <= number; i++ {
		ch <- i
		if i%2 == 1 {
			fmt.Printf("奇数：%d\n", i)
		}
	}
}

func printEvenNumber(ch chan int, number int) {
	defer wg.Done()
	for i := 1; i <= number; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Printf("偶数：%d\n", i)
		}
	}
}

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	wg.Add(2)
	go printEvenNumber(ch, 10)
	go printOddNumber(ch, 10)
	wg.Wait()
}
