package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Always hold mutex when using Condition

// The Bad Condition Usage Scenario:
// 1. Buffer is Full (5/5).
// 2. P1, P2, P3, P4 are all calling produce(). They see it's full and are now Wait()ing on the condition.
// 3. Consumer calls consume(). It takes 1 item (Buffer is now 4/5).
// 4. Consumer calls Signal() and then finishes.
// 5. The "Wrong" Wakeup: The Signal() wakes up P1.
// 6. P1 wakes up, sees the buffer is 4/5, adds an item (Buffer is 5/5).
// 7. P1 calls Signal() (as per your code).
// 8. The Disaster: P1’s signal wakes up P2.
// 9. P2 checks the loop: for len(b.items) == 5. It's true! So P2 goes back to sleep.
// Now, if the Consumer is also sleeping because it's waiting for more items, and all Producers are sleeping because the buffer is full... Deadlock.
// Solution - Broadcast() or several different Conditions

const BUFFERSIZE = 5

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == BUFFERSIZE {
		b.cond.Wait()
	}

	b.items = append(b.items, item)
	fmt.Println("Produced:", item)
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == 0 {
		b.cond.Wait()
	}

	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed:", item)
	b.cond.Signal()
	return item
}

func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range 10 {
		b.produce(i + 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 10 {
		b.consume()
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	buffer := newBuffer(BUFFERSIZE)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
}
