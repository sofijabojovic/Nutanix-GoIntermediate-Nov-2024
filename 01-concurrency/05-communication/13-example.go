package main

import "fmt"

func main() {
	/*
		ch := make(chan int)
		data := <-ch
		ch <- 100
		fmt.Println(data)
	*/

	/*
		ch := make(chan int)
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	ch := make(chan int)
	go func() { // (2.0)
		ch <- 100 // (3.NB)
	}()
	data := <-ch //(1.B, 4.UB)
	fmt.Println(data)

	// modify the above in such a way that the
	// "receive" & "print" happens in a go routine (not in "main")
}
