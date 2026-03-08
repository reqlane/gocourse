package advanced

import (
	"fmt"
	"time"
)

type rateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

func newRateLimiter(rateLimit int, refillTime time.Duration) *rateLimiter {
	rl := &rateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}
	for range rateLimit {
		rl.tokens <- struct{}{}
	}
	go rl.startRefill()
	return rl
}

func (rl *rateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case rl.tokens <- struct{}{}:
		default:
		}
	}
}

func (rl *rateLimiter) allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {

	rateLimiter := newRateLimiter(5, time.Second)

	for range 10 {
		if rateLimiter.allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
