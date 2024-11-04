package main

import "fmt"

var result int

func main() {
	go add(100, 200)
	fmt.Println("Result :", result)
}

func add(x, y int) {
	result = x + y
}
