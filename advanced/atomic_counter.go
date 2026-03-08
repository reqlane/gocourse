package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type atomicCounter struct {
	count int64
}

func (ac *atomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *atomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func heavyWork() {
	acc := 0
	for i := range 100 {
		acc += i
	}
}

// Time passed: 2.5346864s. Final counter value (sequential): 100000000
// Time passed: 1.8705976s. Final counter value (concurrent, no sync): 99862173
// Time passed: 1.6884583s. Final counter value (atomic counter): 100000000
// Time passed: 2.0661454s. Final counter value (mutexes): 100000000
func main() {

	var wg sync.WaitGroup
	numGoroutines := 4
	iterations := 100_000_000 // cool
	counter := &atomicCounter{}

	// No concurrency, sequential

	start := time.Now()
	for range iterations {
		heavyWork()
		counter.count++
	}
	timePassed := time.Since(start)
	fmt.Printf("Time passed: %10s. Final counter value (sequential): %d\n", timePassed, counter.getValue())

	// Concurrency, no sync
	counter.count = 0

	start = time.Now()
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range iterations / numGoroutines {
				heavyWork()
				counter.count++
			}
		}()
	}
	wg.Wait()
	timePassed = time.Since(start)
	fmt.Printf("Time passed: %10s. Final counter value (concurrent, no sync): %d\n", timePassed, counter.getValue())

	// Concurrency, atomic counter
	counter.count = 0

	start = time.Now()
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range iterations / numGoroutines {
				heavyWork()
				counter.increment()
			}
		}()
	}
	wg.Wait()
	timePassed = time.Since(start)
	fmt.Printf("Time passed: %10s. Final counter value (atomic counter): %d\n", timePassed, counter.getValue())

	// Concurrency, mutexes
	var mu sync.Mutex
	counter.count = 0

	start = time.Now()
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range iterations / numGoroutines {
				heavyWork()
				mu.Lock()
				counter.count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	timePassed = time.Since(start)
	fmt.Printf("Time passed: %10s. Final counter value (mutexes): %d\n", timePassed, counter.getValue())
}
