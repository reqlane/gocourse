package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwmu    sync.RWMutex
	counter int
)

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock() // Write LOCKED / Read NOT LOCKED
	fmt.Println("Read counter:", counter)
	time.Sleep(time.Second)
	rwmu.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	rwmu.Lock() // Write/Read LOCKED
	counter = value
	fmt.Printf("Writter value %d for counter\n", value)
	time.Sleep(time.Second)
	rwmu.Unlock()
}

func main() {

	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	go writeCounter(&wg, 18)

	wg.Wait()
}
