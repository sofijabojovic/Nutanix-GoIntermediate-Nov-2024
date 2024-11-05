package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var ErrOSInterrupt = errors.New("OS Interrupt received")

func main() {
	defer func() {
		fmt.Println("[main deferred] app shutting down!.. releasing the resources!")
	}()
	fmt.Printf("PID : %d\n", os.Getpid())

	/*
			primeCh, errCh := genPrimes()
		LOOP:
			for {
				select {
				case no, isOpen := <-primeCh:
					if isOpen {
						fmt.Printf("[main] Prime No : %d\n", no)
						continue LOOP
					}
					break LOOP
				case err := <-errCh:
					fmt.Println("[main] Error occurred :", err)
					break LOOP
				}
			}
	*/

	primeCh, _ := genPrimes()
	for no := range primeCh {
		fmt.Printf("[main] Prime No : %d\n", no)
	}
}

func genPrimes() (<-chan int, <-chan error) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)
	primeCh := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		var no = 2
	LOOP:
		for ; ; no++ {
			select {
			case <-stopCh:
				fmt.Println("[genPrimes] OS Interrupt signal received.. exiting!")
				errCh <- ErrOSInterrupt
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
	return primeCh, errCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
