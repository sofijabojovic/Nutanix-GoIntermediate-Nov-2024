package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	defer func() {
		fmt.Println("app shutting down!.. releasing the resources!")
	}()
	fmt.Printf("PID : %d\n", os.Getpid())
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)
	primeCh := genPrimes(stopCh)
	for no := range primeCh {
		fmt.Printf("Prime No : %d\n", no)
	}
}

func genPrimes(stopCh chan os.Signal) <-chan int {
	primeCh := make(chan int)
	go func() {
		var no = 2
	LOOP:
		for ; ; no++ {
			select {
			case <-stopCh:
				fmt.Println("OS Interrupt signal received.. exiting!")
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
