package main

import (
	"fmt"
)

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

	/*
		ch := make(chan int)
		go func() { // (2.0)
			ch <- 100 // (3.NB)
		}()
		data := <-ch //(1.B, 4.UB)
		fmt.Println(data)
	*/

	ch := make(chan int)
	go func() { // (1.0)
		ch <- 100 // (2.B) (4.UB)
	}()
	data := <-ch // (3.NB)
	fmt.Println(data)

	// modify the above in such a way that the
	// "receive" & "print" happens in a go routine (not in "main")
	/*
		ch := make(chan int)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := <-ch
			time.Sleep(100 * time.Millisecond)
			fmt.Println(data)
		}()
		ch <- 100
		wg.Wait()
	*/
}
