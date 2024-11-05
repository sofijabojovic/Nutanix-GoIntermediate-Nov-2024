package main

import (
	"fmt"
	"time"
)

func main() {
	primeCh := genPrimes()
	for no := range primeCh {
		fmt.Printf("Prime No : %d\n", no)
	}
}

func genPrimes() <-chan int {
	primeCh := make(chan int)
	// timeoutCh := timeout(10 * time.Second)
	timeoutCh := time.After(10 * time.Second)
	go func() {
		var no = 2
	LOOP:
		for ; ; no++ {
			select {
			case <-timeoutCh:
				fmt.Println("Timeout occurred.. exiting!")
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

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
