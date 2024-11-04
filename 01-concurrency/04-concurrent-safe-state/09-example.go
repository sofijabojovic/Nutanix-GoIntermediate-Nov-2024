package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

// use sync/atomic package apis

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", atomic.LoadInt64(&count))
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&count, 1)
}
