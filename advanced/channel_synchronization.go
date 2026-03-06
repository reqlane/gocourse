package advanced

import (
	"fmt"
	"time"
)

// func main() {

// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- struct{}{}
// 	}()

// 	<-done
// 	fmt.Println("Finished")
// }

// ====== SYNCHRONIZING MULTIPLE GOROUTINES
// func main() {

// 	numGoroutines := 3
// 	done := make(chan struct{}, 3)

// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(time.Second)
// 			done <- struct{}{}
// 		}(i)
// 	}

// 	for range numGoroutines {
// 		<-done // Wait for each goroutine to finish
// 	}

// 	fmt.Println("All goroutines are finished")
// }

// ===== SYNCHRONIZING DATA EXCHANGE
func main() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- fmt.Sprintf("%s %d", "hello", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data) // should be executed only by the sender, never the receiver
	}()

	// range over Type "chan any" stops receiving after close()
	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	}
}
