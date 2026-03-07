package advanced

import (
	"fmt"
	"time"
)

// === HANDLING MULTIPLE TICKERS
func main() {

	ticker1 := time.NewTicker(time.Second)
	ticker2 := time.NewTicker(480 * time.Millisecond)
	defer ticker1.Stop()
	defer ticker2.Stop()

	stop := time.After(10 * time.Second)

	for {
		select {
		case tick := <-ticker1.C:
			fmt.Println("Ticker 1:", tick)
		case tick := <-ticker2.C:
			fmt.Println("Ticker 2:", tick)
		case <-stop:
			fmt.Println("Tickers stopped")
			return
		}
	}
}

// === HANDLING ONE TICKER
// func main() {

// 	ticker := time.NewTicker(time.Second)
// 	stop := time.After(5 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case tick := <-ticker.C:
// 			fmt.Println("Tick at:", tick)
// 		case <-stop:
// 			fmt.Println("Stopping ticker")
// 			return
// 		}
// 	}
// }

// === SCHEDULING PERIODIC TASKS
// func periodicTask() {
// 	fmt.Println("Performing periodic task at:", time.Now())
// }

// func main() {

// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// === BASIC TICKER USAGE
// func main() {

// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()

// 	// for tick := range ticker.C {
// 	// 	fmt.Println("Tick at:", tick)
// 	// }

// 	i := 1
// 	for range 5 {
// 		i *= 2
// 		fmt.Println(i)
// 	}

// 	// for tick := range ticker.C {
// 	// 	i *= 2
// 	// 	fmt.Println(tick)
// 	// }
// }
