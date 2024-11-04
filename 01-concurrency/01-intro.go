package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling the f1() execution through the scheduler to be scheduled to execute in future
	f2()
	// block the execution of main() so that the scheduler can look for other goroutines that are scheduled and execute them
	// DO NOT try the following approaches (poor man's synchronization techniques)
	time.Sleep(1 * time.Second)
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
