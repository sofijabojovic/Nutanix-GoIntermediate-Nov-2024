/* Make the following concurrent */
package main

import (
	"fmt"
	"sync"
)

func main() {
	primesCh := genPrimes(3, 100, 5)
	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func genPrimes(start, end int, noOfWorkers int) <-chan int {
	primesCh := make(chan int)
	nosCh := produceNos(start, end)
	go func() {
		wg := &sync.WaitGroup{}
		for id := range noOfWorkers {
			wg.Add(1)
			// worker
			fmt.Printf("Starting worker : %d\n", id+1)
			go checkPrime(id+1, wg, nosCh, primesCh)
		}
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func produceNos(start, end int) <-chan int {
	nosCh := make(chan int)

	// input data producer
	go func() {
		for no := start; no <= end; no++ {
			nosCh <- no
		}
		close(nosCh)
	}()
	return nosCh
}

func checkPrime(workerId int, wg *sync.WaitGroup, nosCh <-chan int, primesCh chan int) {
	defer wg.Done()
	for no := range nosCh {
		if isPrime(no) {
			primesCh <- no
		}
	}
	fmt.Printf("Worker #%d completed!\n", workerId)
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
