package advanced

import (
	"fmt"
	"sync"
	"time"
)

type rateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		limit:     limit,
		window:    window,
		resetTime: time.Now().Add(window),
	}
}

func (rl *rateLimiter) allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func main() {

	rateLimiter := NewRateLimiter(5, time.Second)

	// Simulate burst of (2*Limit) requests in a short time
	time.Sleep(700 * time.Millisecond)

	fmt.Println("===Burst of requests allowed in a short time")

	for range 30 {
		if rateLimiter.allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(50 * time.Millisecond)
	}

	// Simulate multiple simultaneous requests
	time.Sleep(time.Second)

	fmt.Println("===Simultaneous 10 requests")

	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if rateLimiter.allow() {
				fmt.Println("Request allowed")
			} else {
				fmt.Println("Request denied")
			}
		}()
	}
	wg.Wait()
}
