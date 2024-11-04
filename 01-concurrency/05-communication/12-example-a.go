package main

import (
	"fmt"
	"sync"
)

func main() {
	// Share memory by communicating (using channels)
	/*
		var ch chan int
		make(chan int)
	*/
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, ch)
	// receive the result from the channel
	result := <-ch //(B)(UB)
	wg.Wait()
	fmt.Println("Result :", result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	// send the result to the channel
	ch <- result // (NB)
	wg.Done()
}
