package advanced

import (
	"fmt"
	"time"
)

func main() {

	// ch := make(chan int)

	// === NON BLOCKING RECEIVE OPERATION
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received:", msg)
	// default:
	// 	fmt.Println("No messages available")
	// }

	// === NON BLOCKING SEND OPERATION
	// select {
	// case ch <- 1:
	// 	fmt.Println("Sent message")
	// default:
	// 	fmt.Println("Channel is not ready to receive")
	// }

	// === NON BLOCKING OPERATION IN REAL TIME SYSTEMS
	data := make(chan int)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case d, ok := <-data:
				if !ok {
					fmt.Println("Channel closed")
				} else {
					fmt.Println("Data received:", d)
				}
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data...")
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}
	close(data)

	time.Sleep(time.Second)
	quit <- struct{}{}
	time.Sleep(time.Second)
}
