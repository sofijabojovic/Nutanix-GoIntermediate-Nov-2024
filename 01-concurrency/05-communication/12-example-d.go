package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := add(100, 200)
	/*
		go func() {
			ch <- 10000
		}()
	*/
	result := <-ch
	fmt.Println("Result :", result)
}

// producer
func add(x, y int) <-chan int /* RECEIVE ONLY CHANNEL */ {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
