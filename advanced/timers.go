package advanced

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(1 * time.Second)
	defer timer1.Stop()
	timer2 := time.NewTimer(2 * time.Second)
	defer timer2.Stop()

	select {
	case <-timer1.C:
		fmt.Println("Timer 1 expired")
	case <-timer2.C:
		fmt.Println("Timer 2 expired")
	}
}

// === SCHEDULING DELAYED OPERATIONS
// func main() {

// 	timer := time.NewTimer(2 * time.Second) // non blocking

// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed operation executed")
// 	}()

// 	fmt.Println("Waiting...")
// 	time.Sleep(3 * time.Second) // blocking
// 	fmt.Println("End of the program")
// }

// === TIMEOUT
// func longRunningOperation() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {

// 	timeout := time.After(3 * time.Second)
// 	done := make(chan struct{})

// 	go func() {
// 		longRunningOperation()
// 		done <- struct{}{}
// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("Operation timed out")
// 	case <-done:
// 		fmt.Println("Operation completed")
// 	}
// }

// === BASIC TIMER USE
// func main() {

// 	fmt.Println("Starting app")
// 	timer := time.NewTimer(2 * time.Second)
// 	fmt.Println("Waiting for timer.C")
// 	timer.Reset(time.Second)
// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Timer stopped")
// 	}
// 	fmt.Println("Timer reset")
// 	timer.Reset(time.Second)
// 	<-timer.C // blocks
// 	fmt.Println("Timer expired")
// }
