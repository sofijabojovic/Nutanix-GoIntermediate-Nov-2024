package main

import (
	"fmt"
)

func main() {

	/*
		ch := make(chan int, 1)
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	ch := make(chan int, 3)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	ch <- 100
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	ch <- 200
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	ch <- 300
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	fmt.Println("<-ch :", <-ch)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	fmt.Println("<-ch :", <-ch)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))
	fmt.Println("<-ch :", <-ch)
	fmt.Printf("len(ch) : %d, cap(ch) : %d\n", len(ch), cap(ch))

}
