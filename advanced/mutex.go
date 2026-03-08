package advanced

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int

	numGoroutines := 5

	increment := func(useMutex bool) {
		defer wg.Done()

		for range 50000000 {
			if useMutex {
				mu.Lock()
				counter++
				mu.Unlock()
			} else {
				counter++
			}
		}
	}

	start := time.Now()
	for range numGoroutines {
		wg.Add(1)
		go increment(false)
	}
	wg.Wait()
	timePassed := time.Since(start)
	fmt.Printf("Final counter value (no mutexes): %d. Time passed: %s\n", counter, timePassed)

	counter = 0
	start = time.Now()
	for range numGoroutines {
		wg.Add(1)
		go increment(true)
	}
	wg.Wait()
	timePassed = time.Since(start)
	fmt.Printf("Final counter value (with mutexes): %d. Time passed: %s\n", counter, timePassed)
}

// type counter struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (c *counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	c.count++
// }

// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	return c.count
// }

// func main() {

// 	var wg sync.WaitGroup
// 	numGoroutines := 10
// 	counter := &counter{}

// 	// wg.Add(numGoroutines)
// 	for range numGoroutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()

// 			for range 1000 {
// 				counter.increment()
// 				// counter.count++
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Printf("Final counter value: %d", counter.getValue())
// }
