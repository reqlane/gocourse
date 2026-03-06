package advanced

import (
	"fmt"
	"time"
)

// func main() {
// 	// ================ BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY
// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 2
// 	}()

// 	fmt.Println("Value:", <-ch)
// 	fmt.Println("Value:", <-ch)
// 	fmt.Println("End of program")
// }

func main() {
	// ================ BLOCKING ON SEND ONLY IF THE BUFFER IS FULL
	// variable := make(chan Type, capacity)
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)
	}()

	// fmt.Println("Blocking starts")
	ch <- 3 // blocks because the buffer is full
	// fmt.Println("Blocking ends")

	// fmt.Println("Received:", <-ch)
	// fmt.Println("Received:", <-ch)
}
