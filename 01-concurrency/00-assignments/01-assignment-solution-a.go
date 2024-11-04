/* Make the following concurrent */
package main

import (
	"fmt"
	"sync"
)

func main() {
	primesCh := genPrimes(3, 100)
	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func genPrimes(start, end int) <-chan int {
	primesCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for no := start; no <= end; no++ {
			wg.Add(1)
			go checkPrime(wg, no, primesCh)
		}
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func checkPrime(wg *sync.WaitGroup, no int, primesCh chan int) {
	defer wg.Done()
	if isPrime(no) {
		primesCh <- no
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
