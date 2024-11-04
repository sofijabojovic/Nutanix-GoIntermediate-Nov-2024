package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(1 * time.Second)
			fmt.Println(data)
			continue
		}
		break
	}
}

func fn(ch chan<- int) {
	var count int = rand.Intn(20)
	fmt.Println("[fn] count :", count)
	for i := range count {
		ch <- (i + 1) * 10
	}
	close(ch)

}
