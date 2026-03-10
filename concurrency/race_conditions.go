package concurrency

import (
	"fmt"
	"sync"
)

// === go run -race file.go

type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func main() {

	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &counter{}

	// wg.Add(numGoroutines)
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range 1000 {
				counter.increment()
				// counter.count++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d", counter.getValue())
}
