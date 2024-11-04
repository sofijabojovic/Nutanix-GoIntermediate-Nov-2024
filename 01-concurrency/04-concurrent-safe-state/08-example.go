package main

import (
	"fmt"
	"sync"
)

// Custom type for concurrenct safe "count" manipulaton
type Counter struct {
	sync.Mutex // compose (inherit) sync.Mutex
	count      int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", counter.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}
