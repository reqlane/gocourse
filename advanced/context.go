package advanced

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	rootCtx := context.Background()
	// ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second)
	ctx, cancel := context.WithCancel(rootCtx)

	go func() {
		time.Sleep(2 * time.Second) // simulating a heavy task
		cancel()                    // manual cancelling
	}()

	ctx = context.WithValue(ctx, "requestID", "id8932457")
	ctx = context.WithValue(ctx, "name", "John Pork")
	ctx = context.WithValue(ctx, "IP", "345:4365:34256")

	go doWork(ctx)

	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request ID found")
	}

	logWithContext(ctx, "This is a test log message")
}

func logWithContext(ctx context.Context, msg string) {
	requestIDVal := ctx.Value("requestID")
	log.Printf("Request ID: %v - %v\n", requestIDVal, msg)
}

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation cancelled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

// func main() {

// 	ctx := context.TODO()

// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println("Result with context.TODO:", result)

// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
// 	defer cancel()

// 	result = checkEvenOdd(ctx, 10)
// 	fmt.Println("Result from timeout context:", result)

// 	time.Sleep(3 * time.Second)
// 	result = checkEvenOdd(ctx, 15)
// 	fmt.Println("Result after timeout:", result)
// }

// === DIFFERENCE BETWEEN context.TODO AND context.Background
// func main() {

// 	todoContext := context.TODO()
// 	contextBkg := context.Background()

// 	ctx := context.WithValue(todoContext, "name", "John")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	ctx1 := context.WithValue(contextBkg, "city", "New York")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))
// }
