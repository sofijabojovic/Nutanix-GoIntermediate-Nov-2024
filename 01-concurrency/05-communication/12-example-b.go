package main

import (
	"fmt"
)

func main() {
	// Share memory by communicating (using channels)
	/*
		var ch chan int
		make(chan int)
	*/
	ch := make(chan int)
	go add(100, 200, ch)
	// receive the result from the channel
	result := <-ch //(B)(UB)
	fmt.Println("Result :", result)
}

func add(x, y int, ch chan int) {
	result := x + y
	// send the result to the channel
	ch <- result // (NB)
}
