package advanced

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type leakyBucket struct {
	mu           sync.Mutex
	capacity     int
	leakInterval time.Duration
	lastQueued   time.Time
}

func newLeakyBucket(capacity int, leakInterval time.Duration) *leakyBucket {
	return &leakyBucket{
		capacity:     capacity,
		leakInterval: leakInterval,
		lastQueued:   time.Now().Add(-leakInterval),
	}
}

func (lb *leakyBucket) allow(ctx context.Context) bool {
	lb.mu.Lock()

	now := time.Now()
	targetTime := lb.lastQueued.Add(lb.leakInterval)

	if time.Until(targetTime) > time.Duration(lb.capacity-1)*lb.leakInterval {
		lb.mu.Unlock()
		return false
	}

	if now.After(targetTime) {
		targetTime = now
	}

	lb.lastQueued = targetTime
	lb.mu.Unlock()

	waitDuration := time.Until(targetTime)
	if waitDuration <= 0 {
		return true
	}

	timer := time.NewTimer(waitDuration)
	defer timer.Stop()

	select {
	case <-timer.C:
		return true
	case <-ctx.Done():
		return false
	}
}

func main() {

	var wg sync.WaitGroup
	leakyBucket := newLeakyBucket(5, time.Second)

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			if leakyBucket.allow(ctx) {
				fmt.Println("Request ALLOWED")
			} else {
				if ctx.Err() == context.DeadlineExceeded {
					fmt.Println("Request TIMED OUT")
				} else {
					fmt.Println("Request DENIED")
				}
			}
		}()
	}

	wg.Wait()
}
