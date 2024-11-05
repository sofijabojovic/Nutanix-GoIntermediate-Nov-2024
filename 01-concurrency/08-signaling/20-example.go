package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	primeCh := genPrimes(stopCh)
	go func() {
		fmt.Println("Hit ENTER to stop...!")
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	for no := range primeCh {
		fmt.Printf("Prime No : %d\n", no)
	}
}

func genPrimes(stopCh chan struct{}) <-chan int {
	primeCh := make(chan int)
	go func() {
		var no = 2
	LOOP:
		for ; ; no++ {
			select {
			case <-stopCh:
				fmt.Println("Stop signal received.. exiting!")
				break LOOP
			default:
				if isPrime(no) {
					primeCh <- no
					time.Sleep(1 * time.Second)
				}
			}
		}
		close(primeCh)
	}()
	return primeCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
