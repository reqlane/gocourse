package advanced

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// === Go by Example
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var reads uint64
	var writes uint64

	readOps := make(chan readOp)
	writeOps := make(chan writeOp)

	go func() {
		state := make(map[int]int)

		for {
			select {
			case readOp := <-readOps:
				readOp.resp <- state[readOp.key]
			case writeOp := <-writeOps:
				state[writeOp.key] = writeOp.val
				writeOp.resp <- true
			}
		}
	}()

	for range 100 {
		go func() {
			for {
				readOp := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				readOps <- readOp
				<-readOp.resp
				atomic.AddUint64(&reads, 1)
				time.Sleep(time.Millisecond)
			}
		}()

		go func() {
			for {
				writeOp := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writeOps <- writeOp
				<-writeOp.resp
				atomic.AddUint64(&writes, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readsFinal := atomic.LoadUint64(&reads)
	writesFinal := atomic.LoadUint64(&writes)
	fmt.Println("Read operations:", readsFinal)
	fmt.Println("Write operations:", writesFinal)
}

// === gocourse
// type statefulWorker struct {
// 	count int
// 	ch    chan int
// }

// func (w *statefulWorker) start() {
// 	go func() {
// 		for value := range w.ch {
// 			w.count += value
// 			fmt.Println("Current count:", w.count)
// 		}
// 	}()
// }

// func (w *statefulWorker) send(value int) {
// 	w.ch <- value
// }

// func main() {

// 	stWorker := &statefulWorker{
// 		ch: make(chan int),
// 	}

// 	stWorker.start()

// 	for i := range 5 {
// 		stWorker.send(i)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }
