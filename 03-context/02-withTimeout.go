package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	primeCh := genPrimes(cancelCtx)
	go func() {
		fmt.Println("Will stop after 5 secs.. Also can hit ENTER to stop...!")
		fmt.Scanln()
		cancel()
	}()
	for no := range primeCh {
		fmt.Printf("Prime No : %d\n", no)
	}
	if cancelCtx.Err() == context.Canceled {
		fmt.Println("Programmatic Stop signal received.. exiting!")
	}
	if cancelCtx.Err() == context.DeadlineExceeded {
		fmt.Println("Stop signal received by timeout.. exiting!")
	}
}

func genPrimes(ctx context.Context) <-chan int {
	primeCh := make(chan int)
	go func() {
		var no = 2
	LOOP:
		for ; ; no++ {
			select {
			case <-ctx.Done():
				/*
					if ctx.Err() == context.Canceled {
						fmt.Println("Programmatic Stop signal received.. exiting!")
					}
					if ctx.Err() == context.DeadlineExceeded {
						fmt.Println(" Stop signal received by timeout.. exiting!")
					}
				*/
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
